package repository

import (
	"context"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
)

type AccountRepository struct {
	querier queries.Querier
}

var _ cashbunny.IAccountRepository = (*AccountRepository)(nil)

func NewAccountRepository(querier queries.Querier) *AccountRepository {
	return &AccountRepository{
		querier: querier,
	}
}

func (r *AccountRepository) CreateAccount(ctx context.Context, db db.DB, params cashbunny.CreateAccountParams) (uint32, error) {
	result, err := r.querier.CreateAccount(ctx, db, queries.CreateAccountParams{
		UserID:      params.UserID,
		Category:    string(params.Category),
		Name:        params.Name,
		Description: params.Description,
		Currency:    params.Currency,
		OrderIndex:  params.OrderIndex,
	})
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

func (r *AccountRepository) ListAccountsAndAmountBetweenDates(ctx context.Context, db db.DB, params cashbunny.ListAccountsAndAmountBetweenDatesParams) ([]*cashbunny.Account, error) {
	data, err := r.querier.ListAccountsAndAmountBetweenDates(ctx, db, queries.ListAccountsAndAmountBetweenDatesParams{
		UserID:           params.UserID,
		FromTransactedAt: params.FromTransactedAt,
		ToTransactedAt:   params.ToTransactedAt,
	})
	if err != nil {
		return nil, err
	}

	accounts := make([]*cashbunny.Account, len(data))
	for idx, d := range data {
		accounts[idx] = cashbunny.NewAccountFromQueries(&d.CashbunnyAccount, &d.Amount)
	}

	return accounts, nil
}

func (r *AccountRepository) ListAccountsAndAmountByCategory(ctx context.Context, db db.DB, params cashbunny.ListAccountsAndAmountByCategoryParams) ([]*cashbunny.Account, error) {
	data, err := r.querier.ListAccountsAndAmountByCategory(ctx, db, queries.ListAccountsAndAmountByCategoryParams{
		UserID:   params.UserID,
		Category: string(params.Category),
	})
	if err != nil {
		return nil, err
	}

	assetAccs := make([]*cashbunny.Account, len(data))
	for idx, d := range data {
		assetAccs[idx] = cashbunny.NewAccountFromQueries(&d.CashbunnyAccount, &d.Amount)
	}

	return assetAccs, nil
}

func (r *AccountRepository) ListAccountsByIDs(ctx context.Context, db db.DB, params cashbunny.ListAccountsByIDsParams) ([]*cashbunny.Account, error) {
	data, err := r.querier.ListAccountsByIDs(ctx, db, queries.ListAccountsByIDsParams{
		UserID: params.UserID,
		IDs:    params.IDs,
	})
	if err != nil {
		return nil, err
	}

	accs := make([]*cashbunny.Account, len(data))
	for idx, d := range data {
		accs[idx] = cashbunny.NewAccountFromQueries(d, nil)
	}

	return accs, nil
}

func (r *AccountRepository) IncrementAccountIndices(ctx context.Context, db db.DB, params cashbunny.IncrementAccountIndicesParams) error {
	err := r.querier.IncrementAccountIndices(ctx, db, queries.IncrementAccountIndicesParams{
		UserID:     params.UserID,
		OrderIndex: params.OrderIndex,
	},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) DeleteAccount(ctx context.Context, db db.DB, params cashbunny.DeleteAccountParams) error {
	return r.querier.DeleteAccount(ctx, db, queries.DeleteAccountParams{
		UserID: params.UserID,
		ID:     params.ID,
	})
}

func (r *AccountRepository) GetAccountByID(ctx context.Context, db db.DB, params cashbunny.GetAccountByIDParams) (*cashbunny.Account, error) {
	data, err := r.querier.GetAccountByID(ctx, db, queries.GetAccountByIDParams{
		UserID: params.UserID,
		ID:     params.ID,
	})
	if err != nil {
		return nil, err
	}

	return cashbunny.NewAccountFromQueries(data, nil), nil
}
