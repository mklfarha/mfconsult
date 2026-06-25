package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/engagement_agreement"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchEngagementAgreementByIdRequest struct {
	ID uuid.UUID
}

func (r FetchEngagementAgreementByIdRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("id", r.ID.String())

	return nil
}

type FetchEngagementAgreementByIdResponse struct {
	Results []main_entity.EngagementAgreement
}
