package system_preference

import (
	"context"

	"github.com/massivebugs/home-feature-server/db"
)

// ### User System Preference ###
type CreateUserSystemPreferenceParams struct {
	UserID   uint32
	Language *string
}

type UpdateUserSystemPreferenceParams struct {
	UserID   uint32
	Language *string
}

type ISystemPreferenceRepository interface {
	CreateUserSystemPreference(ctx context.Context, db db.DB, arg CreateUserSystemPreferenceParams) (uint32, error)
	GetUserSystemPreferenceExists(ctx context.Context, db db.DB, userID uint32) (bool, error)
	GetUserSystemPreference(ctx context.Context, db db.DB, userID uint32) (*UserSystemPreference, error)
	UpdateUserSystemPreference(ctx context.Context, db db.DB, arg UpdateUserSystemPreferenceParams) error
}
