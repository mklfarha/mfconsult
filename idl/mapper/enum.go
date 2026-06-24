package mapper

import (
	"github.com/mklfarha/mfconsult/enum"
	pb "github.com/mklfarha/mfconsult/idl/gen"
)

func BookingStatusSliceToProto(s []enum.BookingStatus) []pb.BookingStatus {
	res := []pb.BookingStatus{}
	for _, e := range s {
		res = append(res, pb.BookingStatus(e))
	}
	return res
}
func BookingStatusSliceFromProto(s []pb.BookingStatus) []enum.BookingStatus {
	res := []enum.BookingStatus{}
	for _, e := range s {
		res = append(res, enum.BookingStatus(e))
	}
	return res
}

func ReviewDecisionSliceToProto(s []enum.ReviewDecision) []pb.ReviewDecision {
	res := []pb.ReviewDecision{}
	for _, e := range s {
		res = append(res, pb.ReviewDecision(e))
	}
	return res
}
func ReviewDecisionSliceFromProto(s []pb.ReviewDecision) []enum.ReviewDecision {
	res := []enum.ReviewDecision{}
	for _, e := range s {
		res = append(res, enum.ReviewDecision(e))
	}
	return res
}

func HelpTopicSliceToProto(s []enum.HelpTopic) []pb.HelpTopic {
	res := []pb.HelpTopic{}
	for _, e := range s {
		res = append(res, pb.HelpTopic(e))
	}
	return res
}
func HelpTopicSliceFromProto(s []pb.HelpTopic) []enum.HelpTopic {
	res := []enum.HelpTopic{}
	for _, e := range s {
		res = append(res, enum.HelpTopic(e))
	}
	return res
}

func PaymentStatusSliceToProto(s []enum.PaymentStatus) []pb.PaymentStatus {
	res := []pb.PaymentStatus{}
	for _, e := range s {
		res = append(res, pb.PaymentStatus(e))
	}
	return res
}
func PaymentStatusSliceFromProto(s []pb.PaymentStatus) []enum.PaymentStatus {
	res := []enum.PaymentStatus{}
	for _, e := range s {
		res = append(res, enum.PaymentStatus(e))
	}
	return res
}

func DocumentKindSliceToProto(s []enum.DocumentKind) []pb.DocumentKind {
	res := []pb.DocumentKind{}
	for _, e := range s {
		res = append(res, pb.DocumentKind(e))
	}
	return res
}
func DocumentKindSliceFromProto(s []pb.DocumentKind) []enum.DocumentKind {
	res := []enum.DocumentKind{}
	for _, e := range s {
		res = append(res, enum.DocumentKind(e))
	}
	return res
}

func NdaStatusSliceToProto(s []enum.NdaStatus) []pb.NdaStatus {
	res := []pb.NdaStatus{}
	for _, e := range s {
		res = append(res, pb.NdaStatus(e))
	}
	return res
}
func NdaStatusSliceFromProto(s []pb.NdaStatus) []enum.NdaStatus {
	res := []enum.NdaStatus{}
	for _, e := range s {
		res = append(res, enum.NdaStatus(e))
	}
	return res
}

func WebhookSourceSliceToProto(s []enum.WebhookSource) []pb.WebhookSource {
	res := []pb.WebhookSource{}
	for _, e := range s {
		res = append(res, pb.WebhookSource(e))
	}
	return res
}
func WebhookSourceSliceFromProto(s []pb.WebhookSource) []enum.WebhookSource {
	res := []enum.WebhookSource{}
	for _, e := range s {
		res = append(res, enum.WebhookSource(e))
	}
	return res
}
