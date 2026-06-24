package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/nda_document"

	"github.com/gofrs/uuid"
)

type UpsertRequest struct {
	NdaDocument main_entity.NdaDocument
}

type UpsertResponse struct {
	ID uuid.UUID
}
