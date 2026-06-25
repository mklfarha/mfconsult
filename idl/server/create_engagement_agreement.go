package server

import (
	"context"
	"errors"
	engagement_agreementmodule "github.com/mklfarha/mfconsult/core/module/engagement_agreement"
	"github.com/mklfarha/mfconsult/core/module/engagement_agreement/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"
)

func (s *server) CreateEngagementAgreement(ctx context.Context, req *pb.CreateEngagementAgreementRequest) (*pb.EngagementAgreement, error) {
	res, err := s.core.EngagementAgreement().Insert(ctx, types.UpsertRequest{
		EngagementAgreement: pbmapper.EngagementAgreementFromProto(req.GetEngagementAgreement()),
	})
	if err != nil {

		return nil, err
	}

	fetchRes, err := s.core.EngagementAgreement().FetchEngagementAgreementById(ctx, types.FetchEngagementAgreementByIdRequest(res), engagement_agreementmodule.WithSkipCache())
	if err != nil {

		return nil, err
	}

	if len(fetchRes.Results) == 0 {
		err := errors.New("error fetching entity")

		return nil, err
	}

	return pbmapper.EngagementAgreementToProto(fetchRes.Results[0]), nil
}
