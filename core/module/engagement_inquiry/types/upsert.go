package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/engagement_inquiry"

	"github.com/gofrs/uuid"
)

type UpsertRequest struct {
	EngagementInquiry main_entity.EngagementInquiry
}

type UpsertResponse struct {
	ID uuid.UUID
}
