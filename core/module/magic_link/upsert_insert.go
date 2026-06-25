package magic_link

import (
	"context"

	"github.com/mklfarha/mfconsult/core/module/magic_link/types"
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

	_, err := qtx.InsertMagicLink(
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

		ID: req.MagicLink.ID,
	}
}

func mapUpsertRequestToInsertParams(req types.UpsertRequest) mfconsultdb.InsertMagicLinkParams {
	return mfconsultdb.InsertMagicLinkParams{
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
