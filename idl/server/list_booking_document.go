package server

import (
	"context"
	"fmt"
	"log"
	//"encoding/json"

	booking_documentmodule "github.com/mklfarha/mfconsult/core/module/booking_document"
	"github.com/mklfarha/mfconsult/core/module/booking_document/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"

	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
	"go.einride.tech/aip/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BuildListBookingDocumentRequest(ctx context.Context, request *pb.ListBookingDocumentRequest) (types.ListRequest, *pagination.PageToken, error) {
	err := validatePageSizeForListBookingDocument(request)
	if err != nil {
		return types.ListRequest{}, nil, err
	}

	// Use pagination.PageToken for offset-based page tokens.
	pageToken, err := pagination.ParsePageToken(request)
	if err != nil {
		return types.ListRequest{}, nil, status.Errorf(codes.InvalidArgument, "invalid page token")
	}

	// parse filters
	declarations := booking_documentDeclarations()
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

func (s *server) ListBookingDocument(ctx context.Context, request *pb.ListBookingDocumentRequest) (*pb.ListBookingDocumentResponse, error) {

	var err error
	req, pageToken, err := BuildListBookingDocumentRequest(ctx, request)
	if err != nil {

		return nil, err
	}

	// Query the storage.
	var result types.ListResponse
	if request.GetSkipCache() {
		result, err = s.core.BookingDocument().List(ctx, req, booking_documentmodule.WithSkipCache())
	} else {
		result, err = s.core.BookingDocument().List(ctx, req)
	}
	if err != nil {

		return nil, err
	}

	// Build the response.
	response := &pb.ListBookingDocumentResponse{
		BookingDocument: pbmapper.BookingDocumentSliceToProto(result.BookingDocument),
	}
	// Set the next page token.
	if result.HasNextPage {
		response.NextPageToken = pageToken.Next(request).String()
	}

	// Respond.
	return response, nil
}

func validatePageSizeForListBookingDocument(request *pb.ListBookingDocumentRequest) error {
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

func booking_documentDeclarations() *filtering.Declarations {
	declarations, err := filtering.NewDeclarations(
		filtering.DeclareStandardFunctions(),
		// boolean values
		filtering.DeclareIdent("true", filtering.TypeBool),
		filtering.DeclareIdent("false", filtering.TypeBool),

		//booking_document

		filtering.DeclareIdent("id", filtering.TypeString),

		filtering.DeclareIdent("booking_id", filtering.TypeString),

		filtering.DeclareEnumIdent("kind", pb.DocumentKind(0).Type()),

		filtering.DeclareIdent("url", filtering.TypeString),

		filtering.DeclareIdent("label", filtering.TypeString),

		filtering.DeclareIdent("purge_after", filtering.TypeTimestamp),

		filtering.DeclareIdent("created_at", filtering.TypeTimestamp),
	)
	if err != nil {
		log.Printf("error creating declarations: %v", err)
	}
	return declarations
}
