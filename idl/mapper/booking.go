package mapper

import (
	main_entity "github.com/mklfarha/mfconsult/entity/booking"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	"github.com/mklfarha/mfconsult/enum"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/guregu/null/v6"
)

func BookingToProto(e main_entity.Booking) *pb.Booking {
	return &pb.Booking{
		Id:              e.ID.String(),
		ClientId:        e.ClientId.String(),
		Status:          pb.BookingStatus(e.Status),
		ReviewDecision:  pb.ReviewDecision(e.ReviewDecision),
		ReviewedAt:      timestamppb.New(e.ReviewedAt.ValueOrZero()),
		DeclineReason:   e.DeclineReason.ValueOrZero(),
		Intake:          BookingIntakeSliceToProto(e.Intake),
		Payment:         BookingPaymentSliceToProto(e.Payment),
		Scheduling:      BookingSchedulingSliceToProto(e.Scheduling),
		TermsVersion:    e.TermsVersion.ValueOrZero(),
		TermsAcceptedAt: timestamppb.New(e.TermsAcceptedAt.ValueOrZero()),
		TermsAcceptedIp: e.TermsAcceptedIp.ValueOrZero(),
		CreatedAt:       timestamppb.New(e.CreatedAt.ValueOrZero()),
		UpdatedAt:       timestamppb.New(e.UpdatedAt.ValueOrZero()),
	}
}

func BookingSliceToProto(es []main_entity.Booking) []*pb.Booking {
	res := []*pb.Booking{}
	for _, e := range es {
		res = append(res, BookingToProto(e))
	}
	return res
}

func BookingFromProto(m *pb.Booking) main_entity.Booking {
	if m == nil {
		return main_entity.Booking{}
	}
	return main_entity.Booking{
		ID:              StringToUUID(m.GetId()),
		ClientId:        StringToUUID(m.GetClientId()),
		Status:          enum.BookingStatus(m.GetStatus()),
		ReviewDecision:  enum.ReviewDecision(m.GetReviewDecision()),
		ReviewedAt:      null.TimeFrom(m.GetReviewedAt().AsTime()),
		DeclineReason:   null.StringFrom(m.DeclineReason),
		Intake:          BookingIntakeSliceFromProto(m.GetIntake()),
		Payment:         BookingPaymentSliceFromProto(m.GetPayment()),
		Scheduling:      BookingSchedulingSliceFromProto(m.GetScheduling()),
		TermsVersion:    null.StringFrom(m.TermsVersion),
		TermsAcceptedAt: null.TimeFrom(m.GetTermsAcceptedAt().AsTime()),
		TermsAcceptedIp: null.StringFrom(m.TermsAcceptedIp),
		CreatedAt:       null.TimeFrom(m.GetCreatedAt().AsTime()),
		UpdatedAt:       null.TimeFrom(m.GetUpdatedAt().AsTime()),
	}
}

func BookingSliceFromProto(es []*pb.Booking) []main_entity.Booking {
	if es == nil {
		return []main_entity.Booking{}
	}
	res := []main_entity.Booking{}
	for _, e := range es {
		res = append(res, BookingFromProto(e))
	}
	return res
}
