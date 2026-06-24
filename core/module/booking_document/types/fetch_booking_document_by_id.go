package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/booking_document"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchBookingDocumentByIdRequest struct {
	ID uuid.UUID
}

func (r FetchBookingDocumentByIdRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("id", r.ID.String())

	return nil
}

type FetchBookingDocumentByIdResponse struct {
	Results []main_entity.BookingDocument
}
