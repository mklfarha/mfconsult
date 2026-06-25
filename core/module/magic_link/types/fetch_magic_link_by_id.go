package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/magic_link"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchMagicLinkByIdRequest struct {
	ID uuid.UUID
}

func (r FetchMagicLinkByIdRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("id", r.ID.String())

	return nil
}

type FetchMagicLinkByIdResponse struct {
	Results []main_entity.MagicLink
}
