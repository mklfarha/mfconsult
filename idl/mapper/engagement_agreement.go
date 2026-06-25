package mapper

import (
	main_entity "github.com/mklfarha/mfconsult/entity/engagement_agreement"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	"github.com/guregu/null/v6"

	"github.com/mklfarha/mfconsult/enum"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func EngagementAgreementToProto(e main_entity.EngagementAgreement) *pb.EngagementAgreement {
	return &pb.EngagementAgreement{
		Id:                  e.ID.String(),
		ClientId:            e.ClientId.String(),
		NdaUrl:              e.NdaURL.ValueOrZero(),
		Status:              pb.AgreementStatus(e.Status),
		SignedAt:            timestamppb.New(e.SignedAt.ValueOrZero()),
		CreatedAt:           timestamppb.New(e.CreatedAt.ValueOrZero()),
		EnvelopeId:          e.EnvelopeId.ValueOrZero(),
		CertificateUrl:      e.CertificateURL.ValueOrZero(),
		ContractUrl:         e.ContractURL.ValueOrZero(),
		EngagementInquiryId: e.EngagementInquiryId.String(),
		UpdatedAt:           timestamppb.New(e.UpdatedAt.ValueOrZero()),
	}
}

func EngagementAgreementSliceToProto(es []main_entity.EngagementAgreement) []*pb.EngagementAgreement {
	res := []*pb.EngagementAgreement{}
	for _, e := range es {
		res = append(res, EngagementAgreementToProto(e))
	}
	return res
}

func EngagementAgreementFromProto(m *pb.EngagementAgreement) main_entity.EngagementAgreement {
	if m == nil {
		return main_entity.EngagementAgreement{}
	}
	return main_entity.EngagementAgreement{
		ID:                  StringToUUID(m.GetId()),
		ClientId:            StringToUUID(m.GetClientId()),
		NdaURL:              null.StringFrom(m.NdaUrl),
		Status:              enum.AgreementStatus(m.GetStatus()),
		SignedAt:            null.TimeFrom(m.GetSignedAt().AsTime()),
		CreatedAt:           null.TimeFrom(m.GetCreatedAt().AsTime()),
		EnvelopeId:          null.StringFrom(m.EnvelopeId),
		CertificateURL:      null.StringFrom(m.CertificateUrl),
		ContractURL:         null.StringFrom(m.ContractUrl),
		EngagementInquiryId: StringToUUID(m.GetEngagementInquiryId()),
		UpdatedAt:           null.TimeFrom(m.GetUpdatedAt().AsTime()),
	}
}

func EngagementAgreementSliceFromProto(es []*pb.EngagementAgreement) []main_entity.EngagementAgreement {
	if es == nil {
		return []main_entity.EngagementAgreement{}
	}
	res := []main_entity.EngagementAgreement{}
	for _, e := range es {
		res = append(res, EngagementAgreementFromProto(e))
	}
	return res
}
