package cashbunny

import (
	"context"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
)

type CreateUserCurrencyParams struct {
	UserID       uint32
	CurrencyCode string
}

type IUserPreferenceRepository interface {
	GetUserPreferenceExistsByUserID(ctx context.Context, db db.DB, userID uint32) (bool, error)
	CreateUserPreferences(ctx context.Context, db db.DB, userID uint32) (uint32, error)
	GetUserPreferencesByUserID(ctx context.Context, db db.DB, userID uint32) (*UserPreferences, error)
	CreateUserCurrency(ctx context.Context, db db.DB, params CreateUserCurrencyParams) (uint32, error)
	ListUserCurrencies(ctx context.Context, db db.DB, userID uint32) ([]string, error)
}

type UserPreferencesRepository struct {
	querier queries.Querier
}

var _ IUserPreferenceRepository = (*UserPreferencesRepository)(nil)

func NewUserPreferencesRepository(querier queries.Querier) *UserPreferencesRepository {
	return &UserPreferencesRepository{
		querier: querier,
	}
}

func (r *UserPreferencesRepository) GetUserPreferenceExistsByUserID(ctx context.Context, db db.DB, userID uint32) (bool, error) {
	result, err := r.querier.GetUserPreferenceExistsByUserID(ctx, db, userID)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (r *UserPreferencesRepository) CreateUserPreferences(ctx context.Context, db db.DB, userID uint32) (uint32, error) {
	result, err := r.querier.CreateUserPreferences(ctx, db, userID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

func (r *UserPreferencesRepository) GetUserPreferencesByUserID(ctx context.Context, db db.DB, userID uint32) (*UserPreferences, error) {
	ucs, err := r.ListUserCurrencies(ctx, db, userID)
	if err != nil {
		return nil, err
	}

	upData, err := r.querier.GetUserPreferenceByUserID(ctx, db, userID)
	if err != nil {
		return nil, err
	}

	up := NewUserPreferencesFromQueries(upData, ucs)

	return up, nil
}

func (r *UserPreferencesRepository) CreateUserCurrency(ctx context.Context, db db.DB, params CreateUserCurrencyParams) (uint32, error) {
	result, err := r.querier.CreateUserCurrency(ctx, db, queries.CreateUserCurrencyParams{
		UserID:       params.UserID,
		CurrencyCode: params.CurrencyCode,
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

func (r *UserPreferencesRepository) ListUserCurrencies(ctx context.Context, db db.DB, userID uint32) ([]string, error) {
	data, err := r.querier.ListUserCurrencies(ctx, db, userID)
	if err != nil {
		return nil, err
	}

	ucs := make([]string, len(data))
	for idx, d := range data {
		ucs[idx] = d.CurrencyCode
	}

	return ucs, nil
}
