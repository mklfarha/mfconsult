package engagement_inquiry

import (
	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"
	main_entity "github.com/mklfarha/mfconsult/entity/engagement_inquiry"

	"github.com/mklfarha/mfconsult/entity/mapper"

	"github.com/guregu/null/v6"
	"github.com/mklfarha/mfconsult/enum"
)

func mapModelsToEntities(models []mfconsultdb.EngagementInquiry) []main_entity.EngagementInquiry {
	result := []main_entity.EngagementInquiry{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(m mfconsultdb.EngagementInquiry) main_entity.EngagementInquiry {
	return main_entity.EngagementInquiry{
		ID:                 mapper.StringToUUID(m.ID),
		ClientId:           mapper.StringToUUIDPtr(m.ClientId),
		Name:               m.Name,
		Email:              m.Email,
		Phone:              null.NewString(m.Phone.String, m.Phone.Valid),
		Company:            null.NewString(m.Company.String, m.Company.Valid),
		ProjectSummary:     m.ProjectSummary,
		WhyMoreThanSession: null.NewString(m.WhyMoreThanSession.String, m.WhyMoreThanSession.Valid),
		ScopeDetails:       null.NewString(m.ScopeDetails.String, m.ScopeDetails.Valid),
		BudgetRange:        null.NewString(m.BudgetRange.String, m.BudgetRange.Valid),
		Timeline:           null.NewString(m.Timeline.String, m.Timeline.Valid),
		Status:             enum.InquiryStatus(m.Status.Int64),
		ReviewNotes:        null.NewString(m.ReviewNotes.String, m.ReviewNotes.Valid),
		CreatedAt:          null.NewTime(m.CreatedAt.Time, m.CreatedAt.Valid),
		UpdatedAt:          null.NewTime(m.UpdatedAt.Time, m.UpdatedAt.Valid),
	}
}
