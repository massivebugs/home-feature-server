package handler

import (
	"context"

	"github.com/massivebugs/home-feature-server/db"
	"github.com/massivebugs/home-feature-server/db/queries"
	"github.com/massivebugs/home-feature-server/internal/system_preference"
	"github.com/massivebugs/home-feature-server/rest"
	"github.com/massivebugs/home-feature-server/rest/oapi"
)

type SystemPreferenceHandler struct {
	*rest.Handler
	sp *system_preference.SystemPreference
}

func NewSystemPreferenceHandler(cfg *rest.Config, db *db.Handle, querier queries.Querier) *SystemPreferenceHandler {
	return &SystemPreferenceHandler{
		Handler: rest.NewHandler(cfg),
		sp: system_preference.NewSystemPreference(
			db,
			system_preference.NewUserSystemPreferenceRepository(querier),
		),
	}
}

func (h *SystemPreferenceHandler) GetUserSystemPreference(ctx context.Context, request oapi.GetUserSystemPreferenceRequestObject) (oapi.GetUserSystemPreferenceResponseObject, error) {
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

func (h *SystemPreferenceHandler) CreateDefaultUserSystemPreference(ctx context.Context, request oapi.CreateDefaultUserSystemPreferenceRequestObject) (oapi.CreateDefaultUserSystemPreferenceResponseObject, error) {
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

func (h *SystemPreferenceHandler) UpdateUserSystemPreference(ctx context.Context, request oapi.UpdateUserSystemPreferenceRequestObject) (oapi.UpdateUserSystemPreferenceResponseObject, error) {
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
