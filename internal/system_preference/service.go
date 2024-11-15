package system_preference

import (
	"context"
	"errors"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/internal/app"
)

type SystemPreference struct {
	db      *db.Handle
	uspRepo IUserSystemPreferenceRepository
}

func NewSystemPreference(
	db *db.Handle,
	uspRepo IUserSystemPreferenceRepository,
) *SystemPreference {
	return &SystemPreference{
		db:      db,
		uspRepo: uspRepo,
	}
}

func (s *SystemPreference) GetUserSystemPreference(ctx context.Context, userID uint32) (*UserSystemPreference, error) {
	exists, err := s.uspRepo.GetUserSystemPreferenceExists(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, app.NewAppError(app.CodeNotFound, errors.New("user system preference hasn't been created yet"))
	}

	usp, err := s.uspRepo.GetUserSystemPreference(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	return usp, nil
}

func (s *SystemPreference) CreateDefaultUserSystemPreference(ctx context.Context, userID uint32) (*UserSystemPreference, error) {
	_, err := s.uspRepo.CreateUserSystemPreference(ctx, s.db, CreateUserSystemPreferenceParams{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	usp, err := s.uspRepo.GetUserSystemPreference(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	return usp, nil
}

func (s *SystemPreference) UpdateDefaultUserSystemPreference(ctx context.Context, userID uint32, language *string) (*UserSystemPreference, error) {
	params := UpdateUserSystemPreferenceParams{
		UserID:   userID,
		Language: language,
	}

	err := s.uspRepo.UpdateUserSystemPreference(ctx, s.db, params)
	if err != nil {
		return nil, err
	}

	usp, err := s.uspRepo.GetUserSystemPreference(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	return usp, nil
}
