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
		From:                      ov.From.Format(time.RFC3339Nano),
		To:                        ov.To.Format(time.RFC3339Nano),
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

// func (h *CashbunnyHandler) CreateTransaction(c echo.Context) error {
// 	claims := h.GetTokenClaims(c)

// 	req := new(cashbunny.CreateTransactionRequest)

// 	err := h.Validate(c, req)
// 	if err != nil {
// 		return h.CreateErrorResponse(c, err)
// 	}

// 	err = h.cashbunny.CreateTransaction(c.Request().Context(), claims.UserID, req)

// 	return h.CreateResponse(c, err, nil)
// }

// func (h *CashbunnyHandler) ListTransactions(c echo.Context) error {
// 	claims := h.GetTokenClaims(c)

// 	result, err := h.cashbunny.ListTransactions(c.Request().Context(), claims.UserID)

// 	return h.CreateResponse(c, err, result)
// }

// func (h *CashbunnyHandler) DeleteTransaction(c echo.Context) error {
// 	claims := h.GetTokenClaims(c)

// 	transactionId, err := strconv.ParseInt(c.Param("id"), 10, 32)
// 	if err != nil {
// 		return h.CreateErrorResponse(c, err)
// 	}

// 	err = h.cashbunny.DeleteTransaction(c.Request().Context(), claims.UserID, uint32(transactionId))

// 	return h.CreateResponse(c, err, nil)

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
		CreatedAt:     a.CreatedAt.Format(time.RFC3339Nano),
		UpdatedAt:     a.UpdatedAt.Format(time.RFC3339Nano),
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
		TransactedAt:  tr.TransactedAt.Format(time.RFC3339Nano),
		CreatedAt:     tr.CreatedAt.Format(time.RFC3339Nano),
		UpdatedAt:     tr.UpdatedAt.Format(time.RFC3339Nano),

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
			CreatedAt:     tr.ScheduledTransaction.CreatedAt.Format(time.RFC3339Nano),
			UpdatedAt:     tr.ScheduledTransaction.UpdatedAt.Format(time.RFC3339Nano),

			RecurrenceRule: oapi.CashbunnyRecurrenceRule{
				Freq:     tr.ScheduledTransaction.RecurrenceRule.Rule.Options.Freq.String(),
				Dtstart:  tr.ScheduledTransaction.RecurrenceRule.Rule.Options.Dtstart.Format(time.RFC3339Nano),
				Count:    tr.ScheduledTransaction.RecurrenceRule.Rule.OrigOptions.Count,
				Interval: tr.ScheduledTransaction.RecurrenceRule.Rule.Options.Interval,
				Until:    tr.ScheduledTransaction.RecurrenceRule.Rule.Options.Until.Format(time.RFC3339Nano),
			},

			SourceAccountId:        tr.ScheduledTransaction.SourceAccount.ID,
			SourceAccountName:      tr.ScheduledTransaction.SourceAccount.Name,
			DestinationAccountId:   tr.ScheduledTransaction.DestinationAccount.ID,
			DestinationAccountName: tr.ScheduledTransaction.DestinationAccount.Name,
		}
	}

	return e
}
