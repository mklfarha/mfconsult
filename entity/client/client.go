package client

import (
	"encoding/json"
	"log"

	"github.com/gofrs/uuid"
	"github.com/guregu/null/v6"

	"github.com/mklfarha/mfconsult/entity/mapper"
)

type Client struct {
	ID        uuid.UUID   `json:"id"`
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Timezone  null.String `json:"timezone"`
	Notes     null.String `json:"notes"`
	CreatedAt null.Time   `json:"created_at"`
	UpdatedAt null.Time   `json:"updated_at"`
}

func (e Client) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

func (e Client) PrimaryKeyValues() []string {
	return []string{
		e.ID.String(),
	}
}

func ClientFromJSON(data json.RawMessage) Client {
	entity := Client{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		if err2 := mapper.FlexibleUnmarshal(data, &entity); err2 != nil {
			log.Printf("flexible unmarshal error ClientFromJSON: %v\n", err2)
		}
	}
	return entity
}

func ClientSliceFromJSON(data json.RawMessage) []Client {
	entity := []Client{}
	if data == nil {
		return entity
	}
	if len(data) == 0 {
		return entity
	}

	if err := json.Unmarshal(data, &entity); err != nil {
		entity = []Client{}
		var rawSlice []json.RawMessage
		if err2 := json.Unmarshal(data, &rawSlice); err2 == nil {
			for _, raw := range rawSlice {
				item := Client{}
				if err3 := mapper.FlexibleUnmarshal(raw, &item); err3 != nil {
					log.Printf("flexible unmarshal error ClientSliceFromJSON item: %v\n", err3)
				}
				entity = append(entity, item)
			}
		}
	}
	return entity
}

func (e Client) ToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error ClientToJSON: %v\n", err)
	}
	return res
}

func ClientToJSON(e Client) json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error ClientToJSON: %v\n", err)
	}
	return res
}

func ClientSliceToJSON(e []Client) json.RawMessage {
	if e == nil {
		return json.RawMessage{}
	}
	res, err := json.Marshal(e)
	if err != nil {
		log.Printf("marshal error ClientSliceToJSON: %v\n", err)
	}
	return res
}
