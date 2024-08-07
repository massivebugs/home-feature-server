package handler

import (
	"database/sql"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api/response"
	"github.com/massivebugs/home-feature-server/db/service/cashbunny_repository"
	"github.com/massivebugs/home-feature-server/internal/api"
	"github.com/massivebugs/home-feature-server/internal/auth"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
)

type CashbunnyHandler struct {
	cashbunny *cashbunny.Cashbunny
}

func NewCashbunnyHandler(db *sql.DB) *CashbunnyHandler {
	return &CashbunnyHandler{
		cashbunny: cashbunny.NewCashbunny(
			db,
			cashbunny_repository.New(),
		),
	}
}

func (h *CashbunnyHandler) CreateAccount(ctx echo.Context) *api.APIResponse {
	token := ctx.Get("user").(*jwt.Token)
	claims := token.Claims.(*auth.JWTClaims)

	req := new(cashbunny.CreateAccountRequestDTO)

	if err := ctx.Bind(req); err != nil {
		return api.NewAPIResponse(err, "")
	}

	if err := ctx.Validate(req); err != nil {
		return api.NewAPIResponse(err, "")
	}

	err := h.cashbunny.CreateAccount(ctx.Request().Context(), claims.UserID, req)

	return api.NewAPIResponse(err, nil)
}

func (h *CashbunnyHandler) ListAccounts(ctx echo.Context) *api.APIResponse {
	token := ctx.Get("user").(*jwt.Token)
	claims := token.Claims.(*auth.JWTClaims)

	result, err := h.cashbunny.ListAccounts(ctx.Request().Context(), claims.UserID)

	return api.NewAPIResponse(err, response.NewListAccountsResponseDTO(result))
}

func (h *CashbunnyHandler) DeleteAccount(ctx echo.Context) *api.APIResponse {
	token := ctx.Get("user").(*jwt.Token)
	claims := token.Claims.(*auth.JWTClaims)

	accountID, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	if err != nil {
		return api.NewAPIResponse(err, "")
	}

	err = h.cashbunny.DeleteAccount(ctx.Request().Context(), claims.UserID, uint32(accountID))

	return api.NewAPIResponse(err, nil)

}

func (h *CashbunnyHandler) CreateTransaction(ctx echo.Context) *api.APIResponse {
	token := ctx.Get("user").(*jwt.Token)
	claims := token.Claims.(*auth.JWTClaims)

	req := new(cashbunny.CreateTransactionRequestDTO)

	if err := ctx.Bind(req); err != nil {
		return api.NewAPIResponse(err, "")
	}

	if err := ctx.Validate(req); err != nil {
		return api.NewAPIResponse(err, "")
	}

	err := h.cashbunny.CreateTransaction(ctx.Request().Context(), claims.UserID, req)

	return api.NewAPIResponse(err, nil)
}

func (h *CashbunnyHandler) ListTransactions(ctx echo.Context) *api.APIResponse {
	token := ctx.Get("user").(*jwt.Token)
	claims := token.Claims.(*auth.JWTClaims)

	result, err := h.cashbunny.ListTransactions(ctx.Request().Context(), claims.UserID)

	return api.NewAPIResponse(err, response.NewListTransactionsResponseDTO(result))
}
