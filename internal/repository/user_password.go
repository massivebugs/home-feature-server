package repository

import (
	"context"
	"errors"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/auth"
)

type UserPasswordRepository struct {
	querier queries.Querier
}

var _ auth.IUserPasswordRepository = (*UserPasswordRepository)(nil)

func NewUserPasswordRepository(querier queries.Querier) *UserPasswordRepository {
	return &UserPasswordRepository{
		querier: querier,
	}
}

func (r *UserPasswordRepository) CreateUserPassword(ctx context.Context, db db.DB, arg auth.CreateUserPasswordParams) error {
	_, err := r.querier.CreateUserPassword(ctx, db, queries.CreateUserPasswordParams{
		UserID:       arg.UserID,
		PasswordHash: arg.PasswordHash,
	})

	return err
}

func (r *UserPasswordRepository) GetUserPasswordByUserID(ctx context.Context, db db.DB, userID uint32) (string, error) {
	data, err := r.querier.GetUserPasswordByUserID(ctx, db, userID)
	if err != nil {
		return "", err
	}

	return data.PasswordHash, err
}

func (r *UserPasswordRepository) UpdateUserPassword(ctx context.Context, db db.DB, arg auth.UpdateUserPasswordParams) error {
	return errors.New("not implemented yet")
}