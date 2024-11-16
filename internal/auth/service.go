package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/internal/app"
	"github.com/massivebugs/home-feature-server/internal/util"
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

func (s *Auth) CreateUser(ctx context.Context, username string, email string, password string) error {
	return s.db.WithTx(ctx, func(tx db.DB) error {
		// Check if username or email already exists
		exists, err := s.userRepo.GetUsernameOrEmailExists(ctx, s.db, GetUsernameOrEmailExistsParams{
			Name:  username,
			Email: email,
		})
		if err != nil {
			return err
		}
		if exists {
			return app.NewAppError(app.CodeBadRequest, errors.New("username or email already exists"))
		}

		// Create new user
		userID, err := s.userRepo.CreateUser(ctx, tx, CreateUserParams{
			Name:  username,
			Email: email,
		})
		if err != nil {
			return err
		}

		// Hash password
		hashedPassword, err := GeneratePasswordHash(password)
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
	username string,
	password string,
) (string, error) {
	u, err := s.userRepo.GetUserByName(ctx, s.db, username)
	if err != nil {
		return "", app.NewAppError(app.CodeBadRequest, errors.New("username or password does not match"))
	}

	if u.IsDisabled() {
		return "", app.NewAppError(app.CodeForbidden, errors.New("user is disabled"))
	}

	hash, err := s.passRepo.GetUserPasswordByUserID(ctx, s.db, u.Id)
	if err != nil {
		return "", app.NewAppError(app.CodeBadRequest, errors.New("username or password does not match"))
	}

	err = CheckPasswordHash(hash, password)
	if err != nil {
		return "", app.NewAppError(app.CodeBadRequest, errors.New("username or password does not match"))
	}

	// Set custom claims and build/sign token
	tokenBuilder := NewJWTBuilder(now, jwtExpireSeconds, JWTCustomClaims{UserID: u.Id})
	tokenStr, err := tokenBuilder.CreateAndSignToken(jwtSigningMethod, jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (s *Auth) CreateJWTRefreshToken(
	ctx context.Context,
	now time.Time,
	refreshJwtSigningMethod *jwt.SigningMethodHMAC,
	refreshJwtSecret string,
	refreshJwtExpireSeconds int,
	userId uint32,
) (string, error) {
	u, err := s.userRepo.GetUser(ctx, s.db, userId)
	if err != nil {
		return "", err
	}

	if u.IsDisabled() {
		return "", app.NewAppError(app.CodeForbidden, errors.New("user is disabled"))
	}

	// Set custom claims and build/sign token
	tokenID := util.GenerateRandomString(50)
	refreshTokenBuilder := NewJWTBuilder(now, refreshJwtExpireSeconds, JWTCustomClaims{UserID: u.Id, TokenID: tokenID})
	refreshTokenStr, err := refreshTokenBuilder.CreateAndSignToken(refreshJwtSigningMethod, refreshJwtSecret)
	if err != nil {
		return "", err
	}

	err = s.rtRepo.CreateUserRefreshToken(
		ctx,
		s.db,
		CreateUserRefreshTokenParams{
			UserID:    u.Id,
			Value:     tokenID,
			ExpiresAt: refreshTokenBuilder.claims.ExpiresAt.Time,
		})
	if err != nil {
		return "", err
	}

	return refreshTokenStr, nil
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
) (string, error) {
	exists, err := s.rtRepo.GetUserRefreshTokenExistsByValue(
		ctx,
		s.db,
		GetUserRefreshTokenExistsByValueParams{
			UserID: userID,
			Value:  tokenID,
		},
	)
	if err != nil {
		return "", err
	}

	// If the token doesn't exist, this means this token id(value) has been invalidated
	if !exists {
		return "", app.NewAppError(app.CodeUnauthorized, errors.New("token is invalid"))
	}

	u, err := s.userRepo.GetUser(ctx, s.db, userID)
	if err != nil {
		return "", err
	}

	if u.IsDisabled() {
		return "", app.NewAppError(app.CodeForbidden, errors.New("user is disabled"))
	}

	// Set custom claims and build/sign token
	tokenBuilder := NewJWTBuilder(now, jwtExpireSeconds, JWTCustomClaims{UserID: userID})
	tokenStr, err := tokenBuilder.CreateAndSignToken(jwtSigningMethod, jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (s *Auth) GetAuthUser(ctx context.Context, userID uint32, loginTime time.Time) (*User, error) {
	u, err := s.userRepo.GetUser(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	u.LoggedInAt = loginTime

	return u, nil
}
