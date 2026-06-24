package server

import (
	"context"
	"errors"
	nda_documentmodule "github.com/mklfarha/mfconsult/core/module/nda_document"
	"github.com/mklfarha/mfconsult/core/module/nda_document/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"

	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateNdaDocument(ctx context.Context, req *pb.UpdateNdaDocumentRequest) (*pb.NdaDocument, error) {

	if req.NdaDocument.Id == "" {
		return nil, errors.New("please provide a valid ID to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req.GetNdaDocument())
	if err != nil {

		return nil, err
	}

	isFull := fieldmask.IsFullReplacement(req.UpdateMask)

	if !isFull && req.UpdateMask != nil {

		if !strings.Contains(req.UpdateMask.String(), "id") {
			req.UpdateMask.Append(req.GetNdaDocument(), "id")
		}

		pkEntity := pbmapper.NdaDocumentFromProto(req.GetNdaDocument())
		existingRes, err := s.core.NdaDocument().FetchNdaDocumentById(ctx,
			types.FetchNdaDocumentByIdRequest{
				ID: pkEntity.ID,
			},
			nda_documentmodule.WithSkipCache(),
		)
		if err != nil {

			return nil, err
		}
		if len(existingRes.Results) == 0 {
			return nil, errors.New("entity not found")
		}

		merged := pbmapper.NdaDocumentToProto(existingRes.Results[0])
		fieldmask.Update(req.UpdateMask, merged, req.GetNdaDocument())
		req = &pb.UpdateNdaDocumentRequest{NdaDocument: merged, UpdateMask: req.UpdateMask}
	}

	res, err := s.core.NdaDocument().Update(ctx, types.UpsertRequest{
		NdaDocument: pbmapper.NdaDocumentFromProto(req.GetNdaDocument()),
	})
	if err != nil {

		return nil, err
	}

	fetchRes, err := s.core.NdaDocument().FetchNdaDocumentById(ctx, types.FetchNdaDocumentByIdRequest(res), nda_documentmodule.WithSkipCache())
	if err != nil {

		return nil, err
	}

	if len(fetchRes.Results) == 0 {
		err := errors.New("error fetching entity")

		return nil, err
	}

	return pbmapper.NdaDocumentToProto(fetchRes.Results[0]), nil
}
