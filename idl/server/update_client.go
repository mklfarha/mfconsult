package server

import (
	"context"
	"errors"
	clientmodule "github.com/mklfarha/mfconsult/core/module/client"
	"github.com/mklfarha/mfconsult/core/module/client/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"

	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateClient(ctx context.Context, req *pb.UpdateClientRequest) (*pb.Client, error) {

	if req.Client.Id == "" {
		return nil, errors.New("please provide a valid ID to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req.GetClient())
	if err != nil {

		return nil, err
	}

	isFull := fieldmask.IsFullReplacement(req.UpdateMask)

	if !isFull && req.UpdateMask != nil {

		if !strings.Contains(req.UpdateMask.String(), "id") {
			req.UpdateMask.Append(req.GetClient(), "id")
		}

		pkEntity := pbmapper.ClientFromProto(req.GetClient())
		existingRes, err := s.core.Client().FetchClientById(ctx,
			types.FetchClientByIdRequest{
				ID: pkEntity.ID,
			},
			clientmodule.WithSkipCache(),
		)
		if err != nil {

			return nil, err
		}
		if len(existingRes.Results) == 0 {
			return nil, errors.New("entity not found")
		}

		merged := pbmapper.ClientToProto(existingRes.Results[0])
		fieldmask.Update(req.UpdateMask, merged, req.GetClient())
		req = &pb.UpdateClientRequest{Client: merged, UpdateMask: req.UpdateMask}
	}

	res, err := s.core.Client().Update(ctx, types.UpsertRequest{
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
