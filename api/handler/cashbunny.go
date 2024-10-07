package handler

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api"
	"github.com/massivebugs/home-feature-server/api/response"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
	"github.com/massivebugs/home-feature-server/repository"
)

type CashbunnyHandler struct {
	*api.Handler
	cashbunny *cashbunny.Cashbunny
}

func NewCashbunnyHandler(db *db.Handle, querier queries.Querier) *CashbunnyHandler {
	return &CashbunnyHandler{
		cashbunny: cashbunny.NewCashbunny(
			db,
			repository.NewAccountDBRepository(querier),
			repository.NewScheduledTransactionDBRepository(querier),
			repository.NewTransactionDBRepository(querier),
			repository.NewRecurrenceRuleDBRepository(querier),
			repository.NewCurrencyDBRepository(querier),
			repository.NewUserPreferencesDBRepository(querier),
		),
	}
}

func (h *CashbunnyHandler) GetOverview(c echo.Context) *api.APIResponse {
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

	ledger, transactionsFromScheduled, err := h.cashbunny.GetOverview(c.Request().Context(), claims.UserID, from, to)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	return h.CreateResponse(nil, response.NewGetOverviewResponseDTO(from, to, ledger, transactionsFromScheduled))
}

func (h *CashbunnyHandler) GetPlan(c echo.Context) *api.APIResponse {
	claims := h.GetTokenClaims(c)

	planner, err := h.cashbunny.GetPlan(c.Request().Context(), claims.UserID)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	return h.CreateResponse(nil, response.NewGetPlanResponseDTO(planner))
}

func (h *CashbunnyHandler) GetPlannerParameters(c echo.Context) *api.APIResponse {
	claims := h.GetTokenClaims(c)

	aas, srts, slts, err := h.cashbunny.GetPlannerParameters(c.Request().Context(), claims.UserID)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	return h.CreateResponse(nil, response.NewGetPlannerParametersResponseDTO(aas, srts, slts))
}

func (h *CashbunnyHandler) SavePlannerParameters(c echo.Context) *api.APIResponse {
	claims := h.GetTokenClaims(c)

	req := new(cashbunny.SavePlannerParametersDTO)
	err := h.Validate(c, req)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	planner, err := h.cashbunny.SavePlannerParameters(c.Request().Context(), claims.UserID, req)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	return h.CreateResponse(nil, response.NewGetPlanResponseDTO(planner))
}

func (h *CashbunnyHandler) GetCurrencies(c echo.Context) *api.APIResponse {
	result := h.cashbunny.GetAllCurrencies(c.Request().Context())

	return h.CreateResponse(nil, response.NewGetAllCurrenciesResponseDTO(result))
}

func (h *CashbunnyHandler) GetUserPreferences(c echo.Context) *api.APIResponse {
	claims := h.GetTokenClaims(c)

	result, err := h.cashbunny.GetUserPreferences(c.Request().Context(), claims.UserID)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	return h.CreateResponse(nil, response.NewGetUserPreferencesDTO(result))
}

func (h *CashbunnyHandler) CreateDefaultUserPreferences(c echo.Context) *api.APIResponse {
	claims := h.GetTokenClaims(c)

	result, err := h.cashbunny.CreateDefaultUserPreferences(c.Request().Context(), claims.UserID)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	return h.CreateResponse(nil, response.NewGetUserPreferencesDTO(result))
}

func (h *CashbunnyHandler) CreateAccount(c echo.Context) *api.APIResponse {
	claims := h.GetTokenClaims(c)

	req := new(cashbunny.CreateAccountDTO)

	err := h.Validate(c, req)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	err = h.cashbunny.CreateAccount(c.Request().Context(), claims.UserID, req)

	return h.CreateResponse(err, nil)
}

func (h *CashbunnyHandler) ListAccounts(c echo.Context) *api.APIResponse {
	claims := h.GetTokenClaims(c)

	result, err := h.cashbunny.ListAccounts(c.Request().Context(), claims.UserID, time.Now())

	return h.CreateResponse(err, response.NewListAccountsResponseDTO(result))
}

func (h *CashbunnyHandler) DeleteAccount(c echo.Context) *api.APIResponse {
	claims := h.GetTokenClaims(c)

	accountID, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	err = h.cashbunny.DeleteAccount(c.Request().Context(), claims.UserID, uint32(accountID))

	return h.CreateResponse(err, nil)

}

func (h *CashbunnyHandler) CreateTransaction(c echo.Context) *api.APIResponse {
	claims := h.GetTokenClaims(c)

	req := new(cashbunny.CreateTransactionDTO)

	err := h.Validate(c, req)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	err = h.cashbunny.CreateTransaction(c.Request().Context(), claims.UserID, req)

	return h.CreateResponse(err, nil)
}

func (h *CashbunnyHandler) ListTransactions(c echo.Context) *api.APIResponse {
	claims := h.GetTokenClaims(c)

	result, err := h.cashbunny.ListTransactions(c.Request().Context(), claims.UserID)

	return h.CreateResponse(err, response.NewListTransactionsResponseDTO(result))
}

func (h *CashbunnyHandler) DeleteTransaction(c echo.Context) *api.APIResponse {
	claims := h.GetTokenClaims(c)

	transactionId, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	err = h.cashbunny.DeleteTransaction(c.Request().Context(), claims.UserID, uint32(transactionId))

	return h.CreateResponse(err, nil)

}
