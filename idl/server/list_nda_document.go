package server

import (
	"context"
	"fmt"
	"log"
	//"encoding/json"

	nda_documentmodule "github.com/mklfarha/mfconsult/core/module/nda_document"
	"github.com/mklfarha/mfconsult/core/module/nda_document/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"

	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
	"go.einride.tech/aip/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BuildListNdaDocumentRequest(ctx context.Context, request *pb.ListNdaDocumentRequest) (types.ListRequest, *pagination.PageToken, error) {
	err := validatePageSizeForListNdaDocument(request)
	if err != nil {
		return types.ListRequest{}, nil, err
	}

	// Use pagination.PageToken for offset-based page tokens.
	pageToken, err := pagination.ParsePageToken(request)
	if err != nil {
		return types.ListRequest{}, nil, status.Errorf(codes.InvalidArgument, "invalid page token")
	}

	// parse filters
	declarations := nda_documentDeclarations()
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

func (s *server) ListNdaDocument(ctx context.Context, request *pb.ListNdaDocumentRequest) (*pb.ListNdaDocumentResponse, error) {

	var err error
	req, pageToken, err := BuildListNdaDocumentRequest(ctx, request)
	if err != nil {

		return nil, err
	}

	// Query the storage.
	var result types.ListResponse
	if request.GetSkipCache() {
		result, err = s.core.NdaDocument().List(ctx, req, nda_documentmodule.WithSkipCache())
	} else {
		result, err = s.core.NdaDocument().List(ctx, req)
	}
	if err != nil {

		return nil, err
	}

	// Build the response.
	response := &pb.ListNdaDocumentResponse{
		NdaDocument: pbmapper.NdaDocumentSliceToProto(result.NdaDocument),
	}
	// Set the next page token.
	if result.HasNextPage {
		response.NextPageToken = pageToken.Next(request).String()
	}

	// Respond.
	return response, nil
}

func validatePageSizeForListNdaDocument(request *pb.ListNdaDocumentRequest) error {
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

func nda_documentDeclarations() *filtering.Declarations {
	declarations, err := filtering.NewDeclarations(
		filtering.DeclareStandardFunctions(),
		// boolean values
		filtering.DeclareIdent("true", filtering.TypeBool),
		filtering.DeclareIdent("false", filtering.TypeBool),

		//nda_document

		filtering.DeclareIdent("id", filtering.TypeString),

		filtering.DeclareIdent("client_id", filtering.TypeString),

		filtering.DeclareIdent("url", filtering.TypeString),

		filtering.DeclareEnumIdent("status", pb.NdaStatus(0).Type()),

		filtering.DeclareIdent("signed_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("envelope_id", filtering.TypeString),

		filtering.DeclareIdent("certificate_url", filtering.TypeString),
	)
	if err != nil {
		log.Printf("error creating declarations: %v", err)
	}
	return declarations
}
