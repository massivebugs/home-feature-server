package cashbunny

import (
	"context"
	"database/sql"

	"github.com/massivebugs/home-feature-server/db/service/cashbunny_account"
	"github.com/massivebugs/home-feature-server/db/service/cashbunny_category"
)

type Cashbunny struct {
	db           *sql.DB
	categoryRepo cashbunny_category.Querier
	accountRepo  cashbunny_account.Querier
}

func NewCashbunny(db *sql.DB, categoryRepo cashbunny_category.Querier, accountRepo cashbunny_account.Querier) *Cashbunny {
	return &Cashbunny{
		db:           db,
		categoryRepo: categoryRepo,
		accountRepo:  accountRepo,
	}
}

func (s *Cashbunny) CreateCategory(ctx context.Context, userID uint32, req *CreateCategoryRequestDTO) (*cashbunny_category.CashbunnyCategory, error) {
	result, err := s.categoryRepo.CreateCategory(
		ctx,
		s.db,
		cashbunny_category.CreateCategoryParams{
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

	c, err := s.categoryRepo.GetCategoryByID(
		ctx,
		s.db,
		cashbunny_category.GetCategoryByIDParams{
			UserID: userID,
			ID:     uint32(id),
		},
	)
	if err != nil {
		return nil, err
	}

	return c, err
}

func (s *Cashbunny) ListCategories(ctx context.Context, userID uint32) ([]*cashbunny_category.CashbunnyCategory, error) {
	result, err := s.categoryRepo.ListCategoriesByUserID(
		ctx,
		s.db,
		userID,
	)

	return result, err
}

func (s *Cashbunny) CreateAccount(ctx context.Context, userID uint32, req *CreateAccountRequestDTO) error {
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

	_, err = s.accountRepo.CreateAccount(
		ctx,
		tx,
		cashbunny_account.CreateAccountParams{
			UserID:      userID,
			CategoryID:  req.CategoryID,
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

func (s *Cashbunny) ListAccounts(ctx context.Context, userID uint32) ([]Account, error) {
	allAccountData, err := s.accountRepo.ListAccounts(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	accounts := make([]Account, len(allAccountData))
	for idx, d := range allAccountData {
		a, err := NewAccount(d)
		if err != nil {
			return nil, err
		}
		accounts[idx] = a
	}

	return accounts, nil
}
