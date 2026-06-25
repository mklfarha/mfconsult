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

func AgreementStatusSliceToProto(s []enum.AgreementStatus) []pb.AgreementStatus {
	res := []pb.AgreementStatus{}
	for _, e := range s {
		res = append(res, pb.AgreementStatus(e))
	}
	return res
}
func AgreementStatusSliceFromProto(s []pb.AgreementStatus) []enum.AgreementStatus {
	res := []enum.AgreementStatus{}
	for _, e := range s {
		res = append(res, enum.AgreementStatus(e))
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

func InquiryStatusSliceToProto(s []enum.InquiryStatus) []pb.InquiryStatus {
	res := []pb.InquiryStatus{}
	for _, e := range s {
		res = append(res, pb.InquiryStatus(e))
	}
	return res
}
func InquiryStatusSliceFromProto(s []pb.InquiryStatus) []enum.InquiryStatus {
	res := []enum.InquiryStatus{}
	for _, e := range s {
		res = append(res, enum.InquiryStatus(e))
	}
	return res
}

func MagicLinkPurposeSliceToProto(s []enum.MagicLinkPurpose) []pb.MagicLinkPurpose {
	res := []pb.MagicLinkPurpose{}
	for _, e := range s {
		res = append(res, pb.MagicLinkPurpose(e))
	}
	return res
}
func MagicLinkPurposeSliceFromProto(s []pb.MagicLinkPurpose) []enum.MagicLinkPurpose {
	res := []enum.MagicLinkPurpose{}
	for _, e := range s {
		res = append(res, enum.MagicLinkPurpose(e))
	}
	return res
}
