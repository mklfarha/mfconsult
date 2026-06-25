package server

import (
	"context"
	"errors"
	engagement_agreementmodule "github.com/mklfarha/mfconsult/core/module/engagement_agreement"
	"github.com/mklfarha/mfconsult/core/module/engagement_agreement/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"

	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateEngagementAgreement(ctx context.Context, req *pb.UpdateEngagementAgreementRequest) (*pb.EngagementAgreement, error) {

	if req.EngagementAgreement.Id == "" {
		return nil, errors.New("please provide a valid ID to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req.GetEngagementAgreement())
	if err != nil {

		return nil, err
	}

	isFull := fieldmask.IsFullReplacement(req.UpdateMask)

	if !isFull && req.UpdateMask != nil {

		if !strings.Contains(req.UpdateMask.String(), "id") {
			req.UpdateMask.Append(req.GetEngagementAgreement(), "id")
		}

		pkEntity := pbmapper.EngagementAgreementFromProto(req.GetEngagementAgreement())
		existingRes, err := s.core.EngagementAgreement().FetchEngagementAgreementById(ctx,
			types.FetchEngagementAgreementByIdRequest{
				ID: pkEntity.ID,
			},
			engagement_agreementmodule.WithSkipCache(),
		)
		if err != nil {

			return nil, err
		}
		if len(existingRes.Results) == 0 {
			return nil, errors.New("entity not found")
		}

		merged := pbmapper.EngagementAgreementToProto(existingRes.Results[0])
		fieldmask.Update(req.UpdateMask, merged, req.GetEngagementAgreement())
		req = &pb.UpdateEngagementAgreementRequest{EngagementAgreement: merged, UpdateMask: req.UpdateMask}
	}

	res, err := s.core.EngagementAgreement().Update(ctx, types.UpsertRequest{
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
