package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/webhook_event"

	"go.uber.org/zap/zapcore"
)

type FetchWebhookEventByIdRequest struct {
	ID string
}

func (r FetchWebhookEventByIdRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	return nil
}

type FetchWebhookEventByIdResponse struct {
	Results []main_entity.WebhookEvent
}
