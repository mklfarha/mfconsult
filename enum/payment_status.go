package enum

import (
	"encoding/json"
	"github.com/guregu/null/v6"
	"log"
)

//go:generate go run github.com/dmarkham/enumer -type=PaymentStatus -json
type PaymentStatus int64

const (
	PAYMENT_STATUS_INVALID = iota
	PAYMENT_STATUS_UNPAID
	PAYMENT_STATUS_PAID
	PAYMENT_STATUS_REFUNDED
)

func (e PaymentStatus) ToInt64() int64 {
	return int64(e)
}

func (e PaymentStatus) ToNullInt() null.Int {
	return null.NewInt(int64(e), true)
}

func PaymentStatusFromString(in string) PaymentStatus {
	switch in {
	case "invalid":
		return PAYMENT_STATUS_INVALID
	case "unpaid":
		return PAYMENT_STATUS_UNPAID
	case "paid":
		return PAYMENT_STATUS_PAID
	case "refunded":
		return PAYMENT_STATUS_REFUNDED
	}
	return PAYMENT_STATUS_INVALID
}

func PaymentStatusFromPointerString(in *string) PaymentStatus {
	if in == nil {
		return PAYMENT_STATUS_INVALID
	}
	return PaymentStatusFromString(*in)
}

func (e PaymentStatus) String() string {
	switch e {
	case PAYMENT_STATUS_INVALID:
		return "invalid"
	case PAYMENT_STATUS_UNPAID:
		return "unpaid"
	case PAYMENT_STATUS_PAID:
		return "paid"
	case PAYMENT_STATUS_REFUNDED:
		return "refunded"
	}

	return "invalid"
}

func (e PaymentStatus) StringPtr() *string {
	val := e.String()
	return &val
}

func PaymentStatusSliceToJSON(in []PaymentStatus) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		log.Printf("error marshaling PaymentStatus slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToPaymentStatusSlice(in json.RawMessage) []PaymentStatus {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		log.Printf("error unmarshaling PaymentStatus slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []PaymentStatus{}
	for _, r := range res {
		finalRes = append(finalRes, PaymentStatus(r))
	}
	return finalRes
}
