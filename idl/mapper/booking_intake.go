package mapper

import (
	main_entity "github.com/mklfarha/mfconsult/entity/booking_intake"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	"github.com/guregu/null/v6"

	"github.com/mklfarha/mfconsult/enum"
)

func BookingIntakeToProto(e main_entity.BookingIntake) *pb.BookingIntake {
	return &pb.BookingIntake{
		Reason:       e.Reason.ValueOrZero(),
		HelpTopic:    pb.HelpTopic(e.HelpTopic),
		HelpDetails:  e.HelpDetails.ValueOrZero(),
		StackDetails: e.StackDetails.ValueOrZero(),
		PrepNotes:    e.PrepNotes.ValueOrZero(),
	}
}

func BookingIntakeSliceToProto(es []main_entity.BookingIntake) []*pb.BookingIntake {
	res := []*pb.BookingIntake{}
	for _, e := range es {
		res = append(res, BookingIntakeToProto(e))
	}
	return res
}

func BookingIntakeFromProto(m *pb.BookingIntake) main_entity.BookingIntake {
	if m == nil {
		return main_entity.BookingIntake{}
	}
	return main_entity.BookingIntake{
		Reason:       null.StringFrom(m.Reason),
		HelpTopic:    enum.HelpTopic(m.GetHelpTopic()),
		HelpDetails:  null.StringFrom(m.HelpDetails),
		StackDetails: null.StringFrom(m.StackDetails),
		PrepNotes:    null.StringFrom(m.PrepNotes),
	}
}

func BookingIntakeSliceFromProto(es []*pb.BookingIntake) []main_entity.BookingIntake {
	if es == nil {
		return []main_entity.BookingIntake{}
	}
	res := []main_entity.BookingIntake{}
	for _, e := range es {
		res = append(res, BookingIntakeFromProto(e))
	}
	return res
}
