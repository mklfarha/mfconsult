package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/booking"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchBookingByIdRequest struct {
	ID uuid.UUID
}

func (r FetchBookingByIdRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("id", r.ID.String())

	return nil
}

type FetchBookingByIdResponse struct {
	Results []main_entity.Booking
}
