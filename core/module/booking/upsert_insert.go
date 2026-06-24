package booking

import (
	"context"

	"github.com/mklfarha/mfconsult/core/module/booking/types"
	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"

	"github.com/mklfarha/mfconsult/entity/booking_intake"
	"github.com/mklfarha/mfconsult/entity/booking_payment"
	"github.com/mklfarha/mfconsult/entity/booking_scheduling"
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

	_, err := qtx.InsertBooking(
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

		ID: req.Booking.ID,
	}
}

func mapUpsertRequestToInsertParams(req types.UpsertRequest) mfconsultdb.InsertBookingParams {
	return mfconsultdb.InsertBookingParams{
		ID: req.Booking.ID.String(),

		ClientId: req.Booking.ClientId.String(),

		Status: req.Booking.Status.ToInt64(),

		ReviewDecision: req.Booking.ReviewDecision.ToInt64(),

		ReviewedAt: req.Booking.ReviewedAt,

		DeclineReason: req.Booking.DeclineReason,

		PayLinkToken: req.Booking.PayLinkToken,

		PayLinkExpiresAt: req.Booking.PayLinkExpiresAt,

		PortalToken: req.Booking.PortalToken,

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
