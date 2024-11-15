package auth

import (
	"context"
	"errors"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
)

type GetUsernameOrEmailExistsParams struct {
	Name  string
	Email string
}

type CreateUserParams struct {
	Name  string
	Email string
}

type UpdateUserParams struct {
	Name  string
	Email string
	ID    uint32
}

type IUserRepository interface {
	GetUsernameOrEmailExists(ctx context.Context, db db.DB, arg GetUsernameOrEmailExistsParams) (bool, error)
	CreateUser(ctx context.Context, db db.DB, arg CreateUserParams) (uint32, error)
	DeleteUser(ctx context.Context, db db.DB, id uint32) error
	GetUser(ctx context.Context, db db.DB, id uint32) (*User, error)
	GetUserByName(ctx context.Context, db db.DB, name string) (*User, error)
	UpdateUser(ctx context.Context, db db.DB, arg UpdateUserParams) error
}

type UserRepository struct {
	querier queries.Querier
}

var _ IUserRepository = (*UserRepository)(nil)

func NewUserRepository(querier queries.Querier) *UserRepository {
	return &UserRepository{
		querier: querier,
	}
}

func (r *UserRepository) GetUsernameOrEmailExists(ctx context.Context, db db.DB, arg GetUsernameOrEmailExistsParams) (bool, error) {
	return r.querier.GetUsernameOrEmailExists(ctx, db, queries.GetUsernameOrEmailExistsParams(arg))
}

func (r *UserRepository) CreateUser(ctx context.Context, db db.DB, arg CreateUserParams) (uint32, error) {
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

func (r *UserRepository) GetUser(ctx context.Context, db db.DB, id uint32) (*User, error) {
	data, err := r.querier.GetUser(ctx, db, id)
	if err != nil {
		return nil, err
	}

	return NewUserFromQueries(data), nil
}

func (r *UserRepository) GetUserByName(ctx context.Context, db db.DB, name string) (*User, error) {
	data, err := r.querier.GetUserByName(ctx, db, name)
	if err != nil {
		return nil, err
	}

	return NewUserFromQueries(data), nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, db db.DB, arg UpdateUserParams) error {
	return errors.New("not implemented yet")
}
