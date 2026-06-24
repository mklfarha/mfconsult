package booking_scheduling

import (
	"encoding/json"
	"log"

	"github.com/guregu/null/v6"

	"github.com/mklfarha/mfconsult/entity/mapper"
)

type BookingScheduling struct {
	SlotStart          null.Time   `json:"slot_start"`
	SlotEnd            null.Time   `json:"slot_end"`
	SchedulerBookingId null.String `json:"scheduler_booking_id"`
	VideoURL           null.String `json:"video_url"`
}

func (e BookingScheduling) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e BookingScheduling) PrimaryKeyValues() []string {
	return []string{}
}

func BookingSchedulingFromJSON(data json.RawMessage) BookingScheduling {
	entity := BookingScheduling{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		if err2 := mapper.FlexibleUnmarshal(data, &entity); err2 != nil {
			log.Printf("flexible unmarshal error BookingSchedulingFromJSON: %v\n", err2)
		}
	}
	return entity
}

func BookingSchedulingSliceFromJSON(data json.RawMessage) []BookingScheduling {
	entity := []BookingScheduling{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		entity = []BookingScheduling{}
		var rawSlice []json.RawMessage
		if err2 := json.Unmarshal(data, &rawSlice); err2 == nil {
			for _, raw := range rawSlice {
				item := BookingScheduling{}
				if err3 := mapper.FlexibleUnmarshal(raw, &item); err3 != nil {
					log.Printf("flexible unmarshal error BookingSchedulingSliceFromJSON item: %v\n", err3)
				}
				entity = append(entity, item)
			}
		}
	}
	return entity
}

func (e BookingScheduling) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingSchedulingToJSON: %v\n", err)
	}
	return res
}

func BookingSchedulingToJSON(e BookingScheduling) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingSchedulingToJSON: %v\n", err)
	}
	return res
}

func BookingSchedulingSliceToJSON(e []BookingScheduling) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingSchedulingSliceToJSON: %v\n", err)
	}
	return res
}
