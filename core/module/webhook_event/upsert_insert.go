package webhook_event

import (
	"context"

	"github.com/mklfarha/mfconsult/core/module/webhook_event/types"
	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"
)

func (m *module) Insert(
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
	params := mapUpsertRequestToInsertParams(req)

	_, err := qtx.InsertWebhookEvent(
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

	return buildInsertResponse(req), nil
}

func buildInsertResponse(req types.UpsertRequest) types.UpsertResponse {
	return types.UpsertResponse{

		ID: req.WebhookEvent.ID,
	}
}

func mapUpsertRequestToInsertParams(req types.UpsertRequest) mfconsultdb.InsertWebhookEventParams {
	return mfconsultdb.InsertWebhookEventParams{
		ID: req.WebhookEvent.ID,

		Source: req.WebhookEvent.Source.ToNullInt(),

		EventType: req.WebhookEvent.EventType,

		Payload: req.WebhookEvent.Payload,

		ProcessedAt: req.WebhookEvent.ProcessedAt,

		CreatedAt: req.WebhookEvent.CreatedAt,
	}
}
