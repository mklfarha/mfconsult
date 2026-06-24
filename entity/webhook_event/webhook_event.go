package webhook_event

import (
	"encoding/json"
	"log"

	"github.com/guregu/null/v6"
	"github.com/mklfarha/mfconsult/enum"

	"github.com/mklfarha/mfconsult/entity/mapper"
)

type WebhookEvent struct {
	ID          string             `json:"id"`
	Source      enum.WebhookSource `json:"source"`
	EventType   null.String        `json:"event_type"`
	Payload     null.String        `json:"payload"`
	ProcessedAt null.Time          `json:"processed_at"`
	CreatedAt   null.Time          `json:"created_at"`
}

func (e WebhookEvent) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e WebhookEvent) PrimaryKeyValues() []string {
	return []string{
		e.ID,
	}
}

func WebhookEventFromJSON(data json.RawMessage) WebhookEvent {
	entity := WebhookEvent{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		if err2 := mapper.FlexibleUnmarshal(data, &entity); err2 != nil {
			log.Printf("flexible unmarshal error WebhookEventFromJSON: %v\n", err2)
		}
	}
	return entity
}

func WebhookEventSliceFromJSON(data json.RawMessage) []WebhookEvent {
	entity := []WebhookEvent{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		entity = []WebhookEvent{}
		var rawSlice []json.RawMessage
		if err2 := json.Unmarshal(data, &rawSlice); err2 == nil {
			for _, raw := range rawSlice {
				item := WebhookEvent{}
				if err3 := mapper.FlexibleUnmarshal(raw, &item); err3 != nil {
					log.Printf("flexible unmarshal error WebhookEventSliceFromJSON item: %v\n", err3)
				}
				entity = append(entity, item)
			}
		}
	}
	return entity
}

func (e WebhookEvent) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error WebhookEventToJSON: %v\n", err)
	}
	return res
}

func WebhookEventToJSON(e WebhookEvent) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error WebhookEventToJSON: %v\n", err)
	}
	return res
}

func WebhookEventSliceToJSON(e []WebhookEvent) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error WebhookEventSliceToJSON: %v\n", err)
	}
	return res
}
