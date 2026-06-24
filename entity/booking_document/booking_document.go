package booking_document

import (
	"encoding/json"
	"log"

	"github.com/gofrs/uuid"
	"github.com/guregu/null/v6"
	"github.com/mklfarha/mfconsult/enum"

	"github.com/mklfarha/mfconsult/entity/mapper"
)

type BookingDocument struct {
	ID         uuid.UUID         `json:"id"`
	BookingId  uuid.UUID         `json:"booking_id"`
	Kind       enum.DocumentKind `json:"kind"`
	URL        null.String       `json:"url"`
	Label      null.String       `json:"label"`
	PurgeAfter null.Time         `json:"purge_after"`
	CreatedAt  null.Time         `json:"created_at"`
}

func (e BookingDocument) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e BookingDocument) PrimaryKeyValues() []string {
	return []string{
		e.ID.String(),
	}
}

func BookingDocumentFromJSON(data json.RawMessage) BookingDocument {
	entity := BookingDocument{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		if err2 := mapper.FlexibleUnmarshal(data, &entity); err2 != nil {
			log.Printf("flexible unmarshal error BookingDocumentFromJSON: %v\n", err2)
		}
	}
	return entity
}

func BookingDocumentSliceFromJSON(data json.RawMessage) []BookingDocument {
	entity := []BookingDocument{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		entity = []BookingDocument{}
		var rawSlice []json.RawMessage
		if err2 := json.Unmarshal(data, &rawSlice); err2 == nil {
			for _, raw := range rawSlice {
				item := BookingDocument{}
				if err3 := mapper.FlexibleUnmarshal(raw, &item); err3 != nil {
					log.Printf("flexible unmarshal error BookingDocumentSliceFromJSON item: %v\n", err3)
				}
				entity = append(entity, item)
			}
		}
	}
	return entity
}

func (e BookingDocument) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingDocumentToJSON: %v\n", err)
	}
	return res
}

func BookingDocumentToJSON(e BookingDocument) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingDocumentToJSON: %v\n", err)
	}
	return res
}

func BookingDocumentSliceToJSON(e []BookingDocument) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingDocumentSliceToJSON: %v\n", err)
	}
	return res
}
