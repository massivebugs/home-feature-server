package repository

import (
	"context"
	"errors"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/auth"
)

type UserRepository struct {
	querier queries.Querier
}

var _ auth.IUserRepository = (*UserRepository)(nil)

func NewUserRepository(querier queries.Querier) *UserRepository {
	return &UserRepository{
		querier: querier,
	}
}

func (r *UserRepository) GetUsernameOrEmailExists(ctx context.Context, db db.DB, arg auth.GetUsernameOrEmailExistsParams) (bool, error) {
	return r.querier.GetUsernameOrEmailExists(ctx, db, queries.GetUsernameOrEmailExistsParams(arg))
}

func (r *UserRepository) CreateUser(ctx context.Context, db db.DB, arg auth.CreateUserParams) (uint32, error) {
	result, err := r.querier.CreateUser(ctx, db, queries.CreateUserParams(arg))
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, db db.DB, id uint32) error {
	return errors.New("not implemented yet")
}

func (r *UserRepository) GetUser(ctx context.Context, db db.DB, id uint32) (*auth.User, error) {
	data, err := r.querier.GetUser(ctx, db, id)
	if err != nil {
		return nil, err
	}

	return auth.NewUserFromQueries(data), nil
}

func (r *UserRepository) GetUserByName(ctx context.Context, db db.DB, name string) (*auth.User, error) {
	data, err := r.querier.GetUserByName(ctx, db, name)
	if err != nil {
		return nil, err
	}

	return auth.NewUserFromQueries(data), nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, db db.DB, arg auth.UpdateUserParams) error {
	return errors.New("not implemented yet")
}
