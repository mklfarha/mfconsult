package types

import (
	main_entity "github.com/mklfarha/mfconsult/entity/engagement_inquiry"

	"github.com/gofrs/uuid"
	"go.uber.org/zap/zapcore"
)

type FetchEngagementInquiryByIdRequest struct {
	ID uuid.UUID
}

func (r FetchEngagementInquiryByIdRequest) MarshalLogObject(e zapcore.ObjectEncoder) error {

	e.AddString("id", r.ID.String())

	return nil
}

type FetchEngagementInquiryByIdResponse struct {
	Results []main_entity.EngagementInquiry
}
