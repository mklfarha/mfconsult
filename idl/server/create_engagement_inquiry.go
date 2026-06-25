package server

import (
	"context"
	"errors"
	engagement_inquirymodule "github.com/mklfarha/mfconsult/core/module/engagement_inquiry"
	"github.com/mklfarha/mfconsult/core/module/engagement_inquiry/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"
)

func (s *server) CreateEngagementInquiry(ctx context.Context, req *pb.CreateEngagementInquiryRequest) (*pb.EngagementInquiry, error) {
	res, err := s.core.EngagementInquiry().Insert(ctx, types.UpsertRequest{
		EngagementInquiry: pbmapper.EngagementInquiryFromProto(req.GetEngagementInquiry()),
	})
	if err != nil {

		return nil, err
	}

	fetchRes, err := s.core.EngagementInquiry().FetchEngagementInquiryById(ctx, types.FetchEngagementInquiryByIdRequest(res), engagement_inquirymodule.WithSkipCache())
	if err != nil {

		return nil, err
	}

	if len(fetchRes.Results) == 0 {
		err := errors.New("error fetching entity")

		return nil, err
	}

	return pbmapper.EngagementInquiryToProto(fetchRes.Results[0]), nil
}
