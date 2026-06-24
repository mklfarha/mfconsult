package booking_document

import (
	"context"
	"fmt"

	"github.com/mklfarha/mfconsult/core/module/booking_document/types"
)

func (m *module) FetchBookingDocumentById(
	ctx context.Context,
	req types.FetchBookingDocumentByIdRequest,
	opts ...Option,
) (types.FetchBookingDocumentByIdResponse, error) {

	resolvedOpts := applyAllOptions(opts)
	cacheKey := fmt.Sprintf("FetchBookingDocumentById:%v", req)
	if !resolvedOpts.SkipCache {
		if cached, found := m.cache.Get(cacheKey); found {
			return cached.(types.FetchBookingDocumentByIdResponse), nil
		}
	}
	v, err, _ := m.sg.Do(cacheKey, func() (any, error) {
		models, err := m.repository.Queries.FetchBookingDocumentById(
			ctx,
			req.ID.String(),
		)
		if err != nil {

			return types.FetchBookingDocumentByIdResponse{}, err
		}
		return types.FetchBookingDocumentByIdResponse{
			Results: mapModelsToEntities(models),
		}, nil
	})
	if err != nil {
		return types.FetchBookingDocumentByIdResponse{}, err
	}
	result := v.(types.FetchBookingDocumentByIdResponse)
	if !resolvedOpts.SkipCache {
		m.cache.Set(cacheKey, result, 0)
	}
	return result, nil

}
