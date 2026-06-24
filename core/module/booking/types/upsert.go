package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/booking"

	"github.com/gofrs/uuid"
)

type UpsertRequest struct {
	Booking main_entity.Booking
}

type UpsertResponse struct {
	ID uuid.UUID
}
