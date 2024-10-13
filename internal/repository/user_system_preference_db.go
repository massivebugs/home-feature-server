package repository

import (
	"context"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/system_preference"
)

type UserSystemPreferenceRepository struct {
	querier queries.Querier
}

func NewUserSystemPreferenceRepository(querier queries.Querier) *UserSystemPreferenceRepository {
	return &UserSystemPreferenceRepository{
		querier: querier,
	}
}

func (r *UserSystemPreferenceRepository) CreateUserSystemPreference(ctx context.Context, db db.DB, arg system_preference.CreateUserSystemPreferenceParams) (uint32, error) {
	result, err := r.querier.CreateUserSystemPreference(ctx, db, queries.CreateUserSystemPreferenceParams{
		UserID:   arg.UserID,
		Language: arg.Language,
	})
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

func (r *UserSystemPreferenceRepository) GetUserSystemPreferenceExists(ctx context.Context, db db.DB, userID uint32) (bool, error) {
	return r.querier.GetUserSystemPreferenceExists(ctx, db, userID)
}

func (r *UserSystemPreferenceRepository) GetUserSystemPreference(ctx context.Context, db db.DB, userID uint32) (*system_preference.UserSystemPreference, error) {
	data, err := r.querier.GetUserSystemPreference(ctx, db, userID)
	if err != nil {
		return nil, err
	}

	return system_preference.NewUserSystemPreferenceFromQueries(data), nil
}

func (r *UserSystemPreferenceRepository) UpdateUserSystemPreference(ctx context.Context, db db.DB, arg system_preference.UpdateUserSystemPreferenceParams) error {
	return r.querier.UpdateUserSystemPreference(ctx, db, queries.UpdateUserSystemPreferenceParams{
		UserID:   arg.UserID,
		Language: arg.Language,
	})
}
