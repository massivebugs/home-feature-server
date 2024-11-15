package rest

import (
	"context"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/repository"
	"github.com/massivebugs/home-feature-server/internal/system_preference"
	"github.com/massivebugs/home-feature-server/rest/oapi"
)

type SystemPreferencesHandler struct {
	*Handler
	sp *system_preference.SystemPreference
}

func NewSystemPreferencesHandler(cfg *Config, db *db.Handle, querier queries.Querier) *SystemPreferencesHandler {
	return &SystemPreferencesHandler{
		Handler: &Handler{
			cfg: cfg,
		},
		sp: system_preference.NewSystemPreference(
			db,
			repository.NewUserSystemPreferenceRepository(querier),
		),
	}
}

func (h *SystemPreferencesHandler) GetUserSystemPreference(ctx context.Context, request oapi.GetUserSystemPreferenceRequestObject) (oapi.GetUserSystemPreferenceResponseObject, error) {
	claims := h.GetClaims(ctx)

	result, err := h.sp.GetUserSystemPreference(ctx, claims.UserID)
	if err != nil {
		return nil, err
	}

	return oapi.GetUserSystemPreference200JSONResponse{
		Language: result.Language,
	}, nil
}

// func (h *SystemPreferencesHandler) CreateDefaultUserSystemPreference(c echo.Context) error {
// 	claims := h.GetTokenClaims(c)

// 	result, err := h.sp.CreateDefaultUserSystemPreference(c.Request().Context(), claims.UserID)
// 	if err != nil {
// 		return h.CreateErrorResponse(c, err)
// 	}

// 	return h.CreateResponse(c, nil, result)
// }

// func (h *SystemPreferencesHandler) UpdateUserSystemPreference(c echo.Context) error {
// 	claims := h.GetTokenClaims(c)

// 	req := new(system_preference.UserSystemPreferenceDTO)

// 	err := h.Validate(c, req)
// 	if err != nil {
// 		return h.CreateErrorResponse(c, err)
// 	}

// 	err = h.sp.UpdateDefaultUserSystemPreference(c.Request().Context(), claims.UserID, req)
// 	if err != nil {
// 		return h.CreateErrorResponse(c, err)
// 	}

// 	return h.CreateResponse(c, nil, nil)
// }
