package nda_document

import (
	"context"
	"fmt"

	"github.com/mklfarha/mfconsult/core/module/nda_document/types"
)

func (m *module) FetchNdaDocumentById(
	ctx context.Context,
	req types.FetchNdaDocumentByIdRequest,
	opts ...Option,
) (types.FetchNdaDocumentByIdResponse, error) {

	resolvedOpts := applyAllOptions(opts)
	cacheKey := fmt.Sprintf("FetchNdaDocumentById:%v", req)
	if !resolvedOpts.SkipCache {
		if cached, found := m.cache.Get(cacheKey); found {
			return cached.(types.FetchNdaDocumentByIdResponse), nil
		}
	}
	v, err, _ := m.sg.Do(cacheKey, func() (any, error) {
		models, err := m.repository.Queries.FetchNdaDocumentById(
			ctx,
			req.ID.String(),
		)
		if err != nil {

			return types.FetchNdaDocumentByIdResponse{}, err
		}
		return types.FetchNdaDocumentByIdResponse{
			Results: mapModelsToEntities(models),
		}, nil
	})
	if err != nil {
		return types.FetchNdaDocumentByIdResponse{}, err
	}
	result := v.(types.FetchNdaDocumentByIdResponse)
	if !resolvedOpts.SkipCache {
		m.cache.Set(cacheKey, result, 0)
	}
	return result, nil

}
