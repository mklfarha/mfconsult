package nda_document

import (
	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"
	main_entity "github.com/mklfarha/mfconsult/entity/nda_document"

	"github.com/mklfarha/mfconsult/entity/mapper"

	"github.com/guregu/null/v6"
	"github.com/mklfarha/mfconsult/enum"
)

func mapModelsToEntities(models []mfconsultdb.NdaDocument) []main_entity.NdaDocument {
	result := []main_entity.NdaDocument{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(m mfconsultdb.NdaDocument) main_entity.NdaDocument {
	return main_entity.NdaDocument{
		ID:             mapper.StringToUUID(m.ID),
		ClientId:       mapper.StringToUUID(m.ClientId),
		URL:            null.NewString(m.URL.String, m.URL.Valid),
		Status:         enum.NdaStatus(m.Status.Int64),
		SignedAt:       null.NewTime(m.SignedAt.Time, m.SignedAt.Valid),
		CreatedAt:      null.NewTime(m.CreatedAt.Time, m.CreatedAt.Valid),
		EnvelopeId:     null.NewString(m.EnvelopeId.String, m.EnvelopeId.Valid),
		CertificateURL: null.NewString(m.CertificateURL.String, m.CertificateURL.Valid),
	}
}
