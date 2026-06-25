package server

import (
	"context"
	"fmt"
	"log"
	//"encoding/json"

	magic_linkmodule "github.com/mklfarha/mfconsult/core/module/magic_link"
	"github.com/mklfarha/mfconsult/core/module/magic_link/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"

	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
	"go.einride.tech/aip/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BuildListMagicLinkRequest(ctx context.Context, request *pb.ListMagicLinkRequest) (types.ListRequest, *pagination.PageToken, error) {
	err := validatePageSizeForListMagicLink(request)
	if err != nil {
		return types.ListRequest{}, nil, err
	}

	// Use pagination.PageToken for offset-based page tokens.
	pageToken, err := pagination.ParsePageToken(request)
	if err != nil {
		return types.ListRequest{}, nil, status.Errorf(codes.InvalidArgument, "invalid page token")
	}

	// parse filters
	declarations := magic_linkDeclarations()
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

func (s *server) ListMagicLink(ctx context.Context, request *pb.ListMagicLinkRequest) (*pb.ListMagicLinkResponse, error) {

	var err error
	req, pageToken, err := BuildListMagicLinkRequest(ctx, request)
	if err != nil {

		return nil, err
	}

	// Query the storage.
	var result types.ListResponse
	if request.GetSkipCache() {
		result, err = s.core.MagicLink().List(ctx, req, magic_linkmodule.WithSkipCache())
	} else {
		result, err = s.core.MagicLink().List(ctx, req)
	}
	if err != nil {

		return nil, err
	}

	// Build the response.
	response := &pb.ListMagicLinkResponse{
		MagicLink: pbmapper.MagicLinkSliceToProto(result.MagicLink),
	}
	// Set the next page token.
	if result.HasNextPage {
		response.NextPageToken = pageToken.Next(request).String()
	}

	// Respond.
	return response, nil
}

func validatePageSizeForListMagicLink(request *pb.ListMagicLinkRequest) error {
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

func magic_linkDeclarations() *filtering.Declarations {
	declarations, err := filtering.NewDeclarations(
		filtering.DeclareStandardFunctions(),
		// boolean values
		filtering.DeclareIdent("true", filtering.TypeBool),
		filtering.DeclareIdent("false", filtering.TypeBool),

		//magic_link

		filtering.DeclareIdent("id", filtering.TypeString),

		filtering.DeclareIdent("client_id", filtering.TypeString),

		filtering.DeclareIdent("email", filtering.TypeString),

		filtering.DeclareIdent("token", filtering.TypeString),

		filtering.DeclareEnumIdent("purpose", pb.MagicLinkPurpose(0).Type()),

		filtering.DeclareIdent("expires_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("consumed_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("created_ip", filtering.TypeString),
	)
	if err != nil {
		log.Printf("error creating declarations: %v", err)
	}
	return declarations
}
