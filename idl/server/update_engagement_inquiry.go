package server

import (
	"context"
	"errors"
	engagement_inquirymodule "github.com/mklfarha/mfconsult/core/module/engagement_inquiry"
	"github.com/mklfarha/mfconsult/core/module/engagement_inquiry/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"

	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateEngagementInquiry(ctx context.Context, req *pb.UpdateEngagementInquiryRequest) (*pb.EngagementInquiry, error) {

	if req.EngagementInquiry.Id == "" {
		return nil, errors.New("please provide a valid ID to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req.GetEngagementInquiry())
	if err != nil {

		return nil, err
	}

	isFull := fieldmask.IsFullReplacement(req.UpdateMask)

	if !isFull && req.UpdateMask != nil {

		if !strings.Contains(req.UpdateMask.String(), "id") {
			req.UpdateMask.Append(req.GetEngagementInquiry(), "id")
		}

		pkEntity := pbmapper.EngagementInquiryFromProto(req.GetEngagementInquiry())
		existingRes, err := s.core.EngagementInquiry().FetchEngagementInquiryById(ctx,
			types.FetchEngagementInquiryByIdRequest{
				ID: pkEntity.ID,
			},
			engagement_inquirymodule.WithSkipCache(),
		)
		if err != nil {

			return nil, err
		}
		if len(existingRes.Results) == 0 {
			return nil, errors.New("entity not found")
		}

		merged := pbmapper.EngagementInquiryToProto(existingRes.Results[0])
		fieldmask.Update(req.UpdateMask, merged, req.GetEngagementInquiry())
		req = &pb.UpdateEngagementInquiryRequest{EngagementInquiry: merged, UpdateMask: req.UpdateMask}
	}

	res, err := s.core.EngagementInquiry().Update(ctx, types.UpsertRequest{
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
