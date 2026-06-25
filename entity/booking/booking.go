package booking

import (
	"encoding/json"
	"log"

	"github.com/gofrs/uuid"
	"github.com/guregu/null/v6"
	"github.com/mklfarha/mfconsult/entity/booking_intake"
	"github.com/mklfarha/mfconsult/entity/booking_payment"
	"github.com/mklfarha/mfconsult/entity/booking_scheduling"
	"github.com/mklfarha/mfconsult/enum"

	"github.com/mklfarha/mfconsult/entity/mapper"
)

type Booking struct {
	ID              uuid.UUID                              `json:"id"`
	ClientId        uuid.UUID                              `json:"client_id"`
	Status          enum.BookingStatus                     `json:"status"`
	ReviewDecision  enum.ReviewDecision                    `json:"review_decision"`
	ReviewedAt      null.Time                              `json:"reviewed_at"`
	DeclineReason   null.String                            `json:"decline_reason"`
	Intake          []booking_intake.BookingIntake         `json:"intake"`
	Payment         []booking_payment.BookingPayment       `json:"payment"`
	Scheduling      []booking_scheduling.BookingScheduling `json:"scheduling"`
	TermsVersion    null.String                            `json:"terms_version"`
	TermsAcceptedAt null.Time                              `json:"terms_accepted_at"`
	TermsAcceptedIp null.String                            `json:"terms_accepted_ip"`
	CreatedAt       null.Time                              `json:"created_at"`
	UpdatedAt       null.Time                              `json:"updated_at"`
}

func (e Booking) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e Booking) PrimaryKeyValues() []string {
	return []string{
		e.ID.String(),
	}
}

func BookingFromJSON(data json.RawMessage) Booking {
	entity := Booking{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		if err2 := mapper.FlexibleUnmarshal(data, &entity); err2 != nil {
			log.Printf("flexible unmarshal error BookingFromJSON: %v\n", err2)
		}
	}
	return entity
}

func BookingSliceFromJSON(data json.RawMessage) []Booking {
	entity := []Booking{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		entity = []Booking{}
		var rawSlice []json.RawMessage
		if err2 := json.Unmarshal(data, &rawSlice); err2 == nil {
			for _, raw := range rawSlice {
				item := Booking{}
				if err3 := mapper.FlexibleUnmarshal(raw, &item); err3 != nil {
					log.Printf("flexible unmarshal error BookingSliceFromJSON item: %v\n", err3)
				}
				entity = append(entity, item)
			}
		}
	}
	return entity
}

func (e Booking) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingToJSON: %v\n", err)
	}
	return res
}

func BookingToJSON(e Booking) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingToJSON: %v\n", err)
	}
	return res
}

func BookingSliceToJSON(e []Booking) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingSliceToJSON: %v\n", err)
	}
	return res
}
