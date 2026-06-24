package booking_document

import (
	"context"
	"errors"

	"github.com/mklfarha/mfconsult/core/module/booking_document/types"
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
	existing, err := qtx.FetchBookingDocumentByIdForUpdate(ctx,
		req.BookingDocument.ID.String(),
	)
	if err != nil {

		return types.UpsertResponse{}, err
	}

	if len(existing) == 0 {
		err := errors.New("entity not found")

		return types.UpsertResponse{}, err
	}

	params := mapUpsertRequestToUpdateParams(req)
	err = qtx.UpdateBookingDocument(
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

		ID: req.BookingDocument.ID,
	}
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest) mfconsultdb.UpdateBookingDocumentParams {
	return mfconsultdb.UpdateBookingDocumentParams{
		ID: req.BookingDocument.ID.String(),

		BookingId: req.BookingDocument.BookingId.String(),

		Kind: req.BookingDocument.Kind.ToNullInt(),

		URL: req.BookingDocument.URL,

		Label: req.BookingDocument.Label,

		PurgeAfter: req.BookingDocument.PurgeAfter,

		CreatedAt: req.BookingDocument.CreatedAt,
	}
}
