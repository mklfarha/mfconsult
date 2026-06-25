package server

import (
	"context"
	"fmt"
	"log"
	//"encoding/json"

	engagement_agreementmodule "github.com/mklfarha/mfconsult/core/module/engagement_agreement"
	"github.com/mklfarha/mfconsult/core/module/engagement_agreement/types"
	pb "github.com/mklfarha/mfconsult/idl/gen"

	pbmapper "github.com/mklfarha/mfconsult/idl/mapper"

	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
	"go.einride.tech/aip/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BuildListEngagementAgreementRequest(ctx context.Context, request *pb.ListEngagementAgreementRequest) (types.ListRequest, *pagination.PageToken, error) {
	err := validatePageSizeForListEngagementAgreement(request)
	if err != nil {
		return types.ListRequest{}, nil, err
	}

	// Use pagination.PageToken for offset-based page tokens.
	pageToken, err := pagination.ParsePageToken(request)
	if err != nil {
		return types.ListRequest{}, nil, status.Errorf(codes.InvalidArgument, "invalid page token")
	}

	// parse filters
	declarations := engagement_agreementDeclarations()
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

func (s *server) ListEngagementAgreement(ctx context.Context, request *pb.ListEngagementAgreementRequest) (*pb.ListEngagementAgreementResponse, error) {

	var err error
	req, pageToken, err := BuildListEngagementAgreementRequest(ctx, request)
	if err != nil {

		return nil, err
	}

	// Query the storage.
	var result types.ListResponse
	if request.GetSkipCache() {
		result, err = s.core.EngagementAgreement().List(ctx, req, engagement_agreementmodule.WithSkipCache())
	} else {
		result, err = s.core.EngagementAgreement().List(ctx, req)
	}
	if err != nil {

		return nil, err
	}

	// Build the response.
	response := &pb.ListEngagementAgreementResponse{
		EngagementAgreement: pbmapper.EngagementAgreementSliceToProto(result.EngagementAgreement),
	}
	// Set the next page token.
	if result.HasNextPage {
		response.NextPageToken = pageToken.Next(request).String()
	}

	// Respond.
	return response, nil
}

func validatePageSizeForListEngagementAgreement(request *pb.ListEngagementAgreementRequest) error {
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

func engagement_agreementDeclarations() *filtering.Declarations {
	declarations, err := filtering.NewDeclarations(
		filtering.DeclareStandardFunctions(),
		// boolean values
		filtering.DeclareIdent("true", filtering.TypeBool),
		filtering.DeclareIdent("false", filtering.TypeBool),

		//engagement_agreement

		filtering.DeclareIdent("id", filtering.TypeString),

		filtering.DeclareIdent("client_id", filtering.TypeString),

		filtering.DeclareIdent("nda_url", filtering.TypeString),

		filtering.DeclareEnumIdent("status", pb.AgreementStatus(0).Type()),

		filtering.DeclareIdent("signed_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("envelope_id", filtering.TypeString),

		filtering.DeclareIdent("certificate_url", filtering.TypeString),

		filtering.DeclareIdent("contract_url", filtering.TypeString),

		filtering.DeclareIdent("engagement_inquiry_id", filtering.TypeString),

		filtering.DeclareIdent("updated_at", filtering.TypeTimestamp),
	)
	if err != nil {
		log.Printf("error creating declarations: %v", err)
	}
	return declarations
}
