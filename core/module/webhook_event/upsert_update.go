package webhook_event

import (
	"context"
	"errors"

	"github.com/mklfarha/mfconsult/core/module/webhook_event/types"
	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"
)

func (m *module) Update(
	ctx context.Context,
	req types.UpsertRequest,
	opts ...Option,
) (types.UpsertResponse, error) {
	optConfig := applyAllOptions(opts)

	tx := optConfig.SQLTx
	createdTx := false
	if tx == nil {
		ntx, err := m.repository.DB.Begin()
		if err != nil {
			return types.UpsertResponse{}, err
		}
		tx = ntx
		defer tx.Rollback()
		createdTx = true
	}

	qtx := m.repository.Queries.WithTx(tx)
	existing, err := qtx.FetchWebhookEventByIdForUpdate(ctx,
		req.WebhookEvent.ID,
	)
	if err != nil {

		return types.UpsertResponse{}, err
	}

	if len(existing) == 0 {
		err := errors.New("entity not found")

		return types.UpsertResponse{}, err
	}

	params := mapUpsertRequestToUpdateParams(req)
	err = qtx.UpdateWebhookEvent(
		ctx,
		params,
	)
	if err != nil {

		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {

			return types.UpsertResponse{}, err
		}
	}

	return buildUpdateResponse(req), nil
}

func buildUpdateResponse(req types.UpsertRequest) types.UpsertResponse {
	return types.UpsertResponse{

		ID: req.WebhookEvent.ID,
	}
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest) mfconsultdb.UpdateWebhookEventParams {
	return mfconsultdb.UpdateWebhookEventParams{
		ID: req.WebhookEvent.ID,

		Source: req.WebhookEvent.Source.ToNullInt(),

		EventType: req.WebhookEvent.EventType,

		Payload: req.WebhookEvent.Payload,

		ProcessedAt: req.WebhookEvent.ProcessedAt,

		CreatedAt: req.WebhookEvent.CreatedAt,
	}
}
