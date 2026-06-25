package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/magic_link"

	"github.com/gofrs/uuid"
)

type UpsertRequest struct {
	MagicLink main_entity.MagicLink
}

type UpsertResponse struct {
	ID uuid.UUID
}
