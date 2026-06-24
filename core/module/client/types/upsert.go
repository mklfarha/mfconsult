package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/client"

	"github.com/gofrs/uuid"
)

type UpsertRequest struct {
	Client main_entity.Client
}

type UpsertResponse struct {
	ID uuid.UUID
}
