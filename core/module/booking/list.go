package booking

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mklfarha/mfconsult/core/module/booking/types"
	repogen "github.com/mklfarha/mfconsult/core/repository/gen"
	main_entity "github.com/mklfarha/mfconsult/entity/booking"

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
		main_entity.Booking{},
		false)
	if err != nil {

		return types.ListResponse{}, err
	}

	resolvedOpts := applyAllOptions(opts)
	cacheKey := fmt.Sprintf("ListBooking:%v", request)
	if !resolvedOpts.SkipCache {
		if cached, found := m.cache.Get(cacheKey); found {
			return cached.(types.ListResponse), nil
		}
	}

	v, listErr, _ := m.sg.Do(cacheKey, func() (any, error) {
		var rows *sql.Rows
		rows, err = m.repository.DB.QueryContext(ctx, query)
		if err != nil {

			m.logger.Error("error in executing query for ListBooking", zap.String("query", query), zap.Error(err))
			return types.ListResponse{}, err
		}
		defer rows.Close()
		var scanGetters []func(*repogen.Booking) any
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

		var items []repogen.Booking
		for rows.Next() {
			var i repogen.Booking
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
			Booking:     mapModelsToEntities(items),
			HasNextPage: hasNextPage,
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

	"status",

	"review_decision",

	"reviewed_at",

	"decline_reason",

	"intake",

	"payment",

	"scheduling",

	"terms_version",

	"terms_accepted_at",

	"terms_accepted_ip",

	"created_at",

	"updated_at",
}

var listFieldRegistry = map[string]func(*repogen.Booking) any{

	"id": func(i *repogen.Booking) any { return &i.ID },

	"client_id": func(i *repogen.Booking) any { return &i.ClientId },

	"status": func(i *repogen.Booking) any { return &i.Status },

	"review_decision": func(i *repogen.Booking) any { return &i.ReviewDecision },

	"reviewed_at": func(i *repogen.Booking) any { return &i.ReviewedAt },

	"decline_reason": func(i *repogen.Booking) any { return &i.DeclineReason },

	"intake": func(i *repogen.Booking) any { return &i.Intake },

	"payment": func(i *repogen.Booking) any { return &i.Payment },

	"scheduling": func(i *repogen.Booking) any { return &i.Scheduling },

	"terms_version": func(i *repogen.Booking) any { return &i.TermsVersion },

	"terms_accepted_at": func(i *repogen.Booking) any { return &i.TermsAcceptedAt },

	"terms_accepted_ip": func(i *repogen.Booking) any { return &i.TermsAcceptedIp },

	"created_at": func(i *repogen.Booking) any { return &i.CreatedAt },

	"updated_at": func(i *repogen.Booking) any { return &i.UpdatedAt },
}
