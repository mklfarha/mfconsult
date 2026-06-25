package engagement_agreement

import (
	"context"
	"fmt"

	"github.com/mklfarha/mfconsult/core/module/engagement_agreement/types"
)

func (m *module) FetchEngagementAgreementById(
	ctx context.Context,
	req types.FetchEngagementAgreementByIdRequest,
	opts ...Option,
) (types.FetchEngagementAgreementByIdResponse, error) {

	resolvedOpts := applyAllOptions(opts)
	cacheKey := fmt.Sprintf("FetchEngagementAgreementById:%v", req)
	if !resolvedOpts.SkipCache {
		if cached, found := m.cache.Get(cacheKey); found {
			return cached.(types.FetchEngagementAgreementByIdResponse), nil
		}
	}
	v, err, _ := m.sg.Do(cacheKey, func() (any, error) {
		models, err := m.repository.Queries.FetchEngagementAgreementById(
			ctx,
			req.ID.String(),
		)
		if err != nil {

			return types.FetchEngagementAgreementByIdResponse{}, err
		}
		return types.FetchEngagementAgreementByIdResponse{
			Results: mapModelsToEntities(models),
		}, nil
	})
	if err != nil {
		return types.FetchEngagementAgreementByIdResponse{}, err
	}
	result := v.(types.FetchEngagementAgreementByIdResponse)
	if !resolvedOpts.SkipCache {
		m.cache.Set(cacheKey, result, 0)
	}
	return result, nil

}
