package server

import (
	"context"
	"errors"
	booking_documentmodule "github.com/mklfarha/mfconsult/core/module/booking_document"
	"github.com/mklfarha/mfconsult/core/module/booking_document/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"

	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateBookingDocument(ctx context.Context, req *pb.UpdateBookingDocumentRequest) (*pb.BookingDocument, error) {

	if req.BookingDocument.Id == "" {
		return nil, errors.New("please provide a valid ID to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req.GetBookingDocument())
	if err != nil {

		return nil, err
	}

	isFull := fieldmask.IsFullReplacement(req.UpdateMask)

	if !isFull && req.UpdateMask != nil {

		if !strings.Contains(req.UpdateMask.String(), "id") {
			req.UpdateMask.Append(req.GetBookingDocument(), "id")
		}

		pkEntity := pbmapper.BookingDocumentFromProto(req.GetBookingDocument())
		existingRes, err := s.core.BookingDocument().FetchBookingDocumentById(ctx,
			types.FetchBookingDocumentByIdRequest{
				ID: pkEntity.ID,
			},
			booking_documentmodule.WithSkipCache(),
		)
		if err != nil {

			return nil, err
		}
		if len(existingRes.Results) == 0 {
			return nil, errors.New("entity not found")
		}

		merged := pbmapper.BookingDocumentToProto(existingRes.Results[0])
		fieldmask.Update(req.UpdateMask, merged, req.GetBookingDocument())
		req = &pb.UpdateBookingDocumentRequest{BookingDocument: merged, UpdateMask: req.UpdateMask}
	}

	res, err := s.core.BookingDocument().Update(ctx, types.UpsertRequest{
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
