package mapper

import (
	main_entity "github.com/mklfarha/mfconsult/entity/client"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	"github.com/guregu/null/v6"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ClientToProto(e main_entity.Client) *pb.Client {
	return &pb.Client{
		Id:        e.ID.String(),
		Name:      e.Name,
		Email:     e.Email,
		Timezone:  e.Timezone.ValueOrZero(),
		Notes:     e.Notes.ValueOrZero(),
		CreatedAt: timestamppb.New(e.CreatedAt.ValueOrZero()),
		UpdatedAt: timestamppb.New(e.UpdatedAt.ValueOrZero()),
	}
}

func ClientSliceToProto(es []main_entity.Client) []*pb.Client {
	res := []*pb.Client{}
	for _, e := range es {
		res = append(res, ClientToProto(e))
	}
	return res
}

func ClientFromProto(m *pb.Client) main_entity.Client {
	if m == nil {
		return main_entity.Client{}
	}
	return main_entity.Client{
		ID:        StringToUUID(m.GetId()),
		Name:      m.GetName(),
		Email:     m.GetEmail(),
		Timezone:  null.StringFrom(m.Timezone),
		Notes:     null.StringFrom(m.Notes),
		CreatedAt: null.TimeFrom(m.GetCreatedAt().AsTime()),
		UpdatedAt: null.TimeFrom(m.GetUpdatedAt().AsTime()),
	}
}

func ClientSliceFromProto(es []*pb.Client) []main_entity.Client {
	if es == nil {
		return []main_entity.Client{}
	}
	res := []main_entity.Client{}
	for _, e := range es {
		res = append(res, ClientFromProto(e))
	}
	return res
}
