package mapper

import (
	main_entity "github.com/mklfarha/mfconsult/entity/nda_document"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	"github.com/guregu/null/v6"

	"github.com/mklfarha/mfconsult/enum"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func NdaDocumentToProto(e main_entity.NdaDocument) *pb.NdaDocument {
	return &pb.NdaDocument{
		Id:             e.ID.String(),
		ClientId:       e.ClientId.String(),
		Url:            e.URL.ValueOrZero(),
		Status:         pb.NdaStatus(e.Status),
		SignedAt:       timestamppb.New(e.SignedAt.ValueOrZero()),
		CreatedAt:      timestamppb.New(e.CreatedAt.ValueOrZero()),
		EnvelopeId:     e.EnvelopeId.ValueOrZero(),
		CertificateUrl: e.CertificateURL.ValueOrZero(),
	}
}

func NdaDocumentSliceToProto(es []main_entity.NdaDocument) []*pb.NdaDocument {
	res := []*pb.NdaDocument{}
	for _, e := range es {
		res = append(res, NdaDocumentToProto(e))
	}
	return res
}

func NdaDocumentFromProto(m *pb.NdaDocument) main_entity.NdaDocument {
	if m == nil {
		return main_entity.NdaDocument{}
	}
	return main_entity.NdaDocument{
		ID:             StringToUUID(m.GetId()),
		ClientId:       StringToUUID(m.GetClientId()),
		URL:            null.StringFrom(m.Url),
		Status:         enum.NdaStatus(m.GetStatus()),
		SignedAt:       null.TimeFrom(m.GetSignedAt().AsTime()),
		CreatedAt:      null.TimeFrom(m.GetCreatedAt().AsTime()),
		EnvelopeId:     null.StringFrom(m.EnvelopeId),
		CertificateURL: null.StringFrom(m.CertificateUrl),
	}
}

func NdaDocumentSliceFromProto(es []*pb.NdaDocument) []main_entity.NdaDocument {
	if es == nil {
		return []main_entity.NdaDocument{}
	}
	res := []main_entity.NdaDocument{}
	for _, e := range es {
		res = append(res, NdaDocumentFromProto(e))
	}
	return res
}
