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

func (s *Auth) CreateJWTToken(ctx context.Context, jwtSecret string, req *UserAuthRequestDTO) (string, error) {
	// Retrieve user
	u, err := s.authRepo.GetUserByName(ctx, s.db, req.Username)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return "", api.NewAPIError(api.CodeNotFound, errors.New("username or password does not match"))
	} else if err != nil {
		return "", err
	}

	// Retrieve user password
	up, err := s.authRepo.GetUserPasswordByUserID(ctx, s.db, u.ID)
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
	jb := NewJWTBuilder(time.Now(), 72, u.ID)

	// Generate encoded token and send it as response.
	tokenStr, err := jb.CreateAndSignToken(jwt.SigningMethodHS256, jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
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
