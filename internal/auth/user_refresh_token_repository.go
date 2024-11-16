package auth

import (
	"context"
	"time"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
)

type CreateUserRefreshTokenParams struct {
	UserID    uint32
	Value     string
	ExpiresAt time.Time
}

type DeleteUserRefreshTokenParams struct {
	UserID uint32
	Value  string
}

type GetUserRefreshTokenExistsByValueParams struct {
	UserID uint32
	Value  string
}

type DeleteUserRefreshTokenByValueParams struct {
	UserID uint32
	Value  string
}

type IUserRefreshTokenRepository interface {
	CreateUserRefreshToken(ctx context.Context, db db.DB, arg CreateUserRefreshTokenParams) error
	GetUserRefreshTokenExistsByValue(ctx context.Context, db db.DB, arg GetUserRefreshTokenExistsByValueParams) (bool, error)
	DeleteUserRefreshTokenByValue(ctx context.Context, db db.DB, arg DeleteUserRefreshTokenByValueParams) error
}

type UserRefreshTokenRepository struct {
	querier queries.Querier
}

var _ IUserRefreshTokenRepository = (*UserRefreshTokenRepository)(nil)

func NewUserRefreshTokenRepository(querier queries.Querier) *UserRefreshTokenRepository {
	return &UserRefreshTokenRepository{
		querier: querier,
	}
}

func (r *UserRefreshTokenRepository) CreateUserRefreshToken(ctx context.Context, db db.DB, arg CreateUserRefreshTokenParams) error {
	_, err := r.querier.CreateUserRefreshToken(
		ctx,
		db,
		queries.CreateUserRefreshTokenParams{
			UserID:    arg.UserID,
			Value:     arg.Value,
			ExpiresAt: arg.ExpiresAt,
		})

	return err
}

func (r *UserRefreshTokenRepository) GetUserRefreshTokenExistsByValue(ctx context.Context, db db.DB, arg GetUserRefreshTokenExistsByValueParams) (bool, error) {
	return r.querier.GetUserRefreshTokenExistsByValue(ctx, db, queries.GetUserRefreshTokenExistsByValueParams{
		UserID: arg.UserID,
		Value:  arg.Value,
	})
}

func (r *UserRefreshTokenRepository) DeleteUserRefreshTokenByValue(ctx context.Context, db db.DB, arg DeleteUserRefreshTokenByValueParams) error {
	_, err := r.querier.DeleteUserRefreshTokenByValue(ctx, db, queries.DeleteUserRefreshTokenByValueParams(arg))

	return err
}
