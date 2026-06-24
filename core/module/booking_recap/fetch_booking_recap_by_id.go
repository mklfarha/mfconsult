package booking_recap

import (
	"context"
	"fmt"

	"github.com/mklfarha/mfconsult/core/module/booking_recap/types"
)

func (m *module) FetchBookingRecapById(
	ctx context.Context,
	req types.FetchBookingRecapByIdRequest,
	opts ...Option,
) (types.FetchBookingRecapByIdResponse, error) {

	resolvedOpts := applyAllOptions(opts)
	cacheKey := fmt.Sprintf("FetchBookingRecapById:%v", req)
	if !resolvedOpts.SkipCache {
		if cached, found := m.cache.Get(cacheKey); found {
			return cached.(types.FetchBookingRecapByIdResponse), nil
		}
	}
	v, err, _ := m.sg.Do(cacheKey, func() (any, error) {
		models, err := m.repository.Queries.FetchBookingRecapById(
			ctx,
			req.ID.String(),
		)
		if err != nil {

			return types.FetchBookingRecapByIdResponse{}, err
		}
		return types.FetchBookingRecapByIdResponse{
			Results: mapModelsToEntities(models),
		}, nil
	})
	if err != nil {
		return types.FetchBookingRecapByIdResponse{}, err
	}
	result := v.(types.FetchBookingRecapByIdResponse)
	if !resolvedOpts.SkipCache {
		m.cache.Set(cacheKey, result, 0)
	}
	return result, nil

}
