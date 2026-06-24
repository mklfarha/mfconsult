package server

import (
	"context"
	"errors"
	bookingmodule "github.com/mklfarha/mfconsult/core/module/booking"
	"github.com/mklfarha/mfconsult/core/module/booking/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"
)

func (s *server) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.Booking, error) {
	res, err := s.core.Booking().Insert(ctx, types.UpsertRequest{
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
