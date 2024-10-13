package system_preference

import (
	"context"
	"database/sql"

	"github.com/massivebugs/home-feature-server/db"
)

// ### User System Preference ###
type CreateUserSystemPreferenceParams struct {
	UserID   uint32
	Language sql.NullString
}

type UpdateUserSystemPreferenceParams struct {
	Language sql.NullString
	UserID   uint32
}

type ISystemPreferenceRepository interface {
	CreateUserSystemPreference(ctx context.Context, db db.DB, arg CreateUserSystemPreferenceParams) (uint32, error)
	GetUserSystemPreferenceExists(ctx context.Context, db db.DB, userID uint32) (bool, error)
	GetUserSystemPreference(ctx context.Context, db db.DB, userID uint32) (*UserSystemPreference, error)
	UpdateUserSystemPreference(ctx context.Context, db db.DB, arg UpdateUserSystemPreferenceParams) error
}
