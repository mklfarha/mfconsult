package server

import (
	"context"
	"fmt"
	"log"
	//"encoding/json"

	bookingmodule "github.com/mklfarha/mfconsult/core/module/booking"
	"github.com/mklfarha/mfconsult/core/module/booking/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"

	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
	"go.einride.tech/aip/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BuildListBookingRequest(ctx context.Context, request *pb.ListBookingRequest) (types.ListRequest, *pagination.PageToken, error) {
	err := validatePageSizeForListBooking(request)
	if err != nil {
		return types.ListRequest{}, nil, err
	}

	// Use pagination.PageToken for offset-based page tokens.
	pageToken, err := pagination.ParsePageToken(request)
	if err != nil {
		return types.ListRequest{}, nil, status.Errorf(codes.InvalidArgument, "invalid page token")
	}

	// parse filters
	declarations := bookingDeclarations()
	filter, err := filtering.ParseFilter(request, declarations)
	if err != nil {
		return types.ListRequest{}, nil, fmt.Errorf("error parsing filter: %w", err)
	}

	/* // enable for debugging
	if filter.CheckedExpr != nil {
		b, _ := json.Marshal(filter.CheckedExpr.Expr)
		log.Printf("filtering: %v \n", string(b))
	}
	*/

	orderBy, err := ordering.ParseOrderBy(request)
	if err != nil {
		return types.ListRequest{}, nil, fmt.Errorf("error parsing order by: %w", err)
	}

	/* // enable for debugging
	if orderBy.Fields != nil {
		b, _ := json.Marshal(orderBy.Fields)
		log.Printf("ordering: %v \n", string(b))
	}
	*/

	return types.ListRequest{
		Offset:                pageToken.Offset,
		PageSize:              request.GetPageSize(),
		Filter:                filter,
		FilteringDeclarations: declarations,
		OrderBy:               orderBy,
		IncludeFields:         request.GetIncludeFields(),
		ExcludeFields:         request.GetExcludeFields(),
	}, &pageToken, nil
}

func (s *server) ListBooking(ctx context.Context, request *pb.ListBookingRequest) (*pb.ListBookingResponse, error) {

	var err error
	req, pageToken, err := BuildListBookingRequest(ctx, request)
	if err != nil {

		return nil, err
	}

	// Query the storage.
	var result types.ListResponse
	if request.GetSkipCache() {
		result, err = s.core.Booking().List(ctx, req, bookingmodule.WithSkipCache())
	} else {
		result, err = s.core.Booking().List(ctx, req)
	}
	if err != nil {

		return nil, err
	}

	// Build the response.
	response := &pb.ListBookingResponse{
		Booking: pbmapper.BookingSliceToProto(result.Booking),
	}
	// Set the next page token.
	if result.HasNextPage {
		response.NextPageToken = pageToken.Next(request).String()
	}

	// Respond.
	return response, nil
}

func validatePageSizeForListBooking(request *pb.ListBookingRequest) error {
	// Handle request constraints.
	const (
		defaultPageSize = 10
	)
	switch {
	case request.PageSize < 0:
		return status.Errorf(codes.InvalidArgument, "page size is negative")
	case request.PageSize == 0:
		request.PageSize = defaultPageSize
	}
	return nil
}

func bookingDeclarations() *filtering.Declarations {
	declarations, err := filtering.NewDeclarations(
		filtering.DeclareStandardFunctions(),
		// boolean values
		filtering.DeclareIdent("true", filtering.TypeBool),
		filtering.DeclareIdent("false", filtering.TypeBool),

		//booking_intake

		filtering.DeclareIdent("booking_intake.reason", filtering.TypeString),

		filtering.DeclareEnumIdent("booking_intake.help_topic", pb.HelpTopic(0).Type()),

		filtering.DeclareIdent("booking_intake.help_details", filtering.TypeString),

		filtering.DeclareIdent("booking_intake.stack_details", filtering.TypeString),

		filtering.DeclareIdent("booking_intake.prep_notes", filtering.TypeString),

		//booking_payment

		filtering.DeclareIdent("booking_payment.amount_cents", filtering.TypeInt),

		filtering.DeclareIdent("booking_payment.currency", filtering.TypeString),

		filtering.DeclareEnumIdent("booking_payment.payment_status", pb.PaymentStatus(0).Type()),

		filtering.DeclareIdent("booking_payment.stripe_ref", filtering.TypeString),

		//booking_scheduling

		filtering.DeclareIdent("booking_scheduling.slot_start", filtering.TypeTimestamp),

		filtering.DeclareIdent("booking_scheduling.slot_end", filtering.TypeTimestamp),

		filtering.DeclareIdent("booking_scheduling.scheduler_booking_id", filtering.TypeString),

		filtering.DeclareIdent("booking_scheduling.video_url", filtering.TypeString),

		//booking

		filtering.DeclareIdent("id", filtering.TypeString),

		filtering.DeclareIdent("client_id", filtering.TypeString),

		filtering.DeclareEnumIdent("status", pb.BookingStatus(0).Type()),

		filtering.DeclareEnumIdent("review_decision", pb.ReviewDecision(0).Type()),

		filtering.DeclareIdent("reviewed_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("decline_reason", filtering.TypeString),

		filtering.DeclareIdent("pay_link_token", filtering.TypeString),

		filtering.DeclareIdent("pay_link_expires_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("portal_token", filtering.TypeString),

		filtering.DeclareIdent("terms_version", filtering.TypeString),

		filtering.DeclareIdent("terms_accepted_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("terms_accepted_ip", filtering.TypeString),

		filtering.DeclareIdent("created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("updated_at", filtering.TypeTimestamp),
	)
	if err != nil {
		log.Printf("error creating declarations: %v", err)
	}
	return declarations
}
