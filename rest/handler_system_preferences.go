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
		UserSystemPreference: oapi.UserSystemPreference{
			Language: result.Language,
		},
	}, nil
}

func (h *SystemPreferencesHandler) CreateDefaultUserSystemPreference(ctx context.Context, request oapi.CreateDefaultUserSystemPreferenceRequestObject) (oapi.CreateDefaultUserSystemPreferenceResponseObject, error) {
	claims := h.GetClaims(ctx)

	result, err := h.sp.CreateDefaultUserSystemPreference(ctx, claims.UserID)
	if err != nil {
		return nil, err
	}

	return oapi.CreateDefaultUserSystemPreference200JSONResponse{
		UserSystemPreference: oapi.UserSystemPreference{
			Language: result.Language,
		},
	}, nil
}

func (h *SystemPreferencesHandler) UpdateUserSystemPreference(ctx context.Context, request oapi.UpdateUserSystemPreferenceRequestObject) (oapi.UpdateUserSystemPreferenceResponseObject, error) {
	claims := h.GetClaims(ctx)

	usp, err := h.sp.UpdateDefaultUserSystemPreference(ctx, claims.UserID, request.Body.Language)
	if err != nil {
		return nil, err
	}

	return oapi.UpdateUserSystemPreference200JSONResponse{
		UserSystemPreference: oapi.UserSystemPreference{
			Language: usp.Language,
		},
	}, nil
}
