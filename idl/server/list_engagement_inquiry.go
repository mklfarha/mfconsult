package server

import (
	"context"
	"fmt"
	"log"
	//"encoding/json"

	engagement_inquirymodule "github.com/mklfarha/mfconsult/core/module/engagement_inquiry"
	"github.com/mklfarha/mfconsult/core/module/engagement_inquiry/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"

	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
	"go.einride.tech/aip/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BuildListEngagementInquiryRequest(ctx context.Context, request *pb.ListEngagementInquiryRequest) (types.ListRequest, *pagination.PageToken, error) {
	err := validatePageSizeForListEngagementInquiry(request)
	if err != nil {
		return types.ListRequest{}, nil, err
	}

	// Use pagination.PageToken for offset-based page tokens.
	pageToken, err := pagination.ParsePageToken(request)
	if err != nil {
		return types.ListRequest{}, nil, status.Errorf(codes.InvalidArgument, "invalid page token")
	}

	// parse filters
	declarations := engagement_inquiryDeclarations()
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

func (s *server) ListEngagementInquiry(ctx context.Context, request *pb.ListEngagementInquiryRequest) (*pb.ListEngagementInquiryResponse, error) {

	var err error
	req, pageToken, err := BuildListEngagementInquiryRequest(ctx, request)
	if err != nil {

		return nil, err
	}

	// Query the storage.
	var result types.ListResponse
	if request.GetSkipCache() {
		result, err = s.core.EngagementInquiry().List(ctx, req, engagement_inquirymodule.WithSkipCache())
	} else {
		result, err = s.core.EngagementInquiry().List(ctx, req)
	}
	if err != nil {

		return nil, err
	}

	// Build the response.
	response := &pb.ListEngagementInquiryResponse{
		EngagementInquiry: pbmapper.EngagementInquirySliceToProto(result.EngagementInquiry),
	}
	// Set the next page token.
	if result.HasNextPage {
		response.NextPageToken = pageToken.Next(request).String()
	}

	// Respond.
	return response, nil
}

func validatePageSizeForListEngagementInquiry(request *pb.ListEngagementInquiryRequest) error {
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

func engagement_inquiryDeclarations() *filtering.Declarations {
	declarations, err := filtering.NewDeclarations(
		filtering.DeclareStandardFunctions(),
		// boolean values
		filtering.DeclareIdent("true", filtering.TypeBool),
		filtering.DeclareIdent("false", filtering.TypeBool),

		//engagement_inquiry

		filtering.DeclareIdent("id", filtering.TypeString),

		filtering.DeclareIdent("client_id", filtering.TypeString),

		filtering.DeclareIdent("name", filtering.TypeString),

		filtering.DeclareIdent("email", filtering.TypeString),

		filtering.DeclareIdent("phone", filtering.TypeString),

		filtering.DeclareIdent("company", filtering.TypeString),

		filtering.DeclareIdent("project_summary", filtering.TypeString),

		filtering.DeclareIdent("why_more_than_session", filtering.TypeString),

		filtering.DeclareIdent("scope_details", filtering.TypeString),

		filtering.DeclareIdent("budget_range", filtering.TypeString),

		filtering.DeclareIdent("timeline", filtering.TypeString),

		filtering.DeclareEnumIdent("status", pb.InquiryStatus(0).Type()),

		filtering.DeclareIdent("review_notes", filtering.TypeString),

		filtering.DeclareIdent("created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("updated_at", filtering.TypeTimestamp),
	)
	if err != nil {
		log.Printf("error creating declarations: %v", err)
	}
	return declarations
}
