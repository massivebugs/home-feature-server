package handler

import (
	"context"
	"time"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
	"github.com/massivebugs/home-feature-server/rest"
	"github.com/massivebugs/home-feature-server/rest/oapi"
)

type CashbunnyHandler struct {
	*rest.Handler
	cashbunny *cashbunny.Cashbunny
}

func NewCashbunnyHandler(cfg *rest.Config, db *db.Handle, querier queries.Querier) *CashbunnyHandler {
	return &CashbunnyHandler{
		Handler: rest.NewHandler(cfg),
		cashbunny: cashbunny.NewCashbunny(
			db,
			cashbunny.NewAccountRepository(querier),
			cashbunny.NewScheduledTransactionRepository(querier),
			cashbunny.NewTransactionRepository(querier),
			cashbunny.NewTransactionCategoryRepository(querier),
			cashbunny.NewRecurrenceRuleRepository(querier),
			cashbunny.NewCurrencyRepository(querier),
			cashbunny.NewUserPreferencesRepository(querier),
		),
	}
}

func (h *CashbunnyHandler) GetCashbunnyUserPreference(ctx context.Context, request oapi.GetCashbunnyUserPreferenceRequestObject) (oapi.GetCashbunnyUserPreferenceResponseObject, error) {
	claims := h.GetClaims(ctx)

	result, err := h.cashbunny.GetUserPreference(ctx, claims.UserID)
	if err != nil {
		return nil, err
	}

	return oapi.GetCashbunnyUserPreference200JSONResponse{
		UserPreference: oapi.CashbunnyUserPreference{
			UserCurrencies: result.UserCurrencies,
		},
	}, nil
}

func (h *CashbunnyHandler) CreateCashbunnyDefaultUserPreference(ctx context.Context, request oapi.CreateCashbunnyDefaultUserPreferenceRequestObject) (oapi.CreateCashbunnyDefaultUserPreferenceResponseObject, error) {
	claims := h.GetClaims(ctx)

	result, err := h.cashbunny.CreateDefaultUserPreferences(ctx, claims.UserID)
	if err != nil {
		return nil, err
	}

	return oapi.CreateCashbunnyDefaultUserPreference200JSONResponse{
		UserPreference: oapi.CashbunnyUserPreference{
			UserCurrencies: result.UserCurrencies,
		},
	}, nil
}

func (h *CashbunnyHandler) GetCashbunnySupportedCurrencies(ctx context.Context, request oapi.GetCashbunnySupportedCurrenciesRequestObject) (oapi.GetCashbunnySupportedCurrenciesResponseObject, error) {
	result := h.cashbunny.GetSupportedCurrencies(ctx)

	return oapi.GetCashbunnySupportedCurrencies200JSONResponse{
		CurrenciesAndGrapheme: result,
	}, nil
}

func (h *CashbunnyHandler) GetCashbunnyOverview(ctx context.Context, request oapi.GetCashbunnyOverviewRequestObject) (oapi.GetCashbunnyOverviewResponseObject, error) {
	claims := h.GetClaims(ctx)

	var from time.Time
	var to time.Time

	if request.Params.From == nil {
		from = time.Time{}
	} else {
		from = time.Unix(*request.Params.From, 0)
	}

	if request.Params.To == nil {
		to = time.Now()
	} else {
		to = time.Unix(*request.Params.To, 0)
	}

	ov, err := h.cashbunny.GetOverview(ctx, claims.UserID, from, to)
	if err != nil {
		return nil, err
	}

	res := oapi.GetCashbunnyOverview200JSONResponse{
		From:                      ov.From.Format(h.Config.APIDateTimeFormat),
		To:                        ov.To.Format(h.Config.APIDateTimeFormat),
		NetWorth:                  ov.NetWorth,
		ProfitLossSummary:         ov.ProfitLossSummary,
		AssetAccounts:             make([]oapi.CashbunnyAccount, len(ov.AssetAccounts)),
		LiabilityAccounts:         make([]oapi.CashbunnyAccount, len(ov.LiabilityAccounts)),
		Transactions:              make([]oapi.CashbunnyTransaction, len(ov.Transactions)),
		TransactionsFromScheduled: make([]oapi.CashbunnyTransaction, len(ov.TransactionsFromScheduled)),
	}

	for idx, a := range ov.AssetAccounts {
		res.AssetAccounts[idx] = h.accountToResponse(a)
	}

	for idx, a := range ov.LiabilityAccounts {
		res.LiabilityAccounts[idx] = h.accountToResponse(a)
	}

	for idx, tr := range ov.Transactions {
		res.Transactions[idx] = h.transactionToResponse(tr)
	}

	for idx, tr := range ov.TransactionsFromScheduled {
		res.TransactionsFromScheduled[idx] = h.transactionToResponse(tr)
	}

	return res, nil
}

func (h *CashbunnyHandler) GetCashbunnyAccounts(ctx context.Context, request oapi.GetCashbunnyAccountsRequestObject) (oapi.GetCashbunnyAccountsResponseObject, error) {
	claims := h.GetClaims(ctx)

	accounts, err := h.cashbunny.ListAccounts(ctx, claims.UserID, time.Now())
	if err != nil {
		return nil, err
	}

	res := oapi.GetCashbunnyAccounts200JSONResponse{
		Accounts: make([]oapi.CashbunnyAccount, len(accounts)),
	}

	for idx, a := range accounts {
		res.Accounts[idx] = h.accountToResponse(a)
	}

	return res, nil
}

func (h *CashbunnyHandler) CreateCashbunnyAccount(ctx context.Context, request oapi.CreateCashbunnyAccountRequestObject) (oapi.CreateCashbunnyAccountResponseObject, error) {
	claims := h.GetClaims(ctx)

	if err := h.cashbunny.CreateAccount(
		ctx,
		claims.UserID,
		request.Body.Name,
		request.Body.Category,
		request.Body.Description,
		request.Body.Currency,
		request.Body.OrderIndex,
	); err != nil {
		return nil, err
	}

	return oapi.CreateCashbunnyAccount200Response{}, nil
}

func (h *CashbunnyHandler) DeleteCashbunnyAccount(ctx context.Context, request oapi.DeleteCashbunnyAccountRequestObject) (oapi.DeleteCashbunnyAccountResponseObject, error) {
	claims := h.GetClaims(ctx)

	if err := h.cashbunny.DeleteAccount(ctx, claims.UserID, request.AccountId); err != nil {
		return nil, err
	}

	return oapi.DeleteCashbunnyAccount200Response{}, nil
}

func (h *CashbunnyHandler) UpdateCashbunnyAccount(ctx context.Context, request oapi.UpdateCashbunnyAccountRequestObject) (oapi.UpdateCashbunnyAccountResponseObject, error) {
	claims := h.GetClaims(ctx)

	if err := h.cashbunny.UpdateAccount(ctx, claims.UserID, request.AccountId, struct {
		Name        string
		Description string
		OrderIndex  uint32
	}{
		Name:        request.Body.Name,
		Description: request.Body.Description,
		OrderIndex:  request.Body.OrderIndex,
	}); err != nil {
		return nil, err
	}

	return oapi.UpdateCashbunnyAccount200Response{}, nil
}

func (h *CashbunnyHandler) GetCashbunnyTransactions(ctx context.Context, request oapi.GetCashbunnyTransactionsRequestObject) (oapi.GetCashbunnyTransactionsResponseObject, error) {
	claims := h.GetClaims(ctx)

	trs, err := h.cashbunny.ListTransactions(ctx, claims.UserID)
	if err != nil {
		return nil, err
	}

	res := oapi.GetCashbunnyTransactions200JSONResponse{
		Transactions: make([]oapi.CashbunnyTransaction, len(trs)),
	}

	for idx, tr := range trs {
		res.Transactions[idx] = h.transactionToResponse(tr)
	}

	return res, nil
}

func (h *CashbunnyHandler) CreateCashbunnyTransaction(ctx context.Context, request oapi.CreateCashbunnyTransactionRequestObject) (oapi.CreateCashbunnyTransactionResponseObject, error) {
	claims := h.GetClaims(ctx)

	transactedAt, err := time.Parse(h.Config.APIDateTimeFormat, request.Body.TransactedAt)
	if err != nil {
		return nil, err
	}

	if err := h.cashbunny.CreateTransaction(ctx, claims.UserID, struct {
		Description          string
		Amount               float64
		Currency             string
		SourceAccountID      uint32
		DestinationAccountID uint32
		TransactedAt         time.Time
	}{
		Description:          request.Body.Description,
		Amount:               request.Body.Amount,
		Currency:             request.Body.Currency,
		SourceAccountID:      request.Body.SourceAccountId,
		DestinationAccountID: request.Body.DestinationAccountId,
		TransactedAt:         transactedAt,
	}); err != nil {
		return nil, err
	}

	return oapi.CreateCashbunnyTransaction200Response{}, nil
}

func (h *CashbunnyHandler) UpdateCashbunnyTransaction(ctx context.Context, request oapi.UpdateCashbunnyTransactionRequestObject) (oapi.UpdateCashbunnyTransactionResponseObject, error) {
	claims := h.GetClaims(ctx)

	transactedAt, err := time.Parse(h.Config.APIDateTimeFormat, request.Body.TransactedAt)
	if err != nil {
		return nil, err
	}

	if err := h.cashbunny.UpdateTransaction(ctx, claims.UserID, request.TransactionId, struct {
		Description  string
		Amount       float64
		TransactedAt time.Time
	}{
		Description:  request.Body.Description,
		Amount:       request.Body.Amount,
		TransactedAt: transactedAt,
	}); err != nil {
		return nil, err
	}

	return oapi.UpdateCashbunnyTransaction200Response{}, nil
}

func (h *CashbunnyHandler) DeleteCashbunnyTransaction(ctx context.Context, request oapi.DeleteCashbunnyTransactionRequestObject) (oapi.DeleteCashbunnyTransactionResponseObject, error) {
	claims := h.GetClaims(ctx)

	if err := h.cashbunny.DeleteTransaction(ctx, claims.UserID, request.TransactionId); err != nil {
		return nil, err
	}

	return oapi.DeleteCashbunnyTransaction200Response{}, nil
}

// func (h *CashbunnyHandler) GetPlan(c echo.Context) error {
// 	claims := h.GetTokenClaims(c)

// 	result, err := h.cashbunny.GetPlan(c.Request().Context(), claims.UserID)
// 	if err != nil {
// 		return h.CreateErrorResponse(c, err)
// 	}

// 	return h.CreateResponse(c, nil, result)
// }

// func (h *CashbunnyHandler) GetPlannerParameters(c echo.Context) error {
// 	claims := h.GetTokenClaims(c)

// 	result, err := h.cashbunny.GetPlannerParameters(c.Request().Context(), claims.UserID)
// 	if err != nil {
// 		return h.CreateErrorResponse(c, err)
// 	}

// 	return h.CreateResponse(c, nil, result)
// }

// func (h *CashbunnyHandler) SavePlannerParameters(c echo.Context) error {
// 	claims := h.GetTokenClaims(c)

// 	req := new(cashbunny.SavePlannerParametersRequest)
// 	err := h.Validate(c, req)
// 	if err != nil {
// 		return h.CreateErrorResponse(c, err)
// 	}

// 	result, err := h.cashbunny.SavePlannerParameters(c.Request().Context(), claims.UserID, req)
// 	if err != nil {
// 		return h.CreateErrorResponse(c, err)
// 	}

// 	return h.CreateResponse(c, nil, result)
// }

func (h *CashbunnyHandler) accountToResponse(a *cashbunny.Account) oapi.CashbunnyAccount {
	amount := a.Amount.AsMajorUnits()
	amountDisplay := a.Amount.Display()

	return oapi.CashbunnyAccount{
		Id:            a.ID,
		Category:      string(a.Category),
		Name:          a.Name,
		Description:   a.Description,
		Currency:      a.Currency,
		Type:          string(a.GetType()),
		OrderIndex:    a.OrderIndex,
		CreatedAt:     a.CreatedAt.Format(h.Config.APIDateTimeFormat),
		UpdatedAt:     a.UpdatedAt.Format(h.Config.APIDateTimeFormat),
		Amount:        &amount,
		AmountDisplay: &amountDisplay,
	}
}

func (h *CashbunnyHandler) transactionToResponse(tr *cashbunny.Transaction) oapi.CashbunnyTransaction {
	e := oapi.CashbunnyTransaction{
		Id:            tr.ID,
		Description:   tr.Description,
		Amount:        tr.Amount.AsMajorUnits(),
		Currency:      tr.Amount.Currency().Code,
		AmountDisplay: tr.Amount.Display(),
		TransactedAt:  tr.TransactedAt.Format(h.Config.APIDateTimeFormat),
		CreatedAt:     tr.CreatedAt.Format(h.Config.APIDateTimeFormat),
		UpdatedAt:     tr.UpdatedAt.Format(h.Config.APIDateTimeFormat),

		SourceAccountId:        tr.SourceAccount.ID,
		SourceAccountName:      tr.SourceAccount.Name,
		DestinationAccountId:   tr.DestinationAccount.ID,
		DestinationAccountName: tr.DestinationAccount.Name,
	}

	if tr.ScheduledTransaction != nil {
		e.ScheduledTransaction = oapi.CashbunnyScheduledTransaction{
			Id:            tr.ScheduledTransaction.ID,
			Description:   tr.ScheduledTransaction.Description,
			Amount:        tr.ScheduledTransaction.Amount.AsMajorUnits(),
			Currency:      tr.ScheduledTransaction.Amount.Currency().Code,
			AmountDisplay: tr.ScheduledTransaction.Amount.Display(),
			CreatedAt:     tr.ScheduledTransaction.CreatedAt.Format(h.Config.APIDateTimeFormat),
			UpdatedAt:     tr.ScheduledTransaction.UpdatedAt.Format(h.Config.APIDateTimeFormat),

			RecurrenceRule: oapi.CashbunnyRecurrenceRule{
				Freq:     tr.ScheduledTransaction.RecurrenceRule.Rule.Options.Freq.String(),
				Dtstart:  tr.ScheduledTransaction.RecurrenceRule.Rule.Options.Dtstart.Format(h.Config.APIDateTimeFormat),
				Count:    tr.ScheduledTransaction.RecurrenceRule.Rule.OrigOptions.Count,
				Interval: tr.ScheduledTransaction.RecurrenceRule.Rule.Options.Interval,
				Until:    tr.ScheduledTransaction.RecurrenceRule.Rule.Options.Until.Format(h.Config.APIDateTimeFormat),
			},

			SourceAccountId:        tr.ScheduledTransaction.SourceAccount.ID,
			SourceAccountName:      tr.ScheduledTransaction.SourceAccount.Name,
			DestinationAccountId:   tr.ScheduledTransaction.DestinationAccount.ID,
			DestinationAccountName: tr.ScheduledTransaction.DestinationAccount.Name,
		}
	}

	return e
}
