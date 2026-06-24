package booking_payment

import (
	"encoding/json"
	"log"

	"github.com/guregu/null/v6"
	"github.com/mklfarha/mfconsult/enum"

	"github.com/mklfarha/mfconsult/entity/mapper"
)

type BookingPayment struct {
	AmountCents   null.Int64         `json:"amount_cents"`
	Currency      null.String        `json:"currency"`
	PaymentStatus enum.PaymentStatus `json:"payment_status"`
	StripeRef     null.String        `json:"stripe_ref"`
}

func (e BookingPayment) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e BookingPayment) PrimaryKeyValues() []string {
	return []string{}
}

func BookingPaymentFromJSON(data json.RawMessage) BookingPayment {
	entity := BookingPayment{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		if err2 := mapper.FlexibleUnmarshal(data, &entity); err2 != nil {
			log.Printf("flexible unmarshal error BookingPaymentFromJSON: %v\n", err2)
		}
	}
	return entity
}

func BookingPaymentSliceFromJSON(data json.RawMessage) []BookingPayment {
	entity := []BookingPayment{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		entity = []BookingPayment{}
		var rawSlice []json.RawMessage
		if err2 := json.Unmarshal(data, &rawSlice); err2 == nil {
			for _, raw := range rawSlice {
				item := BookingPayment{}
				if err3 := mapper.FlexibleUnmarshal(raw, &item); err3 != nil {
					log.Printf("flexible unmarshal error BookingPaymentSliceFromJSON item: %v\n", err3)
				}
				entity = append(entity, item)
			}
		}
	}
	return entity
}

func (e BookingPayment) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingPaymentToJSON: %v\n", err)
	}
	return res
}

func BookingPaymentToJSON(e BookingPayment) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingPaymentToJSON: %v\n", err)
	}
	return res
}

func BookingPaymentSliceToJSON(e []BookingPayment) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error BookingPaymentSliceToJSON: %v\n", err)
	}
	return res
}
