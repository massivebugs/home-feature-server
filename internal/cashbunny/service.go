package cashbunny

import (
	"context"
	"errors"

	// mapset "github.com/deckarep/golang-set/v2"
	"github.com/Rhymond/go-money"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/internal/app"
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

// func (s *Cashbunny) GetOverview(ctx context.Context, userID uint32, from time.Time, to time.Time) (overviewResponse, error) {
// 	accounts, err := s.accRepo.ListAccountsAndAmountBetweenDates(ctx, s.db, ListAccountsAndAmountBetweenDatesParams{
// 		UserID:           userID,
// 		FromTransactedAt: time.Time{},
// 		ToTransactedAt:   to,
// 	})
// 	if err != nil {
// 		return overviewResponse{}, err
// 	}

// 	trs, err := s.trRepo.ListTransactionsBetweenDates(
// 		ctx,
// 		s.db,
// 		ListTransactionsBetweenDatesParams{
// 			UserID:           userID,
// 			FromTransactedAt: time.Time{},
// 			ToTransactedAt:   to,
// 		},
// 	)
// 	if err != nil {
// 		return overviewResponse{}, err
// 	}

// 	strs, err := s.strRepo.ListScheduledTransactionsWithAllRelations(ctx, s.db, userID)
// 	if err != nil {
// 		return overviewResponse{}, err
// 	}

// 	var tsFromScheduled []*Transaction
// 	for _, str := range strs {
// 		tsFromScheduled = append(tsFromScheduled, str.toTransactions(from, to)...)
// 	}

// 	return newOverviewResponse(from, to, NewLedger(accounts, trs), tsFromScheduled), nil
// }

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

// func (s *Cashbunny) CreateAccount(ctx context.Context, userID uint32, req *CreateAccountRequest) error {
// 	return s.db.WithTx(ctx, func(tx db.DB) error {
// 		orderIndexParam := sql.NullInt32{}
// 		if req.OrderIndex != nil {
// 			err := s.accRepo.IncrementAccountIndices(
// 				ctx,
// 				tx,
// 				IncrementAccountIndicesParams{
// 					UserID:     userID,
// 					OrderIndex: *req.OrderIndex,
// 				},
// 			)
// 			if err != nil {
// 				return err
// 			}

// 			orderIndexParam.Valid = true
// 			orderIndexParam.Int32 = int32(*req.OrderIndex)
// 		}

// 		_, err := s.accRepo.CreateAccount(ctx, tx,
// 			CreateAccountParams{
// 				UserID:      userID,
// 				Category:    AccountCategory(req.Category),
// 				Name:        req.Name,
// 				Description: req.Description,
// 				Currency:    req.Currency,
// 				OrderIndex:  orderIndexParam,
// 			},
// 		)
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	})
// }

// func (s *Cashbunny) ListAccounts(ctx context.Context, userID uint32, now time.Time) ([]accountResponse, error) {
// 	accounts, err := s.accRepo.ListAccountsAndAmountBetweenDates(ctx, s.db, ListAccountsAndAmountBetweenDatesParams{
// 		UserID:           userID,
// 		FromTransactedAt: time.Time{},
// 		ToTransactedAt:   now,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return newListAccountsResponse(accounts), nil
// }

// func (s *Cashbunny) DeleteAccount(ctx context.Context, userID uint32, accountID uint32) error {
// 	return s.db.WithTx(ctx, func(tx db.DB) error {
// 		err := s.accRepo.DeleteAccount(ctx, tx, DeleteAccountParams{
// 			UserID: userID,
// 			ID:     accountID,
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		err = s.trRepo.DeleteTransactionsByAccountID(ctx, tx, DeleteTransactionsByAccountIDParams{
// 			UserID:    userID,
// 			AccountID: accountID,
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	})
// }

// func (s *Cashbunny) CreateTransaction(ctx context.Context, userID uint32, req *CreateTransactionRequest) error {
// 	// Check if source account belong to this user
// 	sa, err := s.accRepo.GetAccountByID(ctx, s.db, GetAccountByIDParams{
// 		UserID: userID,
// 		ID:     req.SourceAccountID,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	// Check if destination account belong to this user
// 	da, err := s.accRepo.GetAccountByID(ctx, s.db, GetAccountByIDParams{
// 		UserID: userID,
// 		ID:     req.DestinationAccountID,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	if sa.currency != req.Currency || da.currency != req.Currency {
// 		return app.NewAppError(app.CodeBadRequest, errors.New("currency does not match the currency of either account"))
// 	}

// 	// TODO: Validate 0 balance

// 	transactedAt, err := time.Parse(time.RFC3339Nano, req.TransactedAt)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = s.trRepo.CreateTransaction(ctx, s.db, CreateTransactionParams{
// 		UserID:        userID,
// 		SrcAccountID:  req.SourceAccountID,
// 		DestAccountID: req.DestinationAccountID,
// 		Description:   req.Description,
// 		Amount:        req.Amount,
// 		Currency:      req.Currency,
// 		TransactedAt:  transactedAt,
// 	})

// 	return err
// }

// func (s *Cashbunny) ListTransactions(ctx context.Context, userID uint32) ([]transactionResponse, error) {
// 	trs, err := s.trRepo.ListTransactions(ctx, s.db, userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Preload accounts
// 	accIds := mapset.NewSet[uint32]()
// 	for _, tr := range trs {
// 		accIds.Add(tr.srcAccountID)
// 		accIds.Add(tr.destAccountID)
// 	}

// 	accs, err := s.accRepo.ListAccountsByIDs(ctx, s.db, ListAccountsByIDsParams{
// 		UserID: userID,
// 		IDs:    accIds.ToSlice(),
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Map preloaded Accounts to Transactions
// 	for _, tr := range trs {
// 		srcAcc := util.SliceFind(accs, func(a *Account) bool { return a.id == tr.srcAccountID })
// 		if srcAcc == nil {
// 			return nil, err
// 		}

// 		destAcc := util.SliceFind(accs, func(a *Account) bool { return a.id == tr.destAccountID })
// 		if destAcc == nil {
// 			return nil, err
// 		}

// 		tr.sourceAccount = *srcAcc
// 		tr.destinationAccount = *destAcc
// 	}

// 	return newListTransactionsResponse(trs), nil
// }

// func (s *Cashbunny) DeleteTransaction(ctx context.Context, userID uint32, transactionID uint32) error {
// 	tr, err := s.trRepo.GetTransactionByID(ctx, s.db, GetTransactionByIDParams{
// 		UserID: userID,
// 		ID:     transactionID,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	// Check if source account belong to this user
// 	_, err = s.accRepo.GetAccountByID(ctx, s.db, GetAccountByIDParams{
// 		UserID: userID,
// 		ID:     tr.srcAccountID,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	// Check if destination account belong to this user
// 	_, err = s.accRepo.GetAccountByID(ctx, s.db, GetAccountByIDParams{
// 		UserID: userID,
// 		ID:     tr.destAccountID,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	// TODO: Validate 0 balance

// 	return s.trRepo.DeleteTransaction(ctx, s.db, DeleteTransactionParams{
// 		UserID: userID,
// 		ID:     transactionID,
// 	})
// }
