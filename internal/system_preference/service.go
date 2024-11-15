package system_preference

import (
	"context"
	"database/sql"
	"errors"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/internal/app"
)

type SystemPreference struct {
	db      *db.Handle
	uspRepo ISystemPreferenceRepository
}

func NewSystemPreference(
	db *db.Handle,
	uspRepo ISystemPreferenceRepository,
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
		return nil, app.NewAppError(app.CodeNotFound, errors.New("user system preferences hasn't been created yet"))
	}

	usp, err := s.uspRepo.GetUserSystemPreference(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	return usp, nil
}

func (s *SystemPreference) CreateDefaultUserSystemPreference(ctx context.Context, userID uint32) (userSystemPreferenceResponse, error) {
	_, err := s.uspRepo.CreateUserSystemPreference(ctx, s.db, CreateUserSystemPreferenceParams{
		UserID:   userID,
		Language: sql.NullString{Valid: false},
	})
	if err != nil {
		return userSystemPreferenceResponse{}, err
	}

	usp, err := s.uspRepo.GetUserSystemPreference(ctx, s.db, userID)
	if err != nil {
		return userSystemPreferenceResponse{}, err
	}

	return newUserSystemPreferenceResponse(usp), nil
}

func (s *SystemPreference) UpdateDefaultUserSystemPreference(ctx context.Context, userID uint32, req *UserSystemPreferenceDTO) error {
	params := UpdateUserSystemPreferenceParams{
		UserID: userID,
	}

	if req.Language != nil {
		params.Language = sql.NullString{Valid: true, String: *req.Language}
	}

	return s.uspRepo.UpdateUserSystemPreference(ctx, s.db, params)
}
