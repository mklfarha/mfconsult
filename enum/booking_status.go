package enum

import (
	"encoding/json"
	"github.com/guregu/null/v6"
	"log"
)

//go:generate go run github.com/dmarkham/enumer -type=BookingStatus -json
type BookingStatus int64

const (
	BOOKING_STATUS_INVALID = iota
	BOOKING_STATUS_PENDING_REVIEW
	BOOKING_STATUS_APPROVED
	BOOKING_STATUS_PENDING_PAYMENT
	BOOKING_STATUS_CONFIRMED
	BOOKING_STATUS_COMPLETED
	BOOKING_STATUS_DECLINED
	BOOKING_STATUS_EXPIRED
	BOOKING_STATUS_CANCELLED
	BOOKING_STATUS_NO_SHOW
)

func (e BookingStatus) ToInt64() int64 {
	return int64(e)
}

func (e BookingStatus) ToNullInt() null.Int {
	return null.NewInt(int64(e), true)
}

func BookingStatusFromString(in string) BookingStatus {
	switch in {
	case "invalid":
		return BOOKING_STATUS_INVALID
	case "pending_review":
		return BOOKING_STATUS_PENDING_REVIEW
	case "approved":
		return BOOKING_STATUS_APPROVED
	case "pending_payment":
		return BOOKING_STATUS_PENDING_PAYMENT
	case "confirmed":
		return BOOKING_STATUS_CONFIRMED
	case "completed":
		return BOOKING_STATUS_COMPLETED
	case "declined":
		return BOOKING_STATUS_DECLINED
	case "expired":
		return BOOKING_STATUS_EXPIRED
	case "cancelled":
		return BOOKING_STATUS_CANCELLED
	case "no_show":
		return BOOKING_STATUS_NO_SHOW
	}
	return BOOKING_STATUS_INVALID
}

func BookingStatusFromPointerString(in *string) BookingStatus {
	if in == nil {
		return BOOKING_STATUS_INVALID
	}
	return BookingStatusFromString(*in)
}

func (e BookingStatus) String() string {
	switch e {
	case BOOKING_STATUS_INVALID:
		return "invalid"
	case BOOKING_STATUS_PENDING_REVIEW:
		return "pending_review"
	case BOOKING_STATUS_APPROVED:
		return "approved"
	case BOOKING_STATUS_PENDING_PAYMENT:
		return "pending_payment"
	case BOOKING_STATUS_CONFIRMED:
		return "confirmed"
	case BOOKING_STATUS_COMPLETED:
		return "completed"
	case BOOKING_STATUS_DECLINED:
		return "declined"
	case BOOKING_STATUS_EXPIRED:
		return "expired"
	case BOOKING_STATUS_CANCELLED:
		return "cancelled"
	case BOOKING_STATUS_NO_SHOW:
		return "no_show"
	}

	return "invalid"
}

func (e BookingStatus) StringPtr() *string {
	val := e.String()
	return &val
}

func BookingStatusSliceToJSON(in []BookingStatus) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		log.Printf("error marshaling BookingStatus slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToBookingStatusSlice(in json.RawMessage) []BookingStatus {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		log.Printf("error unmarshaling BookingStatus slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []BookingStatus{}
	for _, r := range res {
		finalRes = append(finalRes, BookingStatus(r))
	}
	return finalRes
}
