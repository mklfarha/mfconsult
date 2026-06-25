package server

import (
	"context"
	"errors"
	magic_linkmodule "github.com/mklfarha/mfconsult/core/module/magic_link"
	"github.com/mklfarha/mfconsult/core/module/magic_link/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"

	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateMagicLink(ctx context.Context, req *pb.UpdateMagicLinkRequest) (*pb.MagicLink, error) {

	if req.MagicLink.Id == "" {
		return nil, errors.New("please provide a valid ID to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req.GetMagicLink())
	if err != nil {

		return nil, err
	}

	isFull := fieldmask.IsFullReplacement(req.UpdateMask)

	if !isFull && req.UpdateMask != nil {

		if !strings.Contains(req.UpdateMask.String(), "id") {
			req.UpdateMask.Append(req.GetMagicLink(), "id")
		}

		pkEntity := pbmapper.MagicLinkFromProto(req.GetMagicLink())
		existingRes, err := s.core.MagicLink().FetchMagicLinkById(ctx,
			types.FetchMagicLinkByIdRequest{
				ID: pkEntity.ID,
			},
			magic_linkmodule.WithSkipCache(),
		)
		if err != nil {

			return nil, err
		}
		if len(existingRes.Results) == 0 {
			return nil, errors.New("entity not found")
		}

		merged := pbmapper.MagicLinkToProto(existingRes.Results[0])
		fieldmask.Update(req.UpdateMask, merged, req.GetMagicLink())
		req = &pb.UpdateMagicLinkRequest{MagicLink: merged, UpdateMask: req.UpdateMask}
	}

	res, err := s.core.MagicLink().Update(ctx, types.UpsertRequest{
		MagicLink: pbmapper.MagicLinkFromProto(req.GetMagicLink()),
	})
	if err != nil {

		return nil, err
	}

	fetchRes, err := s.core.MagicLink().FetchMagicLinkById(ctx, types.FetchMagicLinkByIdRequest(res), magic_linkmodule.WithSkipCache())
	if err != nil {

		return nil, err
	}

	if len(fetchRes.Results) == 0 {
		err := errors.New("error fetching entity")

		return nil, err
	}

	return pbmapper.MagicLinkToProto(fetchRes.Results[0]), nil
}
