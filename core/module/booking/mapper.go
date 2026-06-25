package booking

import (
	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"
	main_entity "github.com/mklfarha/mfconsult/entity/booking"

	"github.com/mklfarha/mfconsult/entity/mapper"

	"github.com/guregu/null/v6"
	"github.com/mklfarha/mfconsult/entity/booking_intake"
	"github.com/mklfarha/mfconsult/entity/booking_payment"
	"github.com/mklfarha/mfconsult/entity/booking_scheduling"
	"github.com/mklfarha/mfconsult/enum"
)

func mapModelsToEntities(models []mfconsultdb.Booking) []main_entity.Booking {
	result := []main_entity.Booking{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(m mfconsultdb.Booking) main_entity.Booking {
	return main_entity.Booking{
		ID:              mapper.StringToUUID(m.ID),
		ClientId:        mapper.StringToUUID(m.ClientId),
		Status:          enum.BookingStatus(m.Status),
		ReviewDecision:  enum.ReviewDecision(m.ReviewDecision),
		ReviewedAt:      null.NewTime(m.ReviewedAt.Time, m.ReviewedAt.Valid),
		DeclineReason:   null.NewString(m.DeclineReason.String, m.DeclineReason.Valid),
		Intake:          booking_intake.BookingIntakeSliceFromJSON(m.Intake),
		Payment:         booking_payment.BookingPaymentSliceFromJSON(m.Payment),
		Scheduling:      booking_scheduling.BookingSchedulingSliceFromJSON(m.Scheduling),
		TermsVersion:    null.NewString(m.TermsVersion.String, m.TermsVersion.Valid),
		TermsAcceptedAt: null.NewTime(m.TermsAcceptedAt.Time, m.TermsAcceptedAt.Valid),
		TermsAcceptedIp: null.NewString(m.TermsAcceptedIp.String, m.TermsAcceptedIp.Valid),
		CreatedAt:       null.NewTime(m.CreatedAt.Time, m.CreatedAt.Valid),
		UpdatedAt:       null.NewTime(m.UpdatedAt.Time, m.UpdatedAt.Valid),
	}
}
