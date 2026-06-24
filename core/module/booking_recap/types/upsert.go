package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/booking_recap"

	"github.com/gofrs/uuid"
)

type UpsertRequest struct {
	BookingRecap main_entity.BookingRecap
}

type UpsertResponse struct {
	ID uuid.UUID
}
