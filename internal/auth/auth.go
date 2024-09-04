package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/massivebugs/home-feature-server/db/service/auth_repository"
	"github.com/massivebugs/home-feature-server/internal/api"
)

type Auth struct {
	db       *sql.DB
	authRepo auth_repository.Querier
}

func NewAuth(
	db *sql.DB,
	authRepo auth_repository.Querier,
) *Auth {
	return &Auth{
		db:       db,
		authRepo: authRepo,
	}
}

func (s *Auth) CreateAuthUser(ctx context.Context, req *CreateUserRequestDTO) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Check if user already exists
	_, err = s.authRepo.GetUserByName(ctx, s.db, req.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	} else if err == nil {
		return api.NewAPIError(api.CodeBadRequest, errors.New("user already exists"))
	}

	// Create new user
	result, err := s.authRepo.CreateUser(ctx, tx, req.Username)
	if err != nil {
		return err
	}

	// Retrieve ID
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Hash password
	hashedPassword, err := GeneratePasswordHash(req.Password)
	if err != nil {
		return err
	}

	// Create user password
	p := auth_repository.CreateUserPasswordParams{
		UserID:       uint32(id),
		PasswordHash: hashedPassword,
	}
	_, err = s.authRepo.CreateUserPassword(ctx, tx, p)
	if err != nil {
		return err
	}

	return tx.Commit()
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
	req *UserAuthRequestDTO,
) (string, string, error) {
	// Retrieve user
	u, err := s.authRepo.GetUserByName(ctx, s.db, req.Username)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return "", "", api.NewAPIError(api.CodeNotFound, errors.New("username or password does not match"))
	} else if err != nil {
		return "", "", err
	}

	// Retrieve user password
	up, err := s.authRepo.GetUserPasswordByUserID(ctx, s.db, u.ID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return "", "", api.NewAPIError(api.CodeNotFound, errors.New("username or password does not match"))
	} else if err != nil {
		return "", "", err
	}

	// Check if hash matches
	err = CheckPasswordHash(up.PasswordHash, req.Password)
	if err != nil {
		return "", "", api.NewAPIError(api.CodeNotFound, errors.New("username or password does not match"))
	}

	// TODO: Generate and store refresh token value in database and in the JWT claim!
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

	_, err = s.authRepo.CreateUserRefreshToken(
		ctx,
		s.db,
		auth_repository.CreateUserRefreshTokenParams{
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
	claims *JWTClaims,
) (string, string, error) {
	urt, err := s.authRepo.GetUserRefreshTokenByValue(
		ctx,
		s.db,
		auth_repository.GetUserRefreshTokenByValueParams{
			UserID: claims.UserID,
			Value:  claims.TokenID,
		},
	)

	// If the token doesn't exist, this means this token id(value) has been invalidated
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return "", "", api.NewAPIError(api.CodeUnauthorized, errors.New("token is invalid"))
	} else if err != nil {
		return "", "", err
	}

	// Set custom claims
	tokenBuilder := NewJWTBuilder(now, jwtExpireSeconds, JWTCustomClaims{UserID: claims.UserID})
	refreshTokenBuilder := NewJWTBuilder(now, refreshJwtExpireSeconds, JWTCustomClaims{UserID: claims.UserID, TokenID: urt.Value})

	// Generate encoded token and send it as response.
	tokenStr, err := tokenBuilder.CreateAndSignToken(jwtSigningMethod, jwtSecret)
	if err != nil {
		return "", "", err
	}
	refreshTokenStr, err := refreshTokenBuilder.CreateAndSignToken(refreshJwtSigningMethod, refreshJwtSecret)
	if err != nil {
		return "", "", err
	}

	// Update refresh token expiry time
	err = s.authRepo.UpdateUserRefreshTokenExpiresAt(ctx, s.db, auth_repository.UpdateUserRefreshTokenExpiresAtParams{
		ID: urt.ID,
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

func (s *Auth) GetAuthUser(ctx context.Context, jwtClaims *JWTClaims) (AuthUser, error) {
	u, err := s.authRepo.GetUser(ctx, s.db, jwtClaims.UserID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return AuthUser{}, api.NewAPIError(api.CodeForbidden, errors.New("user does not exist"))
	} else if err != nil {
		return AuthUser{}, err
	}

	return NewAuthUser(u, jwtClaims), nil
}
