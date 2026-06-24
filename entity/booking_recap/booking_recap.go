package booking_recap

import (
	"encoding/json"
	"log"

	"github.com/gofrs/uuid"
	"github.com/guregu/null/v6"

	"github.com/mklfarha/mfconsult/entity/mapper"
)

type BookingRecap struct {
	ID          uuid.UUID   `json:"id"`
	BookingId   uuid.UUID   `json:"booking_id"`
	Body        null.String `json:"body"`
	PublishedAt null.Time   `json:"published_at"`
	CreatedAt   null.Time   `json:"created_at"`
}

func (e BookingRecap) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e BookingRecap) PrimaryKeyValues() []string {
	return []string{
		e.ID.String(),
	}
}

func BookingRecapFromJSON(data json.RawMessage) BookingRecap {
	entity := BookingRecap{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		if err2 := mapper.FlexibleUnmarshal(data, &entity); err2 != nil {
			log.Printf("flexible unmarshal error BookingRecapFromJSON: %v\n", err2)
		}
	}
	return entity
}

func BookingRecapSliceFromJSON(data json.RawMessage) []BookingRecap {
	entity := []BookingRecap{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		entity = []BookingRecap{}
		var rawSlice []json.RawMessage
		if err2 := json.Unmarshal(data, &rawSlice); err2 == nil {
			for _, raw := range rawSlice {
				item := BookingRecap{}
				if err3 := mapper.FlexibleUnmarshal(raw, &item); err3 != nil {
					log.Printf("flexible unmarshal error BookingRecapSliceFromJSON item: %v\n", err3)
				}
				entity = append(entity, item)
			}
		}
	}
	return entity
}

func (e BookingRecap) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingRecapToJSON: %v\n", err)
	}
	return res
}

func BookingRecapToJSON(e BookingRecap) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingRecapToJSON: %v\n", err)
	}
	return res
}

func BookingRecapSliceToJSON(e []BookingRecap) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingRecapSliceToJSON: %v\n", err)
	}
	return res
}
