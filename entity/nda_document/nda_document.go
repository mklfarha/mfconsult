package nda_document

import (
	"encoding/json"
	"log"

	"github.com/gofrs/uuid"
	"github.com/guregu/null/v6"
	"github.com/mklfarha/mfconsult/enum"

	"github.com/mklfarha/mfconsult/entity/mapper"
)

type NdaDocument struct {
	ID             uuid.UUID      `json:"id"`
	ClientId       uuid.UUID      `json:"client_id"`
	URL            null.String    `json:"url"`
	Status         enum.NdaStatus `json:"status"`
	SignedAt       null.Time      `json:"signed_at"`
	CreatedAt      null.Time      `json:"created_at"`
	EnvelopeId     null.String    `json:"envelope_id"`
	CertificateURL null.String    `json:"certificate_url"`
}

func (e NdaDocument) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e NdaDocument) PrimaryKeyValues() []string {
	return []string{
		e.ID.String(),
	}
}

func NdaDocumentFromJSON(data json.RawMessage) NdaDocument {
	entity := NdaDocument{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		if err2 := mapper.FlexibleUnmarshal(data, &entity); err2 != nil {
			log.Printf("flexible unmarshal error NdaDocumentFromJSON: %v\n", err2)
		}
	}
	return entity
}

func NdaDocumentSliceFromJSON(data json.RawMessage) []NdaDocument {
	entity := []NdaDocument{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		entity = []NdaDocument{}
		var rawSlice []json.RawMessage
		if err2 := json.Unmarshal(data, &rawSlice); err2 == nil {
			for _, raw := range rawSlice {
				item := NdaDocument{}
				if err3 := mapper.FlexibleUnmarshal(raw, &item); err3 != nil {
					log.Printf("flexible unmarshal error NdaDocumentSliceFromJSON item: %v\n", err3)
				}
				entity = append(entity, item)
			}
		}
	}
	return entity
}

func (e NdaDocument) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error NdaDocumentToJSON: %v\n", err)
	}
	return res
}

func NdaDocumentToJSON(e NdaDocument) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error NdaDocumentToJSON: %v\n", err)
	}
	return res
}

func NdaDocumentSliceToJSON(e []NdaDocument) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error NdaDocumentSliceToJSON: %v\n", err)
	}
	return res
}
