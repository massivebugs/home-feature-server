package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/internal/app"
)

type Auth struct {
	db       *db.Handle
	userRepo IUserRepository
	passRepo IUserPasswordRepository
	rtRepo   IUserRefreshTokenRepository
}

func NewAuth(
	db *db.Handle,
	userRepo IUserRepository,
	passRepo IUserPasswordRepository,
	rtRepo IUserRefreshTokenRepository,
) *Auth {
	return &Auth{
		db:       db,
		userRepo: userRepo,
		passRepo: passRepo,
		rtRepo:   rtRepo,
	}
}

func (s *Auth) CreateAuthUser(ctx context.Context, req *CreateUserDTO) error {
	return s.db.WithTx(ctx, func(tx db.DB) error {
		// Check if user already exists
		exists, err := s.userRepo.GetUsernameExists(ctx, s.db, req.Username)
		if err != nil {
			return err
		}
		if exists {
			return app.NewAppError(app.CodeBadRequest, errors.New("user already exists"))
		}

		// Create new user
		userID, err := s.userRepo.CreateUser(ctx, tx, req.Username)
		if err != nil {
			return err
		}

		// Hash password
		hashedPassword, err := GeneratePasswordHash(req.Password)
		if err != nil {
			return err
		}

		// Create user password
		p := CreateUserPasswordParams{
			UserID:       uint32(userID),
			PasswordHash: hashedPassword,
		}
		return s.passRepo.CreateUserPassword(ctx, tx, p)
	})
}

func (s *Auth) CreateJWTToken(
	ctx context.Context,
	now time.Time,
	jwtSigningMethod *jwt.SigningMethodHMAC,
	jwtSecret string,
	jwtExpireSeconds int,
	refreshJwtSigningMethod *jwt.SigningMethodHMAC,
	refreshJwtSecret string,
	refreshJwtExpireSeconds int,
	req *UserAuthDTO,
) (string, string, error) {
	// Retrieve user
	// TODO
	u, err := s.userRepo.GetUserByName(ctx, s.db, req.Username)
	if err != nil {
		return "", "", app.NewAppError(app.CodeNotFound, errors.New("username or password does not match"))
	}

	// Retrieve user password
	// TODO
	hash, err := s.passRepo.GetUserPasswordByUserID(ctx, s.db, u.ID)
	if err != nil {
		return "", "", app.NewAppError(app.CodeNotFound, errors.New("username or password does not match"))
	}

	// Check if hash matches
	err = CheckPasswordHash(hash, req.Password)
	if err != nil {
		return "", "", app.NewAppError(app.CodeNotFound, errors.New("username or password does not match"))
	}

	tokenID := GenerateRandomString(50)

	// Set custom claims
	tokenBuilder := NewJWTBuilder(now, jwtExpireSeconds, JWTCustomClaims{UserID: u.ID})
	refreshTokenBuilder := NewJWTBuilder(now, refreshJwtExpireSeconds, JWTCustomClaims{UserID: u.ID, TokenID: tokenID})

	// Generate encoded token and send it as response.
	tokenStr, err := tokenBuilder.CreateAndSignToken(jwtSigningMethod, jwtSecret)
	if err != nil {
		return "", "", err
	}
	refreshTokenStr, err := refreshTokenBuilder.CreateAndSignToken(refreshJwtSigningMethod, refreshJwtSecret)
	if err != nil {
		return "", "", err
	}

	err = s.rtRepo.CreateUserRefreshToken(
		ctx,
		s.db,
		CreateUserRefreshTokenParams{
			UserID: u.ID,
			Value:  tokenID,
			ExpiresAt: sql.NullTime{
				Time:  refreshTokenBuilder.claims.ExpiresAt.Time,
				Valid: true,
			},
		})
	if err != nil {
		return "", "", err
	}

	return tokenStr, refreshTokenStr, nil
}

func (s *Auth) RefreshJWTToken(
	ctx context.Context,
	now time.Time,
	jwtSigningMethod *jwt.SigningMethodHMAC,
	jwtSecret string,
	jwtExpireSeconds int,
	refreshJwtSigningMethod *jwt.SigningMethodHMAC,
	refreshJwtSecret string,
	refreshJwtExpireSeconds int,
	userID uint32,
	tokenID string,
) (string, string, error) {
	exists, err := s.rtRepo.GetUserRefreshTokenExistsByValue(
		ctx,
		s.db,
		GetUserRefreshTokenExistsByValueParams{
			UserID: userID,
			Value:  tokenID,
		},
	)
	if err != nil {
		return "", "", err
	}

	// If the token doesn't exist, this means this token id(value) has been invalidated
	if !exists {
		return "", "", app.NewAppError(app.CodeUnauthorized, errors.New("token is invalid"))
	}

	newTokenID := GenerateRandomString(50)

	// Set custom claims
	tokenBuilder := NewJWTBuilder(now, jwtExpireSeconds, JWTCustomClaims{UserID: userID})
	refreshTokenBuilder := NewJWTBuilder(now, refreshJwtExpireSeconds, JWTCustomClaims{UserID: userID, TokenID: newTokenID})

	// Generate encoded token and send it as response.
	tokenStr, err := tokenBuilder.CreateAndSignToken(jwtSigningMethod, jwtSecret)
	if err != nil {
		return "", "", err
	}
	refreshTokenStr, err := refreshTokenBuilder.CreateAndSignToken(refreshJwtSigningMethod, refreshJwtSecret)
	if err != nil {
		return "", "", err
	}

	err = s.rtRepo.CreateUserRefreshToken(
		ctx,
		s.db,
		CreateUserRefreshTokenParams{
			UserID: userID,
			Value:  newTokenID,
			ExpiresAt: sql.NullTime{
				Time:  refreshTokenBuilder.claims.ExpiresAt.Time,
				Valid: true,
			},
		})
	if err != nil {
		return "", "", err
	}

	return tokenStr, refreshTokenStr, nil
}

func (s *Auth) GetAuthUser(ctx context.Context, jwtClaims *JWTClaims) (*AuthUser, error) {
	u, err := s.userRepo.GetUser(ctx, s.db, jwtClaims.UserID)
	if err != nil {
		return nil, err
	}

	u.SetLoginTime(jwtClaims)

	return u, nil
}
