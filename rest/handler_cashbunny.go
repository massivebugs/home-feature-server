package rest

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
	"github.com/massivebugs/home-feature-server/internal/repository"
)

type CashbunnyHandler struct {
	*Handler
	cashbunny *cashbunny.Cashbunny
}

func NewCashbunnyHandler(db *db.Handle, querier queries.Querier) *CashbunnyHandler {
	return &CashbunnyHandler{
		cashbunny: cashbunny.NewCashbunny(
			db,
			repository.NewAccountDBRepository(querier),
			repository.NewScheduledTransactionDBRepository(querier),
			repository.NewTransactionDBRepository(querier),
			repository.NewTransactionCategoryDBRepository(querier),
			repository.NewRecurrenceRuleDBRepository(querier),
			repository.NewCurrencyDBRepository(querier),
			repository.NewUserPreferencesDBRepository(querier),
		),
	}
}

func (h *CashbunnyHandler) GetOverview(c echo.Context) *APIResponse {
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
		return h.CreateErrorResponse(err)
	}

	return h.CreateResponse(nil, result)
}

func (h *CashbunnyHandler) GetPlan(c echo.Context) *APIResponse {
	claims := h.GetTokenClaims(c)

	result, err := h.cashbunny.GetPlan(c.Request().Context(), claims.UserID)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	return h.CreateResponse(nil, result)
}

func (h *CashbunnyHandler) GetPlannerParameters(c echo.Context) *APIResponse {
	claims := h.GetTokenClaims(c)

	result, err := h.cashbunny.GetPlannerParameters(c.Request().Context(), claims.UserID)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	return h.CreateResponse(nil, result)
}

func (h *CashbunnyHandler) SavePlannerParameters(c echo.Context) *APIResponse {
	claims := h.GetTokenClaims(c)

	req := new(cashbunny.SavePlannerParametersRequest)
	err := h.Validate(c, req)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	result, err := h.cashbunny.SavePlannerParameters(c.Request().Context(), claims.UserID, req)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	return h.CreateResponse(nil, result)
}

func (h *CashbunnyHandler) GetCurrencies(c echo.Context) *APIResponse {
	result := h.cashbunny.GetAllCurrencies(c.Request().Context())

	return h.CreateResponse(nil, result)
}

func (h *CashbunnyHandler) GetUserPreferences(c echo.Context) *APIResponse {
	claims := h.GetTokenClaims(c)

	result, err := h.cashbunny.GetUserPreferences(c.Request().Context(), claims.UserID)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	return h.CreateResponse(nil, result)
}

func (h *CashbunnyHandler) CreateDefaultUserPreferences(c echo.Context) *APIResponse {
	claims := h.GetTokenClaims(c)

	result, err := h.cashbunny.CreateDefaultUserPreferences(c.Request().Context(), claims.UserID)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	return h.CreateResponse(nil, result)
}

func (h *CashbunnyHandler) CreateAccount(c echo.Context) *APIResponse {
	claims := h.GetTokenClaims(c)

	req := new(cashbunny.CreateAccountRequest)

	err := h.Validate(c, req)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	err = h.cashbunny.CreateAccount(c.Request().Context(), claims.UserID, req)

	return h.CreateResponse(err, nil)
}

func (h *CashbunnyHandler) ListAccounts(c echo.Context) *APIResponse {
	claims := h.GetTokenClaims(c)

	result, err := h.cashbunny.ListAccounts(c.Request().Context(), claims.UserID, time.Now())

	return h.CreateResponse(err, result)
}

func (h *CashbunnyHandler) DeleteAccount(c echo.Context) *APIResponse {
	claims := h.GetTokenClaims(c)

	accountID, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	err = h.cashbunny.DeleteAccount(c.Request().Context(), claims.UserID, uint32(accountID))

	return h.CreateResponse(err, nil)

}

func (h *CashbunnyHandler) CreateTransaction(c echo.Context) *APIResponse {
	claims := h.GetTokenClaims(c)

	req := new(cashbunny.CreateTransactionRequest)

	err := h.Validate(c, req)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	err = h.cashbunny.CreateTransaction(c.Request().Context(), claims.UserID, req)

	return h.CreateResponse(err, nil)
}

func (h *CashbunnyHandler) ListTransactions(c echo.Context) *APIResponse {
	claims := h.GetTokenClaims(c)

	result, err := h.cashbunny.ListTransactions(c.Request().Context(), claims.UserID)

	return h.CreateResponse(err, result)
}

func (h *CashbunnyHandler) DeleteTransaction(c echo.Context) *APIResponse {
	claims := h.GetTokenClaims(c)

	transactionId, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	err = h.cashbunny.DeleteTransaction(c.Request().Context(), claims.UserID, uint32(transactionId))

	return h.CreateResponse(err, nil)

}
