package magic_link

import (
	"encoding/json"
	"log"

	"github.com/gofrs/uuid"
	"github.com/guregu/null/v6"
	"github.com/mklfarha/mfconsult/enum"

	"github.com/mklfarha/mfconsult/entity/mapper"
)

type MagicLink struct {
	ID         uuid.UUID             `json:"id"`
	ClientId   uuid.UUID             `json:"client_id"`
	Email      null.String           `json:"email"`
	Token      string                `json:"token"`
	Purpose    enum.MagicLinkPurpose `json:"purpose"`
	ExpiresAt  null.Time             `json:"expires_at"`
	ConsumedAt null.Time             `json:"consumed_at"`
	CreatedAt  null.Time             `json:"created_at"`
	CreatedIp  null.String           `json:"created_ip"`
}

func (e MagicLink) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e MagicLink) PrimaryKeyValues() []string {
	return []string{
		e.ID.String(),
	}
}

func MagicLinkFromJSON(data json.RawMessage) MagicLink {
	entity := MagicLink{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		if err2 := mapper.FlexibleUnmarshal(data, &entity); err2 != nil {
			log.Printf("flexible unmarshal error MagicLinkFromJSON: %v\n", err2)
		}
	}
	return entity
}

func MagicLinkSliceFromJSON(data json.RawMessage) []MagicLink {
	entity := []MagicLink{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		entity = []MagicLink{}
		var rawSlice []json.RawMessage
		if err2 := json.Unmarshal(data, &rawSlice); err2 == nil {
			for _, raw := range rawSlice {
				item := MagicLink{}
				if err3 := mapper.FlexibleUnmarshal(raw, &item); err3 != nil {
					log.Printf("flexible unmarshal error MagicLinkSliceFromJSON item: %v\n", err3)
				}
				entity = append(entity, item)
			}
		}
	}
	return entity
}

func (e MagicLink) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error MagicLinkToJSON: %v\n", err)
	}
	return res
}

func MagicLinkToJSON(e MagicLink) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error MagicLinkToJSON: %v\n", err)
	}
	return res
}

func MagicLinkSliceToJSON(e []MagicLink) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error MagicLinkSliceToJSON: %v\n", err)
	}
	return res
}
