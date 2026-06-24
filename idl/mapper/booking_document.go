package mapper

import (
	main_entity "github.com/mklfarha/mfconsult/entity/booking_document"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	"github.com/mklfarha/mfconsult/enum"

	"github.com/guregu/null/v6"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func BookingDocumentToProto(e main_entity.BookingDocument) *pb.BookingDocument {
	return &pb.BookingDocument{
		Id:         e.ID.String(),
		BookingId:  e.BookingId.String(),
		Kind:       pb.DocumentKind(e.Kind),
		Url:        e.URL.ValueOrZero(),
		Label:      e.Label.ValueOrZero(),
		PurgeAfter: timestamppb.New(e.PurgeAfter.ValueOrZero()),
		CreatedAt:  timestamppb.New(e.CreatedAt.ValueOrZero()),
	}
}

func BookingDocumentSliceToProto(es []main_entity.BookingDocument) []*pb.BookingDocument {
	res := []*pb.BookingDocument{}
	for _, e := range es {
		res = append(res, BookingDocumentToProto(e))
	}
	return res
}

func BookingDocumentFromProto(m *pb.BookingDocument) main_entity.BookingDocument {
	if m == nil {
		return main_entity.BookingDocument{}
	}
	return main_entity.BookingDocument{
		ID:         StringToUUID(m.GetId()),
		BookingId:  StringToUUID(m.GetBookingId()),
		Kind:       enum.DocumentKind(m.GetKind()),
		URL:        null.StringFrom(m.Url),
		Label:      null.StringFrom(m.Label),
		PurgeAfter: null.TimeFrom(m.GetPurgeAfter().AsTime()),
		CreatedAt:  null.TimeFrom(m.GetCreatedAt().AsTime()),
	}
}

func BookingDocumentSliceFromProto(es []*pb.BookingDocument) []main_entity.BookingDocument {
	if es == nil {
		return []main_entity.BookingDocument{}
	}
	res := []main_entity.BookingDocument{}
	for _, e := range es {
		res = append(res, BookingDocumentFromProto(e))
	}
	return res
}
