package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/massivebugs/home-feature-server/api"
	"github.com/massivebugs/home-feature-server/api/response"
	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/system_preference"
	"github.com/massivebugs/home-feature-server/repository"
)

type SystemPreferencesHandler struct {
	*api.Handler
	cfg *api.Config
	sp  *system_preference.SystemPreference
}

func NewSystemPreferencesHandler(cfg *api.Config, db *db.Handle, querier queries.Querier) *SystemPreferencesHandler {
	return &SystemPreferencesHandler{
		cfg: cfg,
		sp: system_preference.NewSystemPreference(
			db,
			repository.NewUserSystemPreferenceRepository(querier),
		),
	}
}

func (h *SystemPreferencesHandler) GetUserSystemPreferences(c echo.Context) *api.APIResponse {
	claims := h.GetTokenClaims(c)

	usp, err := h.sp.GetUserSystemPreference(c.Request().Context(), claims.UserID)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	return h.CreateResponse(nil, response.NewUserSystemPreferenceDTO(usp))
}

func (h *SystemPreferencesHandler) CreateDefaultUserSystemPreferences(c echo.Context) *api.APIResponse {
	claims := h.GetTokenClaims(c)

	usp, err := h.sp.CreateDefaultUserSystemPreference(c.Request().Context(), claims.UserID)
	if err != nil {
		return h.CreateErrorResponse(err)
	}

	return h.CreateResponse(nil, response.NewUserSystemPreferenceDTO(usp))
}

func (h *SystemPreferencesHandler) UpdateDefaultUserSystemPreferences(c echo.Context) *api.APIResponse {
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
