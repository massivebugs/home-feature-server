package repository

import (
	"context"
	"errors"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/auth"
)

type UserPasswordDBRepository struct {
	querier queries.Querier
}

func NewUserPasswordDBRepository(querier queries.Querier) *UserPasswordDBRepository {
	return &UserPasswordDBRepository{
		querier: querier,
	}
}

func (r *UserPasswordDBRepository) CreateUserPassword(ctx context.Context, db db.DB, arg auth.CreateUserPasswordParams) error {
	_, err := r.querier.CreateUserPassword(ctx, db, queries.CreateUserPasswordParams{
		UserID:       arg.UserID,
		PasswordHash: arg.PasswordHash,
	})

	return err
}

func (r *UserPasswordDBRepository) GetUserPasswordByUserID(ctx context.Context, db db.DB, userID uint32) (string, error) {
	data, err := r.querier.GetUserPasswordByUserID(ctx, db, userID)
	if err != nil {
		return "", err
	}

	return data.PasswordHash, err
}

func (r *UserPasswordDBRepository) UpdateUserPassword(ctx context.Context, db db.DB, arg auth.UpdateUserPasswordParams) error {
	return errors.New("not implemented yet")
}
