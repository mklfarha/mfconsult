package webhook_event

import (
	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"
	main_entity "github.com/mklfarha/mfconsult/entity/webhook_event"

	"github.com/guregu/null/v6"
	"github.com/mklfarha/mfconsult/enum"
)

func mapModelsToEntities(models []mfconsultdb.WebhookEvent) []main_entity.WebhookEvent {
	result := []main_entity.WebhookEvent{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(m mfconsultdb.WebhookEvent) main_entity.WebhookEvent {
	return main_entity.WebhookEvent{
		ID:          m.ID,
		Source:      enum.WebhookSource(m.Source.Int64),
		EventType:   null.NewString(m.EventType.String, m.EventType.Valid),
		Payload:     null.NewString(m.Payload.String, m.Payload.Valid),
		ProcessedAt: null.NewTime(m.ProcessedAt.Time, m.ProcessedAt.Valid),
		CreatedAt:   null.NewTime(m.CreatedAt.Time, m.CreatedAt.Valid),
	}
}
