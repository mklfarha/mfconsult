package server

import (
	"context"
	"errors"
	nda_documentmodule "github.com/mklfarha/mfconsult/core/module/nda_document"
	"github.com/mklfarha/mfconsult/core/module/nda_document/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"
)

func (s *server) CreateNdaDocument(ctx context.Context, req *pb.CreateNdaDocumentRequest) (*pb.NdaDocument, error) {
	res, err := s.core.NdaDocument().Insert(ctx, types.UpsertRequest{
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
