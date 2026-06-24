package server

import (
	"context"
	"errors"
	booking_recapmodule "github.com/mklfarha/mfconsult/core/module/booking_recap"
	"github.com/mklfarha/mfconsult/core/module/booking_recap/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"

	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateBookingRecap(ctx context.Context, req *pb.UpdateBookingRecapRequest) (*pb.BookingRecap, error) {

	if req.BookingRecap.Id == "" {
		return nil, errors.New("please provide a valid ID to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req.GetBookingRecap())
	if err != nil {

		return nil, err
	}

	isFull := fieldmask.IsFullReplacement(req.UpdateMask)

	if !isFull && req.UpdateMask != nil {

		if !strings.Contains(req.UpdateMask.String(), "id") {
			req.UpdateMask.Append(req.GetBookingRecap(), "id")
		}

		pkEntity := pbmapper.BookingRecapFromProto(req.GetBookingRecap())
		existingRes, err := s.core.BookingRecap().FetchBookingRecapById(ctx,
			types.FetchBookingRecapByIdRequest{
				ID: pkEntity.ID,
			},
			booking_recapmodule.WithSkipCache(),
		)
		if err != nil {

			return nil, err
		}
		if len(existingRes.Results) == 0 {
			return nil, errors.New("entity not found")
		}

		merged := pbmapper.BookingRecapToProto(existingRes.Results[0])
		fieldmask.Update(req.UpdateMask, merged, req.GetBookingRecap())
		req = &pb.UpdateBookingRecapRequest{BookingRecap: merged, UpdateMask: req.UpdateMask}
	}

	res, err := s.core.BookingRecap().Update(ctx, types.UpsertRequest{
		BookingRecap: pbmapper.BookingRecapFromProto(req.GetBookingRecap()),
	})
	if err != nil {

		return nil, err
	}

	fetchRes, err := s.core.BookingRecap().FetchBookingRecapById(ctx, types.FetchBookingRecapByIdRequest(res), booking_recapmodule.WithSkipCache())
	if err != nil {

		return nil, err
	}

	if len(fetchRes.Results) == 0 {
		err := errors.New("error fetching entity")

		return nil, err
	}

	return pbmapper.BookingRecapToProto(fetchRes.Results[0]), nil
}
