package engagement_inquiry

import (
	"context"
	"fmt"

	"github.com/mklfarha/mfconsult/core/module/engagement_inquiry/types"
)

func (m *module) FetchEngagementInquiryById(
	ctx context.Context,
	req types.FetchEngagementInquiryByIdRequest,
	opts ...Option,
) (types.FetchEngagementInquiryByIdResponse, error) {

	resolvedOpts := applyAllOptions(opts)
	cacheKey := fmt.Sprintf("FetchEngagementInquiryById:%v", req)
	if !resolvedOpts.SkipCache {
		if cached, found := m.cache.Get(cacheKey); found {
			return cached.(types.FetchEngagementInquiryByIdResponse), nil
		}
	}
	v, err, _ := m.sg.Do(cacheKey, func() (any, error) {
		models, err := m.repository.Queries.FetchEngagementInquiryById(
			ctx,
			req.ID.String(),
		)
		if err != nil {

			return types.FetchEngagementInquiryByIdResponse{}, err
		}
		return types.FetchEngagementInquiryByIdResponse{
			Results: mapModelsToEntities(models),
		}, nil
	})
	if err != nil {
		return types.FetchEngagementInquiryByIdResponse{}, err
	}
	result := v.(types.FetchEngagementInquiryByIdResponse)
	if !resolvedOpts.SkipCache {
		m.cache.Set(cacheKey, result, 0)
	}
	return result, nil

}
