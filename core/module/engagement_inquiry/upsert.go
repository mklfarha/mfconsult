package engagement_inquiry

import (
	"context"

	"github.com/gofrs/uuid"

	"github.com/mklfarha/mfconsult/core/module/engagement_inquiry/types"
)

func (m *module) Upsert(
	ctx context.Context,
	req types.UpsertRequest,
	opts ...Option,
) (types.UpsertResponse, error) {
	// check if primary keys are set to determine if this is an insert or update
	isInsert := true

	if req.EngagementInquiry.ID != uuid.Nil {
		isInsert = false
	}

	if isInsert {
		return m.Insert(ctx, req, opts...)
	}

	return m.Update(ctx, req, opts...)
}
