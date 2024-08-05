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

func (s *Cashbunny) CreateCategory(ctx context.Context, userID uint32, req *CreateCategoryRequestDTO) (*cashbunny_account.CashbunnyCategory, error) {
	result, err := s.accountRepo.CreateCategory(
		ctx,
		s.db,
		cashbunny_account.CreateCategoryParams{
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

	c, err := s.accountRepo.GetCategoryByID(
		ctx,
		s.db,
		cashbunny_account.GetCategoryByIDParams{
			UserID: userID,
			ID:     uint32(id),
		},
	)
	if err != nil {
		return nil, err
	}

	return c, err
}

func (s *Cashbunny) ListCategories(ctx context.Context, userID uint32) ([]*cashbunny_account.CashbunnyCategory, error) {
	result, err := s.accountRepo.ListCategoriesByUserID(
		ctx,
		s.db,
		userID,
	)

	return result, err
}

func (s *Cashbunny) CreateAccount(ctx context.Context, userID uint32, req *CreateAccountRequestDTO) error {
	// Check if category exists
	category, err := s.accountRepo.GetCategoryByName(
		ctx,
		s.db,
		cashbunny_account.GetCategoryByNameParams{
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
		result, err := s.accountRepo.CreateCategory(
			ctx,
			tx,
			cashbunny_account.CreateCategoryParams{
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
		a, err := NewAccount(&d.CashbunnyAccount, &d.CashbunnyCategory)
		if err != nil {
			return nil, err
		}
		accounts[idx] = a
	}

	return accounts, nil
}

// func (s *Cashbunny) CreateTransaction(ctx context.Context, userID uint32) ()
