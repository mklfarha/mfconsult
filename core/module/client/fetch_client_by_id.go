package client

import (
	"context"
	"fmt"

	"github.com/mklfarha/mfconsult/core/module/client/types"
)

func (m *module) FetchClientById(
	ctx context.Context,
	req types.FetchClientByIdRequest,
	opts ...Option,
) (types.FetchClientByIdResponse, error) {

	resolvedOpts := applyAllOptions(opts)
	cacheKey := fmt.Sprintf("FetchClientById:%v", req)
	if !resolvedOpts.SkipCache {
		if cached, found := m.cache.Get(cacheKey); found {
			return cached.(types.FetchClientByIdResponse), nil
		}
	}
	v, err, _ := m.sg.Do(cacheKey, func() (any, error) {
		models, err := m.repository.Queries.FetchClientById(
			ctx,
			req.ID.String(),
		)
		if err != nil {

			return types.FetchClientByIdResponse{}, err
		}
		return types.FetchClientByIdResponse{
			Results: mapModelsToEntities(models),
		}, nil
	})
	if err != nil {
		return types.FetchClientByIdResponse{}, err
	}
	result := v.(types.FetchClientByIdResponse)
	if !resolvedOpts.SkipCache {
		m.cache.Set(cacheKey, result, 0)
	}
	return result, nil

}
