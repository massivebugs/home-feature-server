package seeder

import (
	"context"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/auth"
)

func (s *Seeder) createAuthDataForPublicUser(ctx context.Context, tx db.DB) (uint32, error) {
	// Create new user
	createUserResult, err := s.querier.CreateUser(ctx, tx, "public")
	if err != nil {
		return 0, err
	}
	userID, err := createUserResult.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Hash password
	hashedPassword, err := auth.GeneratePasswordHash("")
	if err != nil {
		return 0, err
	}

	// Create user password
	p := queries.CreateUserPasswordParams{
		UserID:       uint32(userID),
		PasswordHash: hashedPassword,
	}
	_, err = s.querier.CreateUserPassword(ctx, tx, p)
	if err != nil {
		return 0, err
	}

	return uint32(userID), err
}
