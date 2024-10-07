package repository

import (
	"context"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/auth"
)

type UserRefreshTokenDBRepository struct {
	querier queries.Querier
}

func NewUserRefreshTokenDBRepository(querier queries.Querier) *UserRefreshTokenDBRepository {
	return &UserRefreshTokenDBRepository{
		querier: querier,
	}
}

func (r *UserRefreshTokenDBRepository) CreateUserRefreshToken(ctx context.Context, db db.DB, arg auth.CreateUserRefreshTokenParams) error {
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

func (r *UserRefreshTokenDBRepository) GetUserRefreshTokenExistsByValue(ctx context.Context, db db.DB, arg auth.GetUserRefreshTokenExistsByValueParams) (bool, error) {
	return r.querier.GetUserRefreshTokenExistsByValue(ctx, db, queries.GetUserRefreshTokenExistsByValueParams{
		UserID: arg.UserID,
		Value:  arg.Value,
	})
}
