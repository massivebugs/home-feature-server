package http

import (
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/repository"
	"github.com/massivebugs/home-feature-server/internal/system_preference"
)

type SystemPreferencesHandler struct {
	*Handler
	cfg *Config
	sp  *system_preference.SystemPreference
}

func NewSystemPreferencesHandler(cfg *Config, db *db.Handle, querier queries.Querier) *SystemPreferencesHandler {
	return &SystemPreferencesHandler{
		cfg: cfg,
		sp: system_preference.NewSystemPreference(
			db,
			repository.NewUserSystemPreferenceRepository(querier),
		),
	}
}

func (h *SystemPreferencesHandler) GetUserSystemPreferences(c echo.Context) *APIResponse {
	claims := h.GetTokenClaims(c)

	result, err := h.sp.GetUserSystemPreference(c.Request().Context(), claims.UserID)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	return h.CreateResponse(nil, result)
}

func (h *SystemPreferencesHandler) CreateDefaultUserSystemPreferences(c echo.Context) *APIResponse {
	claims := h.GetTokenClaims(c)

	result, err := h.sp.CreateDefaultUserSystemPreference(c.Request().Context(), claims.UserID)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	return h.CreateResponse(nil, result)
}

func (h *SystemPreferencesHandler) UpdateDefaultUserSystemPreferences(c echo.Context) *APIResponse {
	claims := h.GetTokenClaims(c)

	req := new(system_preference.UserSystemPreferenceDTO)

	err := h.Validate(c, req)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	err = h.sp.UpdateDefaultUserSystemPreference(c.Request().Context(), claims.UserID, req)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	return h.CreateResponse(nil, nil)
}
