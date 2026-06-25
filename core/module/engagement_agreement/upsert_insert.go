package engagement_agreement

import (
	"context"

	"github.com/mklfarha/mfconsult/core/module/engagement_agreement/types"
	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"
)

func (m *module) Insert(
	ctx context.Context,
	req types.UpsertRequest,
	opts ...Option,
) (types.UpsertResponse, error) {

	optConfig := applyAllOptions(opts)

	tx := optConfig.SQLTx
	createdTx := false
	if tx == nil {
		ntx, err := m.repository.DB.Begin()
		if err != nil {
			return types.UpsertResponse{}, err
		}
		tx = ntx
		defer tx.Rollback()
		createdTx = true
	}

	qtx := m.repository.Queries.WithTx(tx)
	params := mapUpsertRequestToInsertParams(req)

	_, err := qtx.InsertEngagementAgreement(
		ctx,
		params,
	)
	if err != nil {

		return types.UpsertResponse{}, err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {

			return types.UpsertResponse{}, err
		}
	}

	return buildInsertResponse(req), nil
}

func buildInsertResponse(req types.UpsertRequest) types.UpsertResponse {
	return types.UpsertResponse{

		ID: req.EngagementAgreement.ID,
	}
}

func mapUpsertRequestToInsertParams(req types.UpsertRequest) mfconsultdb.InsertEngagementAgreementParams {
	return mfconsultdb.InsertEngagementAgreementParams{
		ID: req.EngagementAgreement.ID.String(),

		ClientId: req.EngagementAgreement.ClientId.String(),

		NdaURL: req.EngagementAgreement.NdaURL,

		Status: req.EngagementAgreement.Status.ToNullInt(),

		SignedAt: req.EngagementAgreement.SignedAt,

		CreatedAt: req.EngagementAgreement.CreatedAt,

		EnvelopeId: req.EngagementAgreement.EnvelopeId,

		CertificateURL: req.EngagementAgreement.CertificateURL,

		ContractURL: req.EngagementAgreement.ContractURL,

		EngagementInquiryId: req.EngagementAgreement.EngagementInquiryId.String(),

		UpdatedAt: req.EngagementAgreement.UpdatedAt,
	}
}
