package booking

import (
	"context"
	"errors"

	"github.com/mklfarha/mfconsult/core/module/booking/types"
	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"

	"github.com/mklfarha/mfconsult/entity/booking_intake"
	"github.com/mklfarha/mfconsult/entity/booking_payment"
	"github.com/mklfarha/mfconsult/entity/booking_scheduling"
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
	existing, err := qtx.FetchBookingByIdForUpdate(ctx,
		req.Booking.ID.String(),
	)
	if err != nil {

		return types.UpsertResponse{}, err
	}

	if len(existing) == 0 {
		err := errors.New("entity not found")

		return types.UpsertResponse{}, err
	}

	params := mapUpsertRequestToUpdateParams(req)
	err = qtx.UpdateBooking(
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

		ID: req.Booking.ID,
	}
}

func mapUpsertRequestToUpdateParams(req types.UpsertRequest) mfconsultdb.UpdateBookingParams {
	return mfconsultdb.UpdateBookingParams{
		ID: req.Booking.ID.String(),

		ClientId: req.Booking.ClientId.String(),

		Status: req.Booking.Status.ToInt64(),

		ReviewDecision: req.Booking.ReviewDecision.ToInt64(),

		ReviewedAt: req.Booking.ReviewedAt,

		DeclineReason: req.Booking.DeclineReason,

		Intake: booking_intake.BookingIntakeSliceToJSON(req.Booking.Intake),

		Payment: booking_payment.BookingPaymentSliceToJSON(req.Booking.Payment),

		Scheduling: booking_scheduling.BookingSchedulingSliceToJSON(req.Booking.Scheduling),

		TermsVersion: req.Booking.TermsVersion,

		TermsAcceptedAt: req.Booking.TermsAcceptedAt,

		TermsAcceptedIp: req.Booking.TermsAcceptedIp,

		CreatedAt: req.Booking.CreatedAt,

		UpdatedAt: req.Booking.UpdatedAt,
	}
}
