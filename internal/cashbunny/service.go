package cashbunny

import (
	"context"
	"errors"
	"time"

	// mapset "github.com/deckarep/golang-set/v2"
	"github.com/Rhymond/go-money"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/internal/app"
	"github.com/massivebugs/home-feature-server/internal/util"
)

type Cashbunny struct {
	db        *db.Handle
	accRepo   IAccountRepository
	strRepo   IScheduledTransactionRepository
	trRepo    ITransactionRepository
	trcRepo   ITransactionCategoryRepository
	rruleRepo IRecurrenceRuleRepository
	ucrRepo   ICurrencyRepository
	upRepo    IUserPreferenceRepository
}

func NewCashbunny(
	db *db.Handle,
	accRepo IAccountRepository,
	strRepo IScheduledTransactionRepository,
	trRepo ITransactionRepository,
	trcRepo ITransactionCategoryRepository,
	rruleRepo IRecurrenceRuleRepository,
	ucrRepo ICurrencyRepository,
	upRepo IUserPreferenceRepository,
) *Cashbunny {
	return &Cashbunny{
		db:        db,
		accRepo:   accRepo,
		strRepo:   strRepo,
		trRepo:    trRepo,
		trcRepo:   trcRepo,
		rruleRepo: rruleRepo,
		ucrRepo:   ucrRepo,
		upRepo:    upRepo,
	}
}

func (s *Cashbunny) GetUserPreference(ctx context.Context, userID uint32) (*UserPreference, error) {
	exists, err := s.upRepo.GetUserPreferenceExistsByUserID(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, app.NewAppError(app.CodeNotFound, errors.New("user preferences hasn't been created yet"))
	}

	up, err := s.upRepo.GetUserPreferencesByUserID(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	return up, nil
}

func (s *Cashbunny) CreateDefaultUserPreferences(ctx context.Context, userID uint32) (*UserPreference, error) {
	var up *UserPreference
	err := s.db.WithTx(ctx, func(tx db.DB) error {
		_, err := s.upRepo.CreateUserPreferences(ctx, tx, userID)
		if err != nil {
			return err
		}

		_, err = s.upRepo.CreateUserCurrency(ctx, tx, CreateUserCurrencyParams{
			UserID:       userID,
			CurrencyCode: money.CAD,
		})
		if err != nil {
			return err
		}

		_, err = s.upRepo.CreateUserCurrency(ctx, tx, CreateUserCurrencyParams{
			UserID:       userID,
			CurrencyCode: money.JPY,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	up, err = s.upRepo.GetUserPreferencesByUserID(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	return up, nil
}

func (s *Cashbunny) GetSupportedCurrencies(ctx context.Context) map[string]string {
	currencies := map[string]string{}
	for _, code := range supportedCurrencyCodes {
		gmCurrency := money.GetCurrency(code)
		currencies[gmCurrency.Code] = gmCurrency.Grapheme
	}

	return currencies
}

func (s *Cashbunny) GetOverview(ctx context.Context, userID uint32, from time.Time, to time.Time) (overview, error) {
	accounts, err := s.accRepo.ListAccountsAndAmountBetweenDates(ctx, s.db, ListAccountsAndAmountBetweenDatesParams{
		UserID:           userID,
		FromTransactedAt: time.Time{},
		ToTransactedAt:   to,
	})
	if err != nil {
		return overview{}, err
	}

	trs, err := s.trRepo.ListTransactionsBetweenDates(
		ctx,
		s.db,
		ListTransactionsBetweenDatesParams{
			UserID:           userID,
			FromTransactedAt: time.Time{},
			ToTransactedAt:   to,
		},
	)
	if err != nil {
		return overview{}, err
	}

	strs, err := s.strRepo.ListScheduledTransactionsWithAllRelations(ctx, s.db, userID)
	if err != nil {
		return overview{}, err
	}

	var tsFromScheduled []*Transaction
	for _, str := range strs {
		tsFromScheduled = append(tsFromScheduled, str.toTransactions(from, to)...)
	}

	return newOverview(from, to, NewLedger(accounts, trs), tsFromScheduled), nil
}

func (s *Cashbunny) ListAccounts(ctx context.Context, userID uint32, now time.Time) ([]*Account, error) {
	accs, err := s.accRepo.ListAccountsAndAmountBetweenDates(ctx, s.db, ListAccountsAndAmountBetweenDatesParams{
		UserID:           userID,
		FromTransactedAt: time.Time{},
		ToTransactedAt:   now,
	})
	if err != nil {
		return nil, err
	}

	return accs, nil
}

func (s *Cashbunny) CreateAccount(
	ctx context.Context,
	userID uint32,
	name string,
	category string,
	description string,
	currency string,
	orderIndex *uint32,
) error {
	return s.db.WithTx(ctx, func(tx db.DB) error {
		if orderIndex != nil {
			err := s.accRepo.IncrementAccountIndices(
				ctx,
				tx,
				IncrementAccountIndicesParams{
					UserID:     userID,
					OrderIndex: *orderIndex,
				},
			)
			if err != nil {
				return err
			}
		}

		_, err := s.accRepo.CreateAccount(ctx, tx,
			CreateAccountParams{
				UserID:      userID,
				Category:    AccountCategory(category),
				Name:        name,
				Description: description,
				Currency:    currency,
				OrderIndex:  orderIndex,
			},
		)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *Cashbunny) UpdateAccount(
	ctx context.Context,
	userID uint32,
	accountID uint32,
	args struct {
		Name        string
		Description string
		OrderIndex  uint32
	}) error {
	return s.db.WithTx(ctx, func(tx db.DB) error {
		return s.accRepo.UpdateAccount(ctx, tx, UpdateAccountParams{
			UserID:      userID,
			ID:          accountID,
			Name:        args.Name,
			Description: args.Description,
			OrderIndex:  args.OrderIndex,
		})
	})
}

func (s *Cashbunny) DeleteAccount(ctx context.Context, userID uint32, accountID uint32) error {
	return s.db.WithTx(ctx, func(tx db.DB) error {
		err := s.accRepo.DeleteAccount(ctx, tx, DeleteAccountParams{
			UserID: userID,
			ID:     accountID,
		})
		if err != nil {
			return err
		}

		err = s.trRepo.DeleteTransactionsByAccountID(ctx, tx, DeleteTransactionsByAccountIDParams{
			UserID:    userID,
			AccountID: accountID,
		})
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *Cashbunny) ListTransactions(ctx context.Context, userID uint32) ([]*Transaction, error) {
	trs, err := s.trRepo.ListTransactions(ctx, s.db, userID)
	if err != nil {
		return nil, err
	}

	// Preload accounts
	accIds := util.NewSet[uint32]()
	for _, tr := range trs {
		accIds.Add(tr.SrcAccountID)
		accIds.Add(tr.DestAccountID)
	}

	accs, err := s.accRepo.ListAccountsByIDs(ctx, s.db, ListAccountsByIDsParams{
		UserID: userID,
		IDs:    accIds.ToSlice(),
	})
	if err != nil {
		return nil, err
	}

	// Map preloaded Accounts to Transactions
	for _, tr := range trs {
		srcAcc := util.SliceFind(accs, func(a *Account) bool { return a.ID == tr.SrcAccountID })
		if srcAcc == nil {
			return nil, err
		}

		destAcc := util.SliceFind(accs, func(a *Account) bool { return a.ID == tr.DestAccountID })
		if destAcc == nil {
			return nil, err
		}

		tr.SourceAccount = *srcAcc
		tr.DestinationAccount = *destAcc
	}

	return trs, nil
}

func (s *Cashbunny) CreateTransaction(
	ctx context.Context,
	userID uint32,
	args struct {
		Description          string
		Amount               float64
		Currency             string
		SourceAccountID      uint32
		DestinationAccountID uint32
		TransactedAt         time.Time
	},
) error {
	// Check if source account belong to this user
	sa, err := s.accRepo.GetAccountByID(ctx, s.db, GetAccountByIDParams{
		UserID: userID,
		ID:     args.SourceAccountID,
	})
	if err != nil {
		return err
	}

	// Check if destination account belong to this user
	da, err := s.accRepo.GetAccountByID(ctx, s.db, GetAccountByIDParams{
		UserID: userID,
		ID:     args.DestinationAccountID,
	})
	if err != nil {
		return err
	}

	if sa.Currency != args.Currency || da.Currency != args.Currency {
		return app.NewAppError(app.CodeBadRequest, errors.New("currency does not match the currency of either account"))
	}

	// TODO: Validate 0 balance

	_, err = s.trRepo.CreateTransaction(ctx, s.db, CreateTransactionParams{
		UserID:        userID,
		SrcAccountID:  args.SourceAccountID,
		DestAccountID: args.DestinationAccountID,
		Description:   args.Description,
		Amount:        args.Amount,
		Currency:      args.Currency,
		TransactedAt:  args.TransactedAt,
	})

	return err
}

func (s *Cashbunny) UpdateTransaction(
	ctx context.Context,
	userID uint32,
	transactionID uint32,
	args struct {
		Description  string
		Amount       float64
		TransactedAt time.Time
	}) error {
	return s.db.WithTx(ctx, func(tx db.DB) error {
		return s.trRepo.UpdateTransaction(ctx, tx, UpdateTransactionParams{
			UserID:       userID,
			ID:           transactionID,
			Description:  args.Description,
			Amount:       args.Amount,
			TransactedAt: args.TransactedAt,
		})
	})
}

func (s *Cashbunny) DeleteTransaction(ctx context.Context, userID uint32, transactionID uint32) error {
	// TODO: Validate 0 balance

	return s.trRepo.DeleteTransaction(ctx, s.db, DeleteTransactionParams{
		UserID: userID,
		ID:     transactionID,
	})
}

// func (s *Cashbunny) GetPlan(ctx context.Context, userID uint32) (planResponse, error) {
// 	return newPlanResponse(&Planner{}), nil
// }

// func (s *Cashbunny) GetPlannerParameters(
// 	ctx context.Context,
// 	userID uint32,
// ) (
// 	plannerParametersResponse,
// 	error,
// ) {
// 	assetAccounts, err := s.accRepo.ListAccountsAndAmountByCategory(ctx, s.db, ListAccountsAndAmountByCategoryParams{
// 		UserID:   userID,
// 		Category: AccountCategoryAssets,
// 	})
// 	if err != nil {
// 		return plannerParametersResponse{}, err
// 	}

// 	strs, err := s.strRepo.ListScheduledTransactionsWithAllRelations(ctx, s.db, userID)
// 	if err != nil {
// 		return plannerParametersResponse{}, err
// 	}

// 	revenueStrs := []*ScheduledTransaction{}
// 	liabilityStrs := []*ScheduledTransaction{}
// 	expenseStrs := []*ScheduledTransaction{}
// 	for _, str := range strs {
// 		switch {
// 		case str.isRevenueTransaction():
// 			revenueStrs = append(revenueStrs, str)
// 		case str.isLiabilityTransaction():
// 			liabilityStrs = append(liabilityStrs, str)
// 		case str.isExpenseTransaction():
// 			expenseStrs = append(expenseStrs, str)
// 		}
// 	}

// 	trCategories, err := s.trcRepo.ListTransactionCategoriesWithAllocations(ctx, s.db, userID)
// 	if err != nil {
// 		return plannerParametersResponse{}, err
// 	}

// 	result := newPlannerParametersResponse(assetAccounts, revenueStrs, liabilityStrs, expenseStrs, trCategories)
// 	return result, nil
// }

// func (s *Cashbunny) SavePlannerParameters(ctx context.Context, userID uint32, p *SavePlannerParametersRequest) (planResponse, error) {
// 	return newPlanResponse(&Planner{}), nil
// }
