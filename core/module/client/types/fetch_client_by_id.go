package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/client"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchClientByIdRequest struct {
	ID uuid.UUID
}

func (r FetchClientByIdRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("id", r.ID.String())

	return nil
}

type FetchClientByIdResponse struct {
	Results []main_entity.Client
}
