package types

import (
	"github.com/gofrs/uuid"
)

type DeleteRequest struct {
	ID uuid.UUID
}
