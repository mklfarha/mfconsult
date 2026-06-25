package engagement_inquiry

import (
	"context"
	"errors"

	"github.com/mklfarha/mfconsult/core/module/engagement_inquiry/types"
	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"

	"github.com/mklfarha/mfconsult/entity/mapper"
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
	existing, err := qtx.FetchEngagementInquiryByIdForUpdate(ctx,
		req.EngagementInquiry.ID.String(),
	)
	if err != nil {

		return types.UpsertResponse{}, err
	}

	if len(existing) == 0 {
		err := errors.New("entity not found")

		return types.UpsertResponse{}, err
	}

	params := mapUpsertRequestToUpdateParams(req)
	err = qtx.UpdateEngagementInquiry(
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

		ID: req.EngagementInquiry.ID,
	}
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest) mfconsultdb.UpdateEngagementInquiryParams {
	return mfconsultdb.UpdateEngagementInquiryParams{
		ID: req.EngagementInquiry.ID.String(),

		ClientId: mapper.UUIDPtrToNullString(req.EngagementInquiry.ClientId),

		Name: req.EngagementInquiry.Name,

		Email: req.EngagementInquiry.Email,

		Phone: req.EngagementInquiry.Phone,

		Company: req.EngagementInquiry.Company,

		ProjectSummary: req.EngagementInquiry.ProjectSummary,

		WhyMoreThanSession: req.EngagementInquiry.WhyMoreThanSession,

		ScopeDetails: req.EngagementInquiry.ScopeDetails,

		BudgetRange: req.EngagementInquiry.BudgetRange,

		Timeline: req.EngagementInquiry.Timeline,

		Status: req.EngagementInquiry.Status.ToNullInt(),

		ReviewNotes: req.EngagementInquiry.ReviewNotes,

		CreatedAt: req.EngagementInquiry.CreatedAt,

		UpdatedAt: req.EngagementInquiry.UpdatedAt,
	}
}
