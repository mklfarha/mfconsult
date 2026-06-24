package webhook_event

import (
	"context"
	"fmt"

	"github.com/mklfarha/mfconsult/core/module/webhook_event/types"
)

func (m *module) FetchWebhookEventById(
	ctx context.Context,
	req types.FetchWebhookEventByIdRequest,
	opts ...Option,
) (types.FetchWebhookEventByIdResponse, error) {

	resolvedOpts := applyAllOptions(opts)
	cacheKey := fmt.Sprintf("FetchWebhookEventById:%v", req)
	if !resolvedOpts.SkipCache {
		if cached, found := m.cache.Get(cacheKey); found {
			return cached.(types.FetchWebhookEventByIdResponse), nil
		}
	}
	v, err, _ := m.sg.Do(cacheKey, func() (any, error) {
		models, err := m.repository.Queries.FetchWebhookEventById(
			ctx,
			req.ID,
		)
		if err != nil {

			return types.FetchWebhookEventByIdResponse{}, err
		}
		return types.FetchWebhookEventByIdResponse{
			Results: mapModelsToEntities(models),
		}, nil
	})
	if err != nil {
		return types.FetchWebhookEventByIdResponse{}, err
	}
	result := v.(types.FetchWebhookEventByIdResponse)
	if !resolvedOpts.SkipCache {
		m.cache.Set(cacheKey, result, 0)
	}
	return result, nil

}
