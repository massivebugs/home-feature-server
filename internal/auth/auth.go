package auth

import (
	"context"
	"database/sql"

	"github.com/massivebugs/home-feature-server/api/dto"
	"github.com/massivebugs/home-feature-server/db/service/user"
	"github.com/massivebugs/home-feature-server/db/service/user_password"
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

func (s *Auth) CreateAuthUser(ctx context.Context, req *dto.CreateUserRequestDTO) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()
	result, err := s.userRepo.CreateUser(ctx, tx, req.Username)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	hashedPassword, err := GeneratePasswordHash(req.Password)
	if err != nil {
		return err
	}

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
