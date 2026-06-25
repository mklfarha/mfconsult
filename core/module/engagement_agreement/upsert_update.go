package engagement_agreement

import (
	"context"
	"errors"

	"github.com/mklfarha/mfconsult/core/module/engagement_agreement/types"
	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"
)

func (m *module) Update(
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
	existing, err := qtx.FetchEngagementAgreementByIdForUpdate(ctx,
		req.EngagementAgreement.ID.String(),
	)
	if err != nil {

		return types.UpsertResponse{}, err
	}

	if len(existing) == 0 {
		err := errors.New("entity not found")

		return types.UpsertResponse{}, err
	}

	params := mapUpsertRequestToUpdateParams(req)
	err = qtx.UpdateEngagementAgreement(
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

	return buildUpdateResponse(req), nil
}

func buildUpdateResponse(req types.UpsertRequest) types.UpsertResponse {
	return types.UpsertResponse{

		ID: req.EngagementAgreement.ID,
	}
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest) mfconsultdb.UpdateEngagementAgreementParams {
	return mfconsultdb.UpdateEngagementAgreementParams{
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
