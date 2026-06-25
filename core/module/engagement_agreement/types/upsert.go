package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/engagement_agreement"

	"github.com/gofrs/uuid"
)

type UpsertRequest struct {
	EngagementAgreement main_entity.EngagementAgreement
}

type UpsertResponse struct {
	ID uuid.UUID
}
