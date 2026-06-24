package server

import (
	"context"
	"errors"
	booking_recapmodule "github.com/mklfarha/mfconsult/core/module/booking_recap"
	"github.com/mklfarha/mfconsult/core/module/booking_recap/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"
	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"
)

func (s *server) CreateBookingRecap(ctx context.Context, req *pb.CreateBookingRecapRequest) (*pb.BookingRecap, error) {
	res, err := s.core.BookingRecap().Insert(ctx, types.UpsertRequest{
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
