package server

import (
	"context"
	"errors"
	magic_linkmodule "github.com/mklfarha/mfconsult/core/module/magic_link"
	"github.com/mklfarha/mfconsult/core/module/magic_link/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"
)

func (s *server) CreateMagicLink(ctx context.Context, req *pb.CreateMagicLinkRequest) (*pb.MagicLink, error) {
	res, err := s.core.MagicLink().Insert(ctx, types.UpsertRequest{
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
