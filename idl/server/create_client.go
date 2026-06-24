package server

import (
	"context"
	"errors"
	clientmodule "github.com/mklfarha/mfconsult/core/module/client"
	"github.com/mklfarha/mfconsult/core/module/client/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"
)

func (s *server) CreateClient(ctx context.Context, req *pb.CreateClientRequest) (*pb.Client, error) {
	res, err := s.core.Client().Insert(ctx, types.UpsertRequest{
		Client: pbmapper.ClientFromProto(req.GetClient()),
	})
	if err != nil {

		return nil, err
	}

	fetchRes, err := s.core.Client().FetchClientById(ctx, types.FetchClientByIdRequest(res), clientmodule.WithSkipCache())
	if err != nil {

		return nil, err
	}

	if len(fetchRes.Results) == 0 {
		err := errors.New("error fetching entity")

		return nil, err
	}

	return pbmapper.ClientToProto(fetchRes.Results[0]), nil
}
