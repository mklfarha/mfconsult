package engagement_inquiry

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mklfarha/mfconsult/core/module/engagement_inquiry/types"
	repogen "github.com/mklfarha/mfconsult/core/repository/gen"
	main_entity "github.com/mklfarha/mfconsult/entity/engagement_inquiry"

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
		main_entity.EngagementInquiry{},
		false)
	if err != nil {

		return types.ListResponse{}, err
	}

	resolvedOpts := applyAllOptions(opts)
	cacheKey := fmt.Sprintf("ListEngagementInquiry:%v", request)
	if !resolvedOpts.SkipCache {
		if cached, found := m.cache.Get(cacheKey); found {
			return cached.(types.ListResponse), nil
		}
	}

	v, listErr, _ := m.sg.Do(cacheKey, func() (any, error) {
		var rows *sql.Rows
		rows, err = m.repository.DB.QueryContext(ctx, query)
		if err != nil {

			m.logger.Error("error in executing query for ListEngagementInquiry", zap.String("query", query), zap.Error(err))
			return types.ListResponse{}, err
		}
		defer rows.Close()
		var scanGetters []func(*repogen.EngagementInquiry) any
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

		var items []repogen.EngagementInquiry
		for rows.Next() {
			var i repogen.EngagementInquiry
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
			EngagementInquiry: mapModelsToEntities(items),
			HasNextPage:       hasNextPage,
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

	"name",

	"email",

	"phone",

	"company",

	"project_summary",

	"why_more_than_session",

	"scope_details",

	"budget_range",

	"timeline",

	"status",

	"review_notes",

	"created_at",

	"updated_at",
}

var listFieldRegistry = map[string]func(*repogen.EngagementInquiry) any{

	"id": func(i *repogen.EngagementInquiry) any { return &i.ID },

	"client_id": func(i *repogen.EngagementInquiry) any { return &i.ClientId },

	"name": func(i *repogen.EngagementInquiry) any { return &i.Name },

	"email": func(i *repogen.EngagementInquiry) any { return &i.Email },

	"phone": func(i *repogen.EngagementInquiry) any { return &i.Phone },

	"company": func(i *repogen.EngagementInquiry) any { return &i.Company },

	"project_summary": func(i *repogen.EngagementInquiry) any { return &i.ProjectSummary },

	"why_more_than_session": func(i *repogen.EngagementInquiry) any { return &i.WhyMoreThanSession },

	"scope_details": func(i *repogen.EngagementInquiry) any { return &i.ScopeDetails },

	"budget_range": func(i *repogen.EngagementInquiry) any { return &i.BudgetRange },

	"timeline": func(i *repogen.EngagementInquiry) any { return &i.Timeline },

	"status": func(i *repogen.EngagementInquiry) any { return &i.Status },

	"review_notes": func(i *repogen.EngagementInquiry) any { return &i.ReviewNotes },

	"created_at": func(i *repogen.EngagementInquiry) any { return &i.CreatedAt },

	"updated_at": func(i *repogen.EngagementInquiry) any { return &i.UpdatedAt },
}
