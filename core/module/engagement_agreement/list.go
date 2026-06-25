package engagement_agreement

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mklfarha/mfconsult/core/module/engagement_agreement/types"
	repogen "github.com/mklfarha/mfconsult/core/repository/gen"
	main_entity "github.com/mklfarha/mfconsult/entity/engagement_agreement"

	"go.uber.org/zap"
	"slices"
)

func (m *module) List(ctx context.Context,
	request types.ListRequest,
	opts ...Option) (types.ListResponse, error) {

	reqPlusOne := request
	reqPlusOne.PageSize = request.PageSize + 1

	query, err := m.repository.BuildListEntityQuery(
		ctx,
		reqPlusOne,
		main_entity.EngagementAgreement{},
		false)
	if err != nil {

		return types.ListResponse{}, err
	}

	resolvedOpts := applyAllOptions(opts)
	cacheKey := fmt.Sprintf("ListEngagementAgreement:%v", request)
	if !resolvedOpts.SkipCache {
		if cached, found := m.cache.Get(cacheKey); found {
			return cached.(types.ListResponse), nil
		}
	}

	v, listErr, _ := m.sg.Do(cacheKey, func() (any, error) {
		var rows *sql.Rows
		rows, err = m.repository.DB.QueryContext(ctx, query)
		if err != nil {

			m.logger.Error("error in executing query for ListEngagementAgreement", zap.String("query", query), zap.Error(err))
			return types.ListResponse{}, err
		}
		defer rows.Close()
		var scanGetters []func(*repogen.EngagementAgreement) any
		if len(request.GetIncludeFields()) > 0 {
			for _, f := range listFields {
				if slices.Contains(request.GetIncludeFields(), f) {
					scanGetters = append(scanGetters, listFieldRegistry[f])
				}
			}
		} else if len(request.GetExcludeFields()) > 0 {
			for _, f := range listFields {
				if !slices.Contains(request.GetExcludeFields(), f) {
					scanGetters = append(scanGetters, listFieldRegistry[f])
				}
			}
		} else {
			for _, f := range listFields {
				scanGetters = append(scanGetters, listFieldRegistry[f])
			}
		}

		var items []repogen.EngagementAgreement
		for rows.Next() {
			var i repogen.EngagementAgreement
			fields := make([]any, 0, len(scanGetters))
			for _, getter := range scanGetters {
				fields = append(fields, getter(&i))
			}
			if err := rows.Scan(fields...); err != nil {

				return types.ListResponse{}, err
			}
			items = append(items, i)
		}
		if err := rows.Close(); err != nil {

			return types.ListResponse{}, err
		}
		if err := rows.Err(); err != nil {

			return types.ListResponse{}, err
		}

		hasNextPage := false
		if len(items) > int(request.PageSize) {
			hasNextPage = true
			items = items[:request.PageSize]
		}

		return types.ListResponse{
			EngagementAgreement: mapModelsToEntities(items),
			HasNextPage:         hasNextPage,
		}, nil
	})
	if listErr != nil {
		return types.ListResponse{}, listErr
	}
	listResponse := v.(types.ListResponse)
	if !resolvedOpts.SkipCache {
		m.cache.Set(cacheKey, listResponse, 0)
	}
	return listResponse, nil
}

var listFields = []string{

	"id",

	"client_id",

	"nda_url",

	"status",

	"signed_at",

	"created_at",

	"envelope_id",

	"certificate_url",

	"contract_url",

	"engagement_inquiry_id",

	"updated_at",
}

var listFieldRegistry = map[string]func(*repogen.EngagementAgreement) any{

	"id": func(i *repogen.EngagementAgreement) any { return &i.ID },

	"client_id": func(i *repogen.EngagementAgreement) any { return &i.ClientId },

	"nda_url": func(i *repogen.EngagementAgreement) any { return &i.NdaURL },

	"status": func(i *repogen.EngagementAgreement) any { return &i.Status },

	"signed_at": func(i *repogen.EngagementAgreement) any { return &i.SignedAt },

	"created_at": func(i *repogen.EngagementAgreement) any { return &i.CreatedAt },

	"envelope_id": func(i *repogen.EngagementAgreement) any { return &i.EnvelopeId },

	"certificate_url": func(i *repogen.EngagementAgreement) any { return &i.CertificateURL },

	"contract_url": func(i *repogen.EngagementAgreement) any { return &i.ContractURL },

	"engagement_inquiry_id": func(i *repogen.EngagementAgreement) any { return &i.EngagementInquiryId },

	"updated_at": func(i *repogen.EngagementAgreement) any { return &i.UpdatedAt },
}
