package seeder

import (
	"context"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
)

func (s *Seeder) createUserSystemPreferenceDataForPublicUser(ctx context.Context, tx db.DB, userID uint32) error {
	// Create new user system preferences
	_, err := s.querier.CreateUserSystemPreference(ctx, tx, queries.CreateUserSystemPreferenceParams{
		UserID: userID,
	})

	return err
}
