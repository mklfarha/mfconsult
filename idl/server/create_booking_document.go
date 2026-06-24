package server

import (
	"context"
	"errors"
	booking_documentmodule "github.com/mklfarha/mfconsult/core/module/booking_document"
	"github.com/mklfarha/mfconsult/core/module/booking_document/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"
)

func (s *server) CreateBookingDocument(ctx context.Context, req *pb.CreateBookingDocumentRequest) (*pb.BookingDocument, error) {
	res, err := s.core.BookingDocument().Insert(ctx, types.UpsertRequest{
		BookingDocument: pbmapper.BookingDocumentFromProto(req.GetBookingDocument()),
	})
	if err != nil {

		return nil, err
	}

	fetchRes, err := s.core.BookingDocument().FetchBookingDocumentById(ctx, types.FetchBookingDocumentByIdRequest(res), booking_documentmodule.WithSkipCache())
	if err != nil {

		return nil, err
	}

	if len(fetchRes.Results) == 0 {
		err := errors.New("error fetching entity")

		return nil, err
	}

	return pbmapper.BookingDocumentToProto(fetchRes.Results[0]), nil
}
