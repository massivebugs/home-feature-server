package seeder

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/massivebugs/home-feature-server/api/config"
	"github.com/massivebugs/home-feature-server/db/service/auth_repository"
	"github.com/massivebugs/home-feature-server/db/service/cashbunny_repository"
	"github.com/massivebugs/home-feature-server/internal/auth"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
)

type Seeder struct {
	db            *sql.DB
	cfg           *config.Config
	authRepo      auth_repository.Querier
	cashbunnyRepo cashbunny_repository.Querier
}

func NewSeeder(
	db *sql.DB,
	cfg *config.Config,
	authRepo auth_repository.Querier,
	cashbunnyRepo cashbunny_repository.Querier,
) *Seeder {
	return &Seeder{
		db:            db,
		cfg:           cfg,
		authRepo:      authRepo,
		cashbunnyRepo: cashbunnyRepo,
	}
}

func (s *Seeder) Seed(ctx context.Context) error {
	fmt.Printf("Running seeder for environment: %s", s.cfg.Environment)
	fmt.Println()

	switch s.cfg.Environment {
	case config.EnvironmentLocal:
		return s.seedForLocal(ctx)
	case config.EnvironmentProduction:
		return s.seedForProduction(ctx)

	}

	return fmt.Errorf("found no seeders for specified environment: %s", s.cfg.Environment)
}

func (s *Seeder) seedForLocal(ctx context.Context) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Create new user
	createUserResult, err := s.authRepo.CreateUser(ctx, tx, "testuser")
	if err != nil {
		return err
	}

	// Retrieve ID
	userID, err := createUserResult.LastInsertId()
	if err != nil {
		return err
	}

	// Hash password
	hashedPassword, err := auth.GeneratePasswordHash("test_password_123")
	if err != nil {
		return err
	}

	// Create user password
	p := auth_repository.CreateUserPasswordParams{
		UserID:       uint32(userID),
		PasswordHash: hashedPassword,
	}
	_, err = s.authRepo.CreateUserPassword(ctx, tx, p)
	if err != nil {
		return err
	}

	// Cashbunny - create test debit account
	createAccountResult1, err := s.cashbunnyRepo.CreateAccount(
		ctx,
		tx,
		cashbunny_repository.CreateAccountParams{
			UserID:      uint32(userID),
			Category:    string(cashbunny.AccountCategoryAssets),
			Name:        "Foo Bank account #1",
			Description: "For storing income and for every day use",
			Balance:     1486000,
			Currency:    "JPY",
			Type:        string(cashbunny.AccountTypeDebit),
			OrderIndex:  0,
		},
	)
	if err != nil {
		return err
	}

	// Retrieve ID
	account1ID, err := createAccountResult1.LastInsertId()
	if err != nil {
		return err
	}

	// Cashbunny - create test debit account
	createAccountResult2, err := s.cashbunnyRepo.CreateAccount(
		ctx,
		tx,
		cashbunny_repository.CreateAccountParams{
			UserID:      uint32(userID),
			Category:    string(cashbunny.AccountCategoryExpenses),
			Name:        "Groceries",
			Description: "For grocery expenses",
			Balance:     10000,
			Currency:    "JPY",
			Type:        string(cashbunny.AccountTypeDebit),
			OrderIndex:  1,
		},
	)
	if err != nil {
		return err
	}

	// Retrieve ID
	account2ID, err := createAccountResult2.LastInsertId()
	if err != nil {
		return err
	}

	// Cashbunny - create test credit account
	createAccountResult3, err := s.cashbunnyRepo.CreateAccount(
		ctx,
		tx,
		cashbunny_repository.CreateAccountParams{
			UserID:      uint32(userID),
			Category:    string(cashbunny.AccountCategoryRevenues),
			Name:        "Bar Inc.",
			Description: "My workplace",
			Balance:     1496000,
			Currency:    "JPY",
			Type:        string(cashbunny.AccountTypeCredit),
			OrderIndex:  2,
		},
	)
	if err != nil {
		return err
	}

	// Retrieve ID
	account3ID, err := createAccountResult3.LastInsertId()
	if err != nil {
		return err
	}

	// Cashbunny - create test transaction
	_, err = s.cashbunnyRepo.CreateTransaction(ctx, tx, cashbunny_repository.CreateTransactionParams{
		UserID:        uint32(userID),
		SrcAccountID:  uint32(account3ID),
		DestAccountID: uint32(account1ID),
		Description:   "Monthly wage for May",
		Amount:        710000,
		Currency:      "JPY",
		TransactedAt:  time.Date(2024, 6, 15, 10, 0, 0, 0, time.Local),
	})
	if err != nil {
		return err
	}

	// Cashbunny - create test transaction
	_, err = s.cashbunnyRepo.CreateTransaction(ctx, tx, cashbunny_repository.CreateTransactionParams{
		UserID:        uint32(userID),
		SrcAccountID:  uint32(account1ID),
		DestAccountID: uint32(account2ID),
		Description:   "Bought some groceries for 2 weeks",
		Amount:        10000,
		Currency:      "JPY",
		TransactedAt:  time.Date(2024, 6, 28, 15, 0, 0, 0, time.Local),
	})
	if err != nil {
		return err
	}

	// Cashbunny - create test transaction
	_, err = s.cashbunnyRepo.CreateTransaction(ctx, tx, cashbunny_repository.CreateTransactionParams{
		UserID:        uint32(userID),
		SrcAccountID:  uint32(account3ID),
		DestAccountID: uint32(account1ID),
		Description:   "Monthly wage for June",
		Amount:        786000,
		Currency:      "JPY",
		TransactedAt:  time.Date(2024, 7, 15, 10, 0, 0, 0, time.Local),
	})
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Seeder) seedForProduction(_ context.Context) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// TODO: Write inserts here

	return tx.Commit()
}
