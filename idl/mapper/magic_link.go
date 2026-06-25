package mapper

import (
	main_entity "github.com/mklfarha/mfconsult/entity/magic_link"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	"github.com/guregu/null/v6"

	"github.com/mklfarha/mfconsult/enum"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MagicLinkToProto(e main_entity.MagicLink) *pb.MagicLink {
	return &pb.MagicLink{
		Id:         e.ID.String(),
		ClientId:   e.ClientId.String(),
		Email:      e.Email.ValueOrZero(),
		Token:      e.Token,
		Purpose:    pb.MagicLinkPurpose(e.Purpose),
		ExpiresAt:  timestamppb.New(e.ExpiresAt.ValueOrZero()),
		ConsumedAt: timestamppb.New(e.ConsumedAt.ValueOrZero()),
		CreatedAt:  timestamppb.New(e.CreatedAt.ValueOrZero()),
		CreatedIp:  e.CreatedIp.ValueOrZero(),
	}
}

func MagicLinkSliceToProto(es []main_entity.MagicLink) []*pb.MagicLink {
	res := []*pb.MagicLink{}
	for _, e := range es {
		res = append(res, MagicLinkToProto(e))
	}
	return res
}

func MagicLinkFromProto(m *pb.MagicLink) main_entity.MagicLink {
	if m == nil {
		return main_entity.MagicLink{}
	}
	return main_entity.MagicLink{
		ID:         StringToUUID(m.GetId()),
		ClientId:   StringToUUID(m.GetClientId()),
		Email:      null.StringFrom(m.Email),
		Token:      m.GetToken(),
		Purpose:    enum.MagicLinkPurpose(m.GetPurpose()),
		ExpiresAt:  null.TimeFrom(m.GetExpiresAt().AsTime()),
		ConsumedAt: null.TimeFrom(m.GetConsumedAt().AsTime()),
		CreatedAt:  null.TimeFrom(m.GetCreatedAt().AsTime()),
		CreatedIp:  null.StringFrom(m.CreatedIp),
	}
}

func MagicLinkSliceFromProto(es []*pb.MagicLink) []main_entity.MagicLink {
	if es == nil {
		return []main_entity.MagicLink{}
	}
	res := []main_entity.MagicLink{}
	for _, e := range es {
		res = append(res, MagicLinkFromProto(e))
	}
	return res
}
