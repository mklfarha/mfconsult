package engagement_agreement

import (
	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"
	main_entity "github.com/mklfarha/mfconsult/entity/engagement_agreement"

	"github.com/mklfarha/mfconsult/entity/mapper"

	"github.com/guregu/null/v6"
	"github.com/mklfarha/mfconsult/enum"
)

func mapModelsToEntities(models []mfconsultdb.EngagementAgreement) []main_entity.EngagementAgreement {
	result := []main_entity.EngagementAgreement{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(m mfconsultdb.EngagementAgreement) main_entity.EngagementAgreement {
	return main_entity.EngagementAgreement{
		ID:                  mapper.StringToUUID(m.ID),
		ClientId:            mapper.StringToUUID(m.ClientId),
		NdaURL:              null.NewString(m.NdaURL.String, m.NdaURL.Valid),
		Status:              enum.AgreementStatus(m.Status.Int64),
		SignedAt:            null.NewTime(m.SignedAt.Time, m.SignedAt.Valid),
		CreatedAt:           null.NewTime(m.CreatedAt.Time, m.CreatedAt.Valid),
		EnvelopeId:          null.NewString(m.EnvelopeId.String, m.EnvelopeId.Valid),
		CertificateURL:      null.NewString(m.CertificateURL.String, m.CertificateURL.Valid),
		ContractURL:         null.NewString(m.ContractURL.String, m.ContractURL.Valid),
		EngagementInquiryId: mapper.StringToUUID(m.EngagementInquiryId),
		UpdatedAt:           null.NewTime(m.UpdatedAt.Time, m.UpdatedAt.Valid),
	}
}
