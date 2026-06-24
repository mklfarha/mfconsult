package mapper

import (
	main_entity "github.com/mklfarha/mfconsult/entity/booking_scheduling"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/guregu/null/v6"
)

func BookingSchedulingToProto(e main_entity.BookingScheduling) *pb.BookingScheduling {
	return &pb.BookingScheduling{
		SlotStart:          timestamppb.New(e.SlotStart.ValueOrZero()),
		SlotEnd:            timestamppb.New(e.SlotEnd.ValueOrZero()),
		SchedulerBookingId: e.SchedulerBookingId.ValueOrZero(),
		VideoUrl:           e.VideoURL.ValueOrZero(),
	}
}

func BookingSchedulingSliceToProto(es []main_entity.BookingScheduling) []*pb.BookingScheduling {
	res := []*pb.BookingScheduling{}
	for _, e := range es {
		res = append(res, BookingSchedulingToProto(e))
	}
	return res
}

func BookingSchedulingFromProto(m *pb.BookingScheduling) main_entity.BookingScheduling {
	if m == nil {
		return main_entity.BookingScheduling{}
	}
	return main_entity.BookingScheduling{
		SlotStart:          null.TimeFrom(m.GetSlotStart().AsTime()),
		SlotEnd:            null.TimeFrom(m.GetSlotEnd().AsTime()),
		SchedulerBookingId: null.StringFrom(m.SchedulerBookingId),
		VideoURL:           null.StringFrom(m.VideoUrl),
	}
}

func BookingSchedulingSliceFromProto(es []*pb.BookingScheduling) []main_entity.BookingScheduling {
	if es == nil {
		return []main_entity.BookingScheduling{}
	}
	res := []main_entity.BookingScheduling{}
	for _, e := range es {
		res = append(res, BookingSchedulingFromProto(e))
	}
	return res
}
