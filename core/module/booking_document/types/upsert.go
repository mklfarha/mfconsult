package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/booking_document"

	"github.com/gofrs/uuid"
)

type UpsertRequest struct {
	BookingDocument main_entity.BookingDocument
}

type UpsertResponse struct {
	ID uuid.UUID
}
