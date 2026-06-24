package mapper

import (
	main_entity "github.com/mklfarha/mfconsult/entity/booking_recap"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	"github.com/guregu/null/v6"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func BookingRecapToProto(e main_entity.BookingRecap) *pb.BookingRecap {
	return &pb.BookingRecap{
		Id:          e.ID.String(),
		BookingId:   e.BookingId.String(),
		Body:        e.Body.ValueOrZero(),
		PublishedAt: timestamppb.New(e.PublishedAt.ValueOrZero()),
		CreatedAt:   timestamppb.New(e.CreatedAt.ValueOrZero()),
	}
}

func BookingRecapSliceToProto(es []main_entity.BookingRecap) []*pb.BookingRecap {
	res := []*pb.BookingRecap{}
	for _, e := range es {
		res = append(res, BookingRecapToProto(e))
	}
	return res
}

func BookingRecapFromProto(m *pb.BookingRecap) main_entity.BookingRecap {
	if m == nil {
		return main_entity.BookingRecap{}
	}
	return main_entity.BookingRecap{
		ID:          StringToUUID(m.GetId()),
		BookingId:   StringToUUID(m.GetBookingId()),
		Body:        null.StringFrom(m.Body),
		PublishedAt: null.TimeFrom(m.GetPublishedAt().AsTime()),
		CreatedAt:   null.TimeFrom(m.GetCreatedAt().AsTime()),
	}
}

func BookingRecapSliceFromProto(es []*pb.BookingRecap) []main_entity.BookingRecap {
	if es == nil {
		return []main_entity.BookingRecap{}
	}
	res := []main_entity.BookingRecap{}
	for _, e := range es {
		res = append(res, BookingRecapFromProto(e))
	}
	return res
}
