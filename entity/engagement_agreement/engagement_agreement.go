package engagement_agreement

import (
	"encoding/json"
	"log"

	"github.com/gofrs/uuid"
	"github.com/guregu/null/v6"
	"github.com/mklfarha/mfconsult/enum"

	"github.com/mklfarha/mfconsult/entity/mapper"
)

type EngagementAgreement struct {
	ID                  uuid.UUID            `json:"id"`
	ClientId            uuid.UUID            `json:"client_id"`
	NdaURL              null.String          `json:"nda_url"`
	Status              enum.AgreementStatus `json:"status"`
	SignedAt            null.Time            `json:"signed_at"`
	CreatedAt           null.Time            `json:"created_at"`
	EnvelopeId          null.String          `json:"envelope_id"`
	CertificateURL      null.String          `json:"certificate_url"`
	ContractURL         null.String          `json:"contract_url"`
	EngagementInquiryId uuid.UUID            `json:"engagement_inquiry_id"`
	UpdatedAt           null.Time            `json:"updated_at"`
}

func (e EngagementAgreement) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e EngagementAgreement) PrimaryKeyValues() []string {
	return []string{
		e.ID.String(),
	}
}

func EngagementAgreementFromJSON(data json.RawMessage) EngagementAgreement {
	entity := EngagementAgreement{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		if err2 := mapper.FlexibleUnmarshal(data, &entity); err2 != nil {
			log.Printf("flexible unmarshal error EngagementAgreementFromJSON: %v\n", err2)
		}
	}
	return entity
}

func EngagementAgreementSliceFromJSON(data json.RawMessage) []EngagementAgreement {
	entity := []EngagementAgreement{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		entity = []EngagementAgreement{}
		var rawSlice []json.RawMessage
		if err2 := json.Unmarshal(data, &rawSlice); err2 == nil {
			for _, raw := range rawSlice {
				item := EngagementAgreement{}
				if err3 := mapper.FlexibleUnmarshal(raw, &item); err3 != nil {
					log.Printf("flexible unmarshal error EngagementAgreementSliceFromJSON item: %v\n", err3)
				}
				entity = append(entity, item)
			}
		}
	}
	return entity
}

func (e EngagementAgreement) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error EngagementAgreementToJSON: %v\n", err)
	}
	return res
}

func EngagementAgreementToJSON(e EngagementAgreement) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error EngagementAgreementToJSON: %v\n", err)
	}
	return res
}

func EngagementAgreementSliceToJSON(e []EngagementAgreement) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error EngagementAgreementSliceToJSON: %v\n", err)
	}
	return res
}
