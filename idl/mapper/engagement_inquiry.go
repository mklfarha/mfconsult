package mapper

import (
	main_entity "github.com/mklfarha/mfconsult/entity/engagement_inquiry"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	"github.com/guregu/null/v6"

	"github.com/mklfarha/mfconsult/enum"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func EngagementInquiryToProto(e main_entity.EngagementInquiry) *pb.EngagementInquiry {
	return &pb.EngagementInquiry{
		Id:                 e.ID.String(),
		ClientId:           StringFromUUIDPtr(e.ClientId),
		Name:               e.Name,
		Email:              e.Email,
		Phone:              e.Phone.ValueOrZero(),
		Company:            e.Company.ValueOrZero(),
		ProjectSummary:     e.ProjectSummary,
		WhyMoreThanSession: e.WhyMoreThanSession.ValueOrZero(),
		ScopeDetails:       e.ScopeDetails.ValueOrZero(),
		BudgetRange:        e.BudgetRange.ValueOrZero(),
		Timeline:           e.Timeline.ValueOrZero(),
		Status:             pb.InquiryStatus(e.Status),
		ReviewNotes:        e.ReviewNotes.ValueOrZero(),
		CreatedAt:          timestamppb.New(e.CreatedAt.ValueOrZero()),
		UpdatedAt:          timestamppb.New(e.UpdatedAt.ValueOrZero()),
	}
}

func EngagementInquirySliceToProto(es []main_entity.EngagementInquiry) []*pb.EngagementInquiry {
	res := []*pb.EngagementInquiry{}
	for _, e := range es {
		res = append(res, EngagementInquiryToProto(e))
	}
	return res
}

func EngagementInquiryFromProto(m *pb.EngagementInquiry) main_entity.EngagementInquiry {
	if m == nil {
		return main_entity.EngagementInquiry{}
	}
	return main_entity.EngagementInquiry{
		ID:                 StringToUUID(m.GetId()),
		ClientId:           StringToUUIDPtr(m.GetClientId()),
		Name:               m.GetName(),
		Email:              m.GetEmail(),
		Phone:              null.StringFrom(m.Phone),
		Company:            null.StringFrom(m.Company),
		ProjectSummary:     m.GetProjectSummary(),
		WhyMoreThanSession: null.StringFrom(m.WhyMoreThanSession),
		ScopeDetails:       null.StringFrom(m.ScopeDetails),
		BudgetRange:        null.StringFrom(m.BudgetRange),
		Timeline:           null.StringFrom(m.Timeline),
		Status:             enum.InquiryStatus(m.GetStatus()),
		ReviewNotes:        null.StringFrom(m.ReviewNotes),
		CreatedAt:          null.TimeFrom(m.GetCreatedAt().AsTime()),
		UpdatedAt:          null.TimeFrom(m.GetUpdatedAt().AsTime()),
	}
}

func EngagementInquirySliceFromProto(es []*pb.EngagementInquiry) []main_entity.EngagementInquiry {
	if es == nil {
		return []main_entity.EngagementInquiry{}
	}
	res := []main_entity.EngagementInquiry{}
	for _, e := range es {
		res = append(res, EngagementInquiryFromProto(e))
	}
	return res
}
