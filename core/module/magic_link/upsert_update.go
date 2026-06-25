package magic_link

import (
	"context"
	"errors"

	"github.com/mklfarha/mfconsult/core/module/magic_link/types"
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
	existing, err := qtx.FetchMagicLinkByIdForUpdate(ctx,
		req.MagicLink.ID.String(),
	)
	if err != nil {

		return types.UpsertResponse{}, err
	}

	if len(existing) == 0 {
		err := errors.New("entity not found")

		return types.UpsertResponse{}, err
	}

	params := mapUpsertRequestToUpdateParams(req)
	err = qtx.UpdateMagicLink(
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

		ID: req.MagicLink.ID,
	}
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest) mfconsultdb.UpdateMagicLinkParams {
	return mfconsultdb.UpdateMagicLinkParams{
		ID: req.MagicLink.ID.String(),

		ClientId: req.MagicLink.ClientId.String(),

		Email: req.MagicLink.Email,

		Token: req.MagicLink.Token,

		Purpose: req.MagicLink.Purpose.ToNullInt(),

		ExpiresAt: req.MagicLink.ExpiresAt,

		ConsumedAt: req.MagicLink.ConsumedAt,

		CreatedAt: req.MagicLink.CreatedAt,

		CreatedIp: req.MagicLink.CreatedIp,
	}
}
