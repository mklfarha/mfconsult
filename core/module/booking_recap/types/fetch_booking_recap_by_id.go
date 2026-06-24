package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/booking_recap"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchBookingRecapByIdRequest struct {
	ID uuid.UUID
}

func (r FetchBookingRecapByIdRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("id", r.ID.String())

	return nil
}

type FetchBookingRecapByIdResponse struct {
	Results []main_entity.BookingRecap
}
