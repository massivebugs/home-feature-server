package cashbunny

import (
	"context"
	"database/sql"

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
