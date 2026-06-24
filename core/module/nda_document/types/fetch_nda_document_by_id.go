package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/nda_document"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchNdaDocumentByIdRequest struct {
	ID uuid.UUID
}

func (r FetchNdaDocumentByIdRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("id", r.ID.String())

	return nil
}

type FetchNdaDocumentByIdResponse struct {
	Results []main_entity.NdaDocument
}
