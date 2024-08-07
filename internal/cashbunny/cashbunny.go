package cashbunny

import (
	"context"
	"database/sql"
	"errors"

	"github.com/massivebugs/home-feature-server/db/service/cashbunny_account"
)

type Cashbunny struct {
	db          *sql.DB
	accountRepo cashbunny_account.Querier
}

func NewCashbunny(db *sql.DB, accountRepo cashbunny_account.Querier) *Cashbunny {
	return &Cashbunny{
		db:          db,
		accountRepo: accountRepo,
	}
}

func (s *Cashbunny) CreateAccountCategory(ctx context.Context, userID uint32, req *CreateAccountCategoryRequestDTO) (*cashbunny_account.CashbunnyAccountCategory, error) {
	result, err := s.accountRepo.CreateAccountCategory(
		ctx,
		s.db,
		cashbunny_account.CreateAccountCategoryParams{
			UserID:      userID,
			Name:        req.Name,
			Description: req.Description,
		},
	)
	if err != nil {
		return nil, err
	}

	// Retrieve ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	c, err := s.accountRepo.GetAccountCategoryByID(
		ctx,
		s.db,
		cashbunny_account.GetAccountCategoryByIDParams{
			UserID: userID,
			ID:     uint32(id),
		},
	)
	if err != nil {
		return nil, err
	}

	return c, err
}

func (s *Cashbunny) ListAccountCategories(ctx context.Context, userID uint32) ([]*cashbunny_account.CashbunnyAccountCategory, error) {
	result, err := s.accountRepo.ListAccountCategoriesByUserID(
		ctx,
		s.db,
		userID,
	)

	return result, err
}

func (s *Cashbunny) CreateAccount(ctx context.Context, userID uint32, req *CreateAccountRequestDTO) error {
	// Check if category exists
	category, err := s.accountRepo.GetAccountCategoryByName(
		ctx,
		s.db,
		cashbunny_account.GetAccountCategoryByNameParams{
			UserID: userID,
			Name:   req.CategoryName,
		},
	)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	var categoryID uint32
	if category != nil {
		categoryID = category.ID
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = s.accountRepo.IncrementIndex(
		ctx,
		tx,
		cashbunny_account.IncrementIndexParams{
			UserID:     userID,
			OrderIndex: req.OrderIndex,
		},
	)
	if err != nil {
		return err
	}

	if categoryID == 0 {
		result, err := s.accountRepo.CreateAccountCategory(
			ctx,
			tx,
			cashbunny_account.CreateAccountCategoryParams{
				UserID:      userID,
				Name:        req.CategoryName,
				Description: req.CategoryName,
			},
		)
		if err != nil {
			return err
		}

		lastInsertID, err := result.LastInsertId()
		if err != nil {
			return err
		}

		categoryID = uint32(lastInsertID)
	}

	_, err = s.accountRepo.CreateAccount(
		ctx,
		tx,
		cashbunny_account.CreateAccountParams{
			UserID:      userID,
			CategoryID:  categoryID,
			Name:        req.Name,
			Description: req.Description,
			Balance:     req.Balance,
			Currency:    req.Currency,
			Type:        req.Type,
			OrderIndex:  req.OrderIndex,
		},
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Cashbunny) ListAccounts(ctx context.Context, userID uint32) ([]*Account, error) {
	data, err := s.accountRepo.ListAccountsAndCategories(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	accounts := make([]*Account, len(data))
	for idx, d := range data {
		a, err := NewAccount(&d.CashbunnyAccount, &d.CashbunnyAccountCategory)
		if err != nil {
			return nil, err
		}
		accounts[idx] = a
	}

	return accounts, nil
}

func (s *Cashbunny) DeleteAccount(ctx context.Context, userID uint32, accountID uint32) error {
	return s.accountRepo.DeleteAccount(ctx, s.db, cashbunny_account.DeleteAccountParams{
		UserID: userID,
		ID:     accountID,
	})
}

// func (s *Cashbunny) CreateTransaction(ctx context.Context, userID uint32) ()
