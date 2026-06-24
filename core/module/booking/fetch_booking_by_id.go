package booking

import (
	"context"
	"fmt"

	"github.com/mklfarha/mfconsult/core/module/booking/types"
)

func (m *module) FetchBookingById(
	ctx context.Context,
	req types.FetchBookingByIdRequest,
	opts ...Option,
) (types.FetchBookingByIdResponse, error) {

	resolvedOpts := applyAllOptions(opts)
	cacheKey := fmt.Sprintf("FetchBookingById:%v", req)
	if !resolvedOpts.SkipCache {
		if cached, found := m.cache.Get(cacheKey); found {
			return cached.(types.FetchBookingByIdResponse), nil
		}
	}
	v, err, _ := m.sg.Do(cacheKey, func() (any, error) {
		models, err := m.repository.Queries.FetchBookingById(
			ctx,
			req.ID.String(),
		)
		if err != nil {

			return types.FetchBookingByIdResponse{}, err
		}
		return types.FetchBookingByIdResponse{
			Results: mapModelsToEntities(models),
		}, nil
	})
	if err != nil {
		return types.FetchBookingByIdResponse{}, err
	}
	result := v.(types.FetchBookingByIdResponse)
	if !resolvedOpts.SkipCache {
		m.cache.Set(cacheKey, result, 0)
	}
	return result, nil

}
