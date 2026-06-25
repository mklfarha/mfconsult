package engagement_inquiry

import (
	"encoding/json"
	"log"

	"github.com/gofrs/uuid"
	"github.com/guregu/null/v6"
	"github.com/mklfarha/mfconsult/enum"

	"github.com/mklfarha/mfconsult/entity/mapper"
)

type EngagementInquiry struct {
	ID                 uuid.UUID          `json:"id"`
	ClientId           *uuid.UUID         `json:"client_id"`
	Name               string             `json:"name"`
	Email              string             `json:"email"`
	Phone              null.String        `json:"phone"`
	Company            null.String        `json:"company"`
	ProjectSummary     string             `json:"project_summary"`
	WhyMoreThanSession null.String        `json:"why_more_than_session"`
	ScopeDetails       null.String        `json:"scope_details"`
	BudgetRange        null.String        `json:"budget_range"`
	Timeline           null.String        `json:"timeline"`
	Status             enum.InquiryStatus `json:"status"`
	ReviewNotes        null.String        `json:"review_notes"`
	CreatedAt          null.Time          `json:"created_at"`
	UpdatedAt          null.Time          `json:"updated_at"`
}

func (e EngagementInquiry) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e EngagementInquiry) PrimaryKeyValues() []string {
	return []string{
		e.ID.String(),
	}
}

func EngagementInquiryFromJSON(data json.RawMessage) EngagementInquiry {
	entity := EngagementInquiry{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		if err2 := mapper.FlexibleUnmarshal(data, &entity); err2 != nil {
			log.Printf("flexible unmarshal error EngagementInquiryFromJSON: %v\n", err2)
		}
	}
	return entity
}

func EngagementInquirySliceFromJSON(data json.RawMessage) []EngagementInquiry {
	entity := []EngagementInquiry{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		entity = []EngagementInquiry{}
		var rawSlice []json.RawMessage
		if err2 := json.Unmarshal(data, &rawSlice); err2 == nil {
			for _, raw := range rawSlice {
				item := EngagementInquiry{}
				if err3 := mapper.FlexibleUnmarshal(raw, &item); err3 != nil {
					log.Printf("flexible unmarshal error EngagementInquirySliceFromJSON item: %v\n", err3)
				}
				entity = append(entity, item)
			}
		}
	}
	return entity
}

func (e EngagementInquiry) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error EngagementInquiryToJSON: %v\n", err)
	}
	return res
}

func EngagementInquiryToJSON(e EngagementInquiry) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error EngagementInquiryToJSON: %v\n", err)
	}
	return res
}

func EngagementInquirySliceToJSON(e []EngagementInquiry) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error EngagementInquirySliceToJSON: %v\n", err)
	}
	return res
}
