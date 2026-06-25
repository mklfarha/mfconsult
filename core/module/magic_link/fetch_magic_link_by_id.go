package magic_link

import (
	"context"
	"fmt"

	"github.com/mklfarha/mfconsult/core/module/magic_link/types"
)

func (m *module) FetchMagicLinkById(
	ctx context.Context,
	req types.FetchMagicLinkByIdRequest,
	opts ...Option,
) (types.FetchMagicLinkByIdResponse, error) {

	resolvedOpts := applyAllOptions(opts)
	cacheKey := fmt.Sprintf("FetchMagicLinkById:%v", req)
	if !resolvedOpts.SkipCache {
		if cached, found := m.cache.Get(cacheKey); found {
			return cached.(types.FetchMagicLinkByIdResponse), nil
		}
	}
	v, err, _ := m.sg.Do(cacheKey, func() (any, error) {
		models, err := m.repository.Queries.FetchMagicLinkById(
			ctx,
			req.ID.String(),
		)
		if err != nil {

			return types.FetchMagicLinkByIdResponse{}, err
		}
		return types.FetchMagicLinkByIdResponse{
			Results: mapModelsToEntities(models),
		}, nil
	})
	if err != nil {
		return types.FetchMagicLinkByIdResponse{}, err
	}
	result := v.(types.FetchMagicLinkByIdResponse)
	if !resolvedOpts.SkipCache {
		m.cache.Set(cacheKey, result, 0)
	}
	return result, nil

}
