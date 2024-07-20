package handler

import (
	"database/sql"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/internal/api"
	"github.com/massivebugs/home-feature-server/internal/auth"
	"github.com/massivebugs/home-feature-server/internal/cashbunny"
)

type CashbunnyHandler struct {
	db        *sql.DB
	cashbunny *cashbunny.Cashbunny
}

func NewCashbunnyHandler(db *sql.DB) *CashbunnyHandler {
	return &CashbunnyHandler{
		db:        db,
		cashbunny: cashbunny.NewCashbunny(),
	}
}

func (h *CashbunnyHandler) ListAccounts(ctx echo.Context) *api.APIResponse {
	token := ctx.Get("user").(*jwt.Token)
	claims := token.Claims.(*auth.JWTClaims)

	err := h.cashbunny.ListAccounts(ctx.Request().Context(), claims.UserID)

	return api.NewAPIResponse(err, nil)
}
