package seeder

import (
	"context"
	"time"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
)

type Seeder struct {
	db      *db.Handle
	querier queries.Querier
}

func NewSeeder(
	db *db.Handle,
	querier queries.Querier,
) *Seeder {
	return &Seeder{
		db:      db,
		querier: querier,
	}
}

func (s *Seeder) SeedForLocal(ctx context.Context) error {
	now := time.Now()

	return s.db.WithTx(ctx, func(tx db.DB) error {
		userID, err := s.createAuthDataForPublicUser(ctx, tx)
		if err != nil {
			return err
		}

		err = s.createUserSystemPreferenceDataForPublicUser(ctx, tx, userID)
		if err != nil {
			return err
		}

		err = s.createCashbunnyDataForPublicUser(ctx, tx, userID, now)
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *Seeder) SeedForProduction(ctx context.Context) error {
	return s.db.WithTx(ctx, func(tx db.DB) error {
		return nil
	})
}
