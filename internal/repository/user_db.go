package repository

import (
	"context"
	"errors"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/auth"
)

type UserDBRepository struct {
	querier queries.Querier
}

func NewUserDBRepository(querier queries.Querier) *UserDBRepository {
	return &UserDBRepository{
		querier: querier,
	}
}

func (r *UserDBRepository) GetUsernameExists(ctx context.Context, db db.DB, name string) (bool, error) {
	return r.querier.GetUsernameExists(ctx, db, name)
}

func (r *UserDBRepository) CreateUser(ctx context.Context, db db.DB, name string) (uint32, error) {
	result, err := r.querier.CreateUser(ctx, db, name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

func (r *UserDBRepository) DeleteUser(ctx context.Context, db db.DB, id uint32) error {
	return errors.New("not implemented yet")
}

func (r *UserDBRepository) GetUser(ctx context.Context, db db.DB, id uint32) (*auth.AuthUser, error) {
	data, err := r.querier.GetUser(ctx, db, id)
	if err != nil {
		return nil, err
	}

	return auth.NewAuthUserFromQueries(data), nil
}

func (r *UserDBRepository) GetUserByName(ctx context.Context, db db.DB, name string) (*auth.AuthUser, error) {
	data, err := r.querier.GetUserByName(ctx, db, name)
	if err != nil {
		return nil, err
	}

	return auth.NewAuthUserFromQueries(data), nil
}

func (r *UserDBRepository) UpdateUser(ctx context.Context, db db.DB, arg auth.UpdateUserParams) error {
	return errors.New("not implemented yet")
}
