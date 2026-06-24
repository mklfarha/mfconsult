package mapper

import (
	main_entity "github.com/mklfarha/mfconsult/entity/booking_payment"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	"github.com/guregu/null/v6"

	"github.com/mklfarha/mfconsult/enum"
)

func BookingPaymentToProto(e main_entity.BookingPayment) *pb.BookingPayment {
	return &pb.BookingPayment{
		AmountCents:   e.AmountCents.ValueOrZero(),
		Currency:      e.Currency.ValueOrZero(),
		PaymentStatus: pb.PaymentStatus(e.PaymentStatus),
		StripeRef:     e.StripeRef.ValueOrZero(),
	}
}

func BookingPaymentSliceToProto(es []main_entity.BookingPayment) []*pb.BookingPayment {
	res := []*pb.BookingPayment{}
	for _, e := range es {
		res = append(res, BookingPaymentToProto(e))
	}
	return res
}

func BookingPaymentFromProto(m *pb.BookingPayment) main_entity.BookingPayment {
	if m == nil {
		return main_entity.BookingPayment{}
	}
	return main_entity.BookingPayment{
		AmountCents:   null.IntFrom(m.GetAmountCents()),
		Currency:      null.StringFrom(m.Currency),
		PaymentStatus: enum.PaymentStatus(m.GetPaymentStatus()),
		StripeRef:     null.StringFrom(m.StripeRef),
	}
}

func BookingPaymentSliceFromProto(es []*pb.BookingPayment) []main_entity.BookingPayment {
	if es == nil {
		return []main_entity.BookingPayment{}
	}
	res := []main_entity.BookingPayment{}
	for _, e := range es {
		res = append(res, BookingPaymentFromProto(e))
	}
	return res
}
