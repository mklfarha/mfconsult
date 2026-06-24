package server

import (
	"context"
	"errors"
	bookingmodule "github.com/mklfarha/mfconsult/core/module/booking"
	"github.com/mklfarha/mfconsult/core/module/booking/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"

	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateBooking(ctx context.Context, req *pb.UpdateBookingRequest) (*pb.Booking, error) {

	if req.Booking.Id == "" {
		return nil, errors.New("please provide a valid ID to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req.GetBooking())
	if err != nil {

		return nil, err
	}

	isFull := fieldmask.IsFullReplacement(req.UpdateMask)

	if !isFull && req.UpdateMask != nil {

		if !strings.Contains(req.UpdateMask.String(), "id") {
			req.UpdateMask.Append(req.GetBooking(), "id")
		}

		pkEntity := pbmapper.BookingFromProto(req.GetBooking())
		existingRes, err := s.core.Booking().FetchBookingById(ctx,
			types.FetchBookingByIdRequest{
				ID: pkEntity.ID,
			},
			bookingmodule.WithSkipCache(),
		)
		if err != nil {

			return nil, err
		}
		if len(existingRes.Results) == 0 {
			return nil, errors.New("entity not found")
		}

		merged := pbmapper.BookingToProto(existingRes.Results[0])
		fieldmask.Update(req.UpdateMask, merged, req.GetBooking())
		req = &pb.UpdateBookingRequest{Booking: merged, UpdateMask: req.UpdateMask}
	}

	res, err := s.core.Booking().Update(ctx, types.UpsertRequest{
		Booking: pbmapper.BookingFromProto(req.GetBooking()),
	})
	if err != nil {

		return nil, err
	}

	fetchRes, err := s.core.Booking().FetchBookingById(ctx, types.FetchBookingByIdRequest(res), bookingmodule.WithSkipCache())
	if err != nil {

		return nil, err
	}

	if len(fetchRes.Results) == 0 {
		err := errors.New("error fetching entity")

		return nil, err
	}

	return pbmapper.BookingToProto(fetchRes.Results[0]), nil
}
