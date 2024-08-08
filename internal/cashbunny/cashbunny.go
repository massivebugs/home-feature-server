package cashbunny

import (
	"context"
	"database/sql"
	"errors"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/massivebugs/home-feature-server/db/service/cashbunny_repository"
)

type Cashbunny struct {
	db            *sql.DB
	cashbunnyRepo cashbunny_repository.Querier
}

func NewCashbunny(db *sql.DB, cashbunnyRepo cashbunny_repository.Querier) *Cashbunny {
	return &Cashbunny{
		db:            db,
		cashbunnyRepo: cashbunnyRepo,
	}
}

func (s *Cashbunny) CreateAccount(ctx context.Context, userID uint32, req *CreateAccountRequestDTO) error {
	accountType, err := GetAccountTypeForCategory(AccountCategory(req.Category))
	if err != nil {
		// TODO: Flag errors
		return err
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = s.cashbunnyRepo.IncrementIndex(
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

	_, err = s.cashbunnyRepo.CreateAccount(
		ctx,
		tx,
		cashbunny_repository.CreateAccountParams{
			UserID:      userID,
			Category:    req.Category,
			Name:        req.Name,
			Description: req.Description,
			Balance:     req.Balance,
			Currency:    req.Currency,
			Type:        string(accountType),
			OrderIndex:  req.OrderIndex,
		},
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Cashbunny) ListAccounts(ctx context.Context, userID uint32) ([]*Account, error) {
	data, err := s.cashbunnyRepo.ListAccounts(ctx, s.db, userID)
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
	return s.cashbunnyRepo.DeleteAccount(ctx, s.db, cashbunny_repository.DeleteAccountParams{
		UserID: userID,
		ID:     accountID,
	})
}

func (s *Cashbunny) CreateTransaction(ctx context.Context, userID uint32, req *CreateTransactionRequestDTO) error {
	// Check if source account belong to this user
	_, err := s.cashbunnyRepo.GetAccountByID(ctx, s.db, cashbunny_repository.GetAccountByIDParams{
		UserID: userID,
		ID:     req.SourceAccountID,
	})
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		// TODO: flag action
		return err
	} else if err != nil {
		return err
	}

	// Check if destination account belong to this user
	_, err = s.cashbunnyRepo.GetAccountByID(ctx, s.db, cashbunny_repository.GetAccountByIDParams{
		UserID: userID,
		ID:     req.DestinationAccountID,
	})
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		// TODO: flag action
		return err
	} else if err != nil {
		return err
	}

	transactedAt, err := time.Parse(time.DateTime, req.TransactedAt)
	if err != nil {
		return err
	}

	_, err = s.cashbunnyRepo.CreateTransaction(ctx, s.db, cashbunny_repository.CreateTransactionParams{
		UserID:        userID,
		SrcAccountID:  req.SourceAccountID,
		DestAccountID: req.DestinationAccountID,
		Description:   req.Description,
		Amount:        req.Amount,
		Currency:      req.Currency,
		TransactedAt:  transactedAt,
	})

	return err
}

func (s *Cashbunny) ListTransactions(ctx context.Context, userID uint32) ([]*Transaction, error) {
	data, err := s.cashbunnyRepo.ListTransactions(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	// Preload accounts
	accountIds := mapset.NewSet[uint32]()
	for _, d := range data {
		accountIds.Add(d.SrcAccountID)
		accountIds.Add(d.DestAccountID)
	}

	accounts, err := s.cashbunnyRepo.ListAccountsByIDs(ctx, s.db, cashbunny_repository.ListAccountsByIDsParams{
		UserID: userID,
		IDs:    accountIds.ToSlice(),
	})
	if err != nil {
		return nil, err
	}

	trs := make([]*Transaction, len(data))
	for idx, d := range data {
		srcAccount := SliceFind(accounts, func(a *cashbunny_repository.CashbunnyAccount) bool { return a.ID == d.SrcAccountID })
		destAccount := SliceFind(accounts, func(a *cashbunny_repository.CashbunnyAccount) bool { return a.ID == d.DestAccountID })

		a, err := NewTransaction(d, *srcAccount, *destAccount)
		if err != nil {
			return nil, err
		}

		trs[idx] = a
	}

	return trs, nil
}

func (s *Cashbunny) DeleteTransaction(ctx context.Context, userID uint32, transactionID uint32) error {
	return s.cashbunnyRepo.DeleteTransaction(ctx, s.db, cashbunny_repository.DeleteTransactionParams{
		UserID: userID,
		ID:     transactionID,
	})
}
