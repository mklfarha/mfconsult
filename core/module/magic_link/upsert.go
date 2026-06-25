package magic_link

import (
	"context"

	"github.com/gofrs/uuid"

	"github.com/mklfarha/mfconsult/core/module/magic_link/types"
)

func (m *module) Upsert(
	ctx context.Context,
	req types.UpsertRequest,
	opts ...Option,
) (types.UpsertResponse, error) {
	// check if primary keys are set to determine if this is an insert or update
	isInsert := true

	if req.MagicLink.ID != uuid.Nil {
		isInsert = false
	}

	if isInsert {
		return m.Insert(ctx, req, opts...)
	}

	return m.Update(ctx, req, opts...)
}
