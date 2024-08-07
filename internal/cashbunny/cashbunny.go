package cashbunny

import (
	"context"
	"database/sql"

	"github.com/massivebugs/home-feature-server/db/service/cashbunny_repository"
)

type Cashbunny struct {
	db          *sql.DB
	accountRepo cashbunny_repository.Querier
}

func NewCashbunny(db *sql.DB, accountRepo cashbunny_repository.Querier) *Cashbunny {
	return &Cashbunny{
		db:          db,
		accountRepo: accountRepo,
	}
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
		cashbunny_repository.IncrementIndexParams{
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
		cashbunny_repository.CreateAccountParams{
			UserID:      userID,
			Category:    req.Category,
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
	data, err := s.accountRepo.ListAccounts(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	accounts := make([]*Account, len(data))
	for idx, d := range data {
		a, err := NewAccount(d)
		if err != nil {
			return nil, err
		}
		accounts[idx] = a
	}

	return accounts, nil
}

func (s *Cashbunny) DeleteAccount(ctx context.Context, userID uint32, accountID uint32) error {
	return s.accountRepo.DeleteAccount(ctx, s.db, cashbunny_repository.DeleteAccountParams{
		UserID: userID,
		ID:     accountID,
	})
}

// func (s *Cashbunny) CreateTransaction(ctx context.Context, userID uint32) ()
