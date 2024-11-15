package handler

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
	"github.com/massivebugs/home-feature-server/rest"
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

func (h *CashbunnyHandler) GetOverview(c echo.Context) error {
	claims := h.GetTokenClaims(c)

	qFrom, _ := strconv.Atoi(c.QueryParam("from"))

	qTo, _ := strconv.Atoi(c.QueryParam("to"))

	var from time.Time
	var to time.Time

	if qFrom == 0 {
		from = time.Time{}
	} else {
		from = time.Unix(int64(qFrom), 0)
	}

	if qTo == 0 {
		to = time.Now()
	} else {
		to = time.Unix(int64(qTo), 0)
	}

	result, err := h.cashbunny.GetOverview(c.Request().Context(), claims.UserID, from, to)
	if err != nil {
		return h.CreateErrorResponse(c, err)
	}

	return h.CreateResponse(c, nil, result)
}

func (h *CashbunnyHandler) GetPlan(c echo.Context) error {
	claims := h.GetTokenClaims(c)

	result, err := h.cashbunny.GetPlan(c.Request().Context(), claims.UserID)
	if err != nil {
		return h.CreateErrorResponse(c, err)
	}

	return h.CreateResponse(c, nil, result)
}

func (h *CashbunnyHandler) GetPlannerParameters(c echo.Context) error {
	claims := h.GetTokenClaims(c)

	result, err := h.cashbunny.GetPlannerParameters(c.Request().Context(), claims.UserID)
	if err != nil {
		return h.CreateErrorResponse(c, err)
	}

	return h.CreateResponse(c, nil, result)
}

func (h *CashbunnyHandler) SavePlannerParameters(c echo.Context) error {
	claims := h.GetTokenClaims(c)

	req := new(cashbunny.SavePlannerParametersRequest)
	err := h.Validate(c, req)
	if err != nil {
		return h.CreateErrorResponse(c, err)
	}

	result, err := h.cashbunny.SavePlannerParameters(c.Request().Context(), claims.UserID, req)
	if err != nil {
		return h.CreateErrorResponse(c, err)
	}

	return h.CreateResponse(c, nil, result)
}

func (h *CashbunnyHandler) GetCurrencies(c echo.Context) error {
	result := h.cashbunny.GetAllCurrencies(c.Request().Context())

	return h.CreateResponse(c, nil, result)
}

func (h *CashbunnyHandler) GetUserPreferences(c echo.Context) error {
	claims := h.GetTokenClaims(c)

	result, err := h.cashbunny.GetUserPreferences(c.Request().Context(), claims.UserID)
	if err != nil {
		return h.CreateErrorResponse(c, err)
	}

	return h.CreateResponse(c, nil, result)
}

func (h *CashbunnyHandler) CreateDefaultUserPreferences(c echo.Context) error {
	claims := h.GetTokenClaims(c)

	result, err := h.cashbunny.CreateDefaultUserPreferences(c.Request().Context(), claims.UserID)
	if err != nil {
		return h.CreateErrorResponse(c, err)
	}

	return h.CreateResponse(c, nil, result)
}

func (h *CashbunnyHandler) CreateAccount(c echo.Context) error {
	claims := h.GetTokenClaims(c)

	req := new(cashbunny.CreateAccountRequest)

	err := h.Validate(c, req)
	if err != nil {
		return h.CreateErrorResponse(c, err)
	}

	err = h.cashbunny.CreateAccount(c.Request().Context(), claims.UserID, req)

	return h.CreateResponse(c, err, nil)
}

func (h *CashbunnyHandler) ListAccounts(c echo.Context) error {
	claims := h.GetTokenClaims(c)

	result, err := h.cashbunny.ListAccounts(c.Request().Context(), claims.UserID, time.Now())

	return h.CreateResponse(c, err, result)
}

func (h *CashbunnyHandler) DeleteAccount(c echo.Context) error {
	claims := h.GetTokenClaims(c)

	accountID, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return h.CreateErrorResponse(c, err)
	}

	err = h.cashbunny.DeleteAccount(c.Request().Context(), claims.UserID, uint32(accountID))

	return h.CreateResponse(c, err, nil)

}

func (h *CashbunnyHandler) CreateTransaction(c echo.Context) error {
	claims := h.GetTokenClaims(c)

	req := new(cashbunny.CreateTransactionRequest)

	err := h.Validate(c, req)
	if err != nil {
		return h.CreateErrorResponse(c, err)
	}

	err = h.cashbunny.CreateTransaction(c.Request().Context(), claims.UserID, req)

	return h.CreateResponse(c, err, nil)
}

func (h *CashbunnyHandler) ListTransactions(c echo.Context) error {
	claims := h.GetTokenClaims(c)

	result, err := h.cashbunny.ListTransactions(c.Request().Context(), claims.UserID)

	return h.CreateResponse(c, err, result)
}

func (h *CashbunnyHandler) DeleteTransaction(c echo.Context) error {
	claims := h.GetTokenClaims(c)

	transactionId, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return h.CreateErrorResponse(c, err)
	}

	err = h.cashbunny.DeleteTransaction(c.Request().Context(), claims.UserID, uint32(transactionId))

	return h.CreateResponse(c, err, nil)

}
