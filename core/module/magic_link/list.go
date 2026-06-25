package magic_link

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mklfarha/mfconsult/core/module/magic_link/types"
	repogen "github.com/mklfarha/mfconsult/core/repository/gen"
	main_entity "github.com/mklfarha/mfconsult/entity/magic_link"

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
		main_entity.MagicLink{},
		false)
	if err != nil {

		return types.ListResponse{}, err
	}

	resolvedOpts := applyAllOptions(opts)
	cacheKey := fmt.Sprintf("ListMagicLink:%v", request)
	if !resolvedOpts.SkipCache {
		if cached, found := m.cache.Get(cacheKey); found {
			return cached.(types.ListResponse), nil
		}
	}

	v, listErr, _ := m.sg.Do(cacheKey, func() (any, error) {
		var rows *sql.Rows
		rows, err = m.repository.DB.QueryContext(ctx, query)
		if err != nil {

			m.logger.Error("error in executing query for ListMagicLink", zap.String("query", query), zap.Error(err))
			return types.ListResponse{}, err
		}
		defer rows.Close()
		var scanGetters []func(*repogen.MagicLink) any
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

		var items []repogen.MagicLink
		for rows.Next() {
			var i repogen.MagicLink
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
			MagicLink:   mapModelsToEntities(items),
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

	"email",

	"token",

	"purpose",

	"expires_at",

	"consumed_at",

	"created_at",

	"created_ip",
}

var listFieldRegistry = map[string]func(*repogen.MagicLink) any{

	"id": func(i *repogen.MagicLink) any { return &i.ID },

	"client_id": func(i *repogen.MagicLink) any { return &i.ClientId },

	"email": func(i *repogen.MagicLink) any { return &i.Email },

	"token": func(i *repogen.MagicLink) any { return &i.Token },

	"purpose": func(i *repogen.MagicLink) any { return &i.Purpose },

	"expires_at": func(i *repogen.MagicLink) any { return &i.ExpiresAt },

	"consumed_at": func(i *repogen.MagicLink) any { return &i.ConsumedAt },

	"created_at": func(i *repogen.MagicLink) any { return &i.CreatedAt },

	"created_ip": func(i *repogen.MagicLink) any { return &i.CreatedIp },
}
