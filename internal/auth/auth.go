package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/massivebugs/home-feature-server/api/dto"
	"github.com/massivebugs/home-feature-server/db/service/user"
	"github.com/massivebugs/home-feature-server/db/service/user_password"
	"github.com/massivebugs/home-feature-server/internal/api"
)

type Auth struct {
	db               *sql.DB
	userRepo         user.Querier
	userPasswordRepo user_password.Querier
}

func NewAuth(
	db *sql.DB,
	userRepo user.Querier,
	userPasswordRepo user_password.Querier,
) *Auth {
	return &Auth{
		db:               db,
		userRepo:         userRepo,
		userPasswordRepo: userPasswordRepo,
	}
}

func (s *Auth) CreateAuthUser(ctx context.Context, req *dto.UserAuthRequestDTO) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	// Check if user already exists
	existingUser, err := s.userRepo.GetUserByName(ctx, s.db, req.Username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return api.NewAPIError(api.CodeBadRequest, errors.New("user already exists"))
	}

	defer tx.Rollback()
	// Create new user
	result, err := s.userRepo.CreateUser(ctx, tx, req.Username)
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
	p := user_password.CreateUserPasswordParams{
		UserID:       uint32(id),
		PasswordHash: hashedPassword,
	}
	_, err = s.userPasswordRepo.CreateUserPassword(ctx, tx, p)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Auth) CreateJWTToken(ctx context.Context, jwtSecret string, req *dto.UserAuthRequestDTO) (string, error) {
	// Retrieve user
	u, err := s.userRepo.GetUserByName(ctx, s.db, req.Username)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return "", api.NewAPIError(api.CodeNotFound, errors.New("username or password does not match"))
	} else if err != nil {
		return "", err
	}

	// Retrieve user password
	up, err := s.userPasswordRepo.GetUserPasswordByUserID(ctx, s.db, u.ID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return "", api.NewAPIError(api.CodeNotFound, errors.New("username or password does not match"))
	} else if err != nil {
		return "", err
	}

	// Check if hash matches
	err = CheckPasswordHash(up.PasswordHash, req.Password)
	if err != nil {
		return "", api.NewAPIError(api.CodeNotFound, errors.New("username or password does not match"))
	}

	// Set custom claims
	claims := NewJWTClaims(time.Now(), 72, u.ID)

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	tokenStr, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
