package cashbunny

import (
	"context"
	"database/sql"
	"time"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
)

type CreateAccountParams struct {
	UserID      uint32
	Category    AccountCategory
	Name        string
	Description string
	Currency    string
	OrderIndex  *uint32
}

type UpdateAccountParams struct {
	Name        string
	Description string
	OrderIndex  uint32
	UserID      uint32
	ID          uint32
}

type ListAccountsAndAmountBetweenDatesParams struct {
	FromTransactedAt time.Time
	ToTransactedAt   time.Time
	UserID           uint32
}

type ListAccountsAndAmountByCategoryParams struct {
	UserID   uint32
	Category AccountCategory
}

type ListAccountsByIDsParams struct {
	UserID uint32
	IDs    []uint32
}

type DeleteAccountParams struct {
	UserID uint32
	ID     uint32
}

type GetAccountByIDParams struct {
	UserID uint32
	ID     uint32
}

type IncrementAccountIndicesParams struct {
	UserID     uint32
	OrderIndex uint32
}

type IAccountRepository interface {
	CreateAccount(ctx context.Context, db db.DB, params CreateAccountParams) (uint32, error)
	UpdateAccount(ctx context.Context, db db.DB, arg UpdateAccountParams) error
	ListAccountsAndAmountBetweenDates(ctx context.Context, db db.DB, params ListAccountsAndAmountBetweenDatesParams) ([]*Account, error)
	ListAccountsAndAmountByCategory(ctx context.Context, db db.DB, params ListAccountsAndAmountByCategoryParams) ([]*Account, error)
	ListAccountsByIDs(ctx context.Context, db db.DB, params ListAccountsByIDsParams) ([]*Account, error)
	IncrementAccountIndices(ctx context.Context, db db.DB, params IncrementAccountIndicesParams) error
	DeleteAccount(ctx context.Context, db db.DB, params DeleteAccountParams) error
	GetAccountByID(ctx context.Context, db db.DB, params GetAccountByIDParams) (*Account, error)
}

type AccountRepository struct {
	querier queries.Querier
}

var _ IAccountRepository = (*AccountRepository)(nil)

func NewAccountRepository(querier queries.Querier) *AccountRepository {
	return &AccountRepository{
		querier: querier,
	}
}

func (r *AccountRepository) CreateAccount(ctx context.Context, db db.DB, params CreateAccountParams) (uint32, error) {
	p := queries.CreateAccountParams{
		UserID:      params.UserID,
		Category:    string(params.Category),
		Name:        params.Name,
		Description: params.Description,
		Currency:    params.Currency,
	}

	if params.OrderIndex != nil {
		p.OrderIndex = sql.NullInt32{
			Valid: true,
			Int32: int32(*params.OrderIndex),
		}
	}

	result, err := r.querier.CreateAccount(ctx, db, p)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(id), nil
}

func (r *AccountRepository) UpdateAccount(ctx context.Context, db db.DB, arg UpdateAccountParams) error {
	p := queries.UpdateCashbunnyAccountParams{
		UserID:      arg.UserID,
		ID:          arg.ID,
		Name:        arg.Name,
		Description: arg.Description,
		OrderIndex:  arg.OrderIndex,
	}

	return r.querier.UpdateCashbunnyAccount(ctx, db, p)
}

func (r *AccountRepository) ListAccountsAndAmountBetweenDates(ctx context.Context, db db.DB, params ListAccountsAndAmountBetweenDatesParams) ([]*Account, error) {
	data, err := r.querier.ListAccountsAndAmountBetweenDates(ctx, db, queries.ListAccountsAndAmountBetweenDatesParams{
		UserID:           params.UserID,
		FromTransactedAt: params.FromTransactedAt,
		ToTransactedAt:   params.ToTransactedAt,
	})
	if err != nil {
		return nil, err
	}

	accounts := make([]*Account, len(data))
	for idx, d := range data {
		accounts[idx] = NewAccountFromQueries(&d.CashbunnyAccount, &d.Amount)
	}

	return accounts, nil
}

func (r *AccountRepository) ListAccountsAndAmountByCategory(ctx context.Context, db db.DB, params ListAccountsAndAmountByCategoryParams) ([]*Account, error) {
	data, err := r.querier.ListAccountsAndAmountByCategory(ctx, db, queries.ListAccountsAndAmountByCategoryParams{
		UserID:   params.UserID,
		Category: string(params.Category),
	})
	if err != nil {
		return nil, err
	}

	assetAccs := make([]*Account, len(data))
	for idx, d := range data {
		assetAccs[idx] = NewAccountFromQueries(&d.CashbunnyAccount, &d.Amount)
	}

	return assetAccs, nil
}

func (r *AccountRepository) ListAccountsByIDs(ctx context.Context, db db.DB, params ListAccountsByIDsParams) ([]*Account, error) {
	data, err := r.querier.ListAccountsByIDs(ctx, db, queries.ListAccountsByIDsParams{
		UserID: params.UserID,
		IDs:    params.IDs,
	})
	if err != nil {
		return nil, err
	}

	accs := make([]*Account, len(data))
	for idx, d := range data {
		accs[idx] = NewAccountFromQueries(d, nil)
	}

	return accs, nil
}

func (r *AccountRepository) IncrementAccountIndices(ctx context.Context, db db.DB, params IncrementAccountIndicesParams) error {
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

func (r *AccountRepository) DeleteAccount(ctx context.Context, db db.DB, params DeleteAccountParams) error {
	return r.querier.DeleteAccount(ctx, db, queries.DeleteAccountParams{
		UserID: params.UserID,
		ID:     params.ID,
	})
}

func (r *AccountRepository) GetAccountByID(ctx context.Context, db db.DB, params GetAccountByIDParams) (*Account, error) {
	data, err := r.querier.GetAccountByID(ctx, db, queries.GetAccountByIDParams{
		UserID: params.UserID,
		ID:     params.ID,
	})
	if err != nil {
		return nil, err
	}

	return NewAccountFromQueries(data, nil), nil
}
