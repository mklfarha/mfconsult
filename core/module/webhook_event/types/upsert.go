package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/webhook_event"
)

type UpsertRequest struct {
	WebhookEvent main_entity.WebhookEvent
}

type UpsertResponse struct {
	ID string
}
