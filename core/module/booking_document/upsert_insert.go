package booking_document

import (
	"context"

	"github.com/mklfarha/mfconsult/core/module/booking_document/types"
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

	_, err := qtx.InsertBookingDocument(
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

		ID: req.BookingDocument.ID,
	}
}

func mapUpsertRequestToInsertParams(req types.UpsertRequest) mfconsultdb.InsertBookingDocumentParams {
	return mfconsultdb.InsertBookingDocumentParams{
		ID: req.BookingDocument.ID.String(),

		BookingId: req.BookingDocument.BookingId.String(),

		Kind: req.BookingDocument.Kind.ToNullInt(),

		URL: req.BookingDocument.URL,

		Label: req.BookingDocument.Label,

		PurgeAfter: req.BookingDocument.PurgeAfter,

		CreatedAt: req.BookingDocument.CreatedAt,
	}
}
