package booking_intake

import (
	"encoding/json"
	"log"

	"github.com/guregu/null/v6"
	"github.com/mklfarha/mfconsult/enum"

	"github.com/mklfarha/mfconsult/entity/mapper"
)

type BookingIntake struct {
	Reason       null.String    `json:"reason"`
	HelpTopic    enum.HelpTopic `json:"help_topic"`
	HelpDetails  null.String    `json:"help_details"`
	StackDetails null.String    `json:"stack_details"`
	PrepNotes    null.String    `json:"prep_notes"`
}

func (e BookingIntake) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e BookingIntake) PrimaryKeyValues() []string {
	return []string{}
}

func BookingIntakeFromJSON(data json.RawMessage) BookingIntake {
	entity := BookingIntake{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		if err2 := mapper.FlexibleUnmarshal(data, &entity); err2 != nil {
			log.Printf("flexible unmarshal error BookingIntakeFromJSON: %v\n", err2)
		}
	}
	return entity
}

func BookingIntakeSliceFromJSON(data json.RawMessage) []BookingIntake {
	entity := []BookingIntake{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		entity = []BookingIntake{}
		var rawSlice []json.RawMessage
		if err2 := json.Unmarshal(data, &rawSlice); err2 == nil {
			for _, raw := range rawSlice {
				item := BookingIntake{}
				if err3 := mapper.FlexibleUnmarshal(raw, &item); err3 != nil {
					log.Printf("flexible unmarshal error BookingIntakeSliceFromJSON item: %v\n", err3)
				}
				entity = append(entity, item)
			}
		}
	}
	return entity
}

func (e BookingIntake) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingIntakeToJSON: %v\n", err)
	}
	return res
}

func BookingIntakeToJSON(e BookingIntake) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingIntakeToJSON: %v\n", err)
	}
	return res
}

func BookingIntakeSliceToJSON(e []BookingIntake) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingIntakeSliceToJSON: %v\n", err)
	}
	return res
}
