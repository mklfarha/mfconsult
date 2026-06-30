package booking_document

import (
	"context"
	"errors"

	"github.com/mklfarha/mfconsult/core/module/booking_document/types"
)

func (m *module) Delete(
	ctx context.Context,
	req types.DeleteRequest,
	opts ...Option,
) error {
	optConfig := applyAllOptions(opts)

	tx := optConfig.SQLTx
	createdTx := false
	if tx == nil {
		ntx, err := m.repository.DB.Begin()
		if err != nil {
			return err
		}
		tx = ntx
		defer tx.Rollback()
		createdTx = true
	}

	qtx := m.repository.Queries.WithTx(tx)

	existing, err := qtx.FetchBookingDocumentByIdForUpdate(ctx,
		req.ID.String(),
	)
	if err != nil {

		return err
	}

	if len(existing) == 0 {
		err := errors.New("entity not found")

		return err
	}

	_, err = qtx.DeleteBookingDocument(
		ctx,
		req.ID.String(),
	)
	if err != nil {

		return err
	}

	if createdTx {
		err := tx.Commit()
		if err != nil {

			return err
		}
	}

	return nil
}
