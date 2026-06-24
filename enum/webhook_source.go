package enum

import (
	"encoding/json"
	"github.com/guregu/null/v6"
	"log"
)

//go:generate go run github.com/dmarkham/enumer -type=WebhookSource -json
type WebhookSource int64

const (
	WEBHOOK_SOURCE_INVALID = iota
	WEBHOOK_SOURCE_STRIPE
	WEBHOOK_SOURCE_SCHEDULER
	WEBHOOK_SOURCE_ESIGN
)

func (e WebhookSource) ToInt64() int64 {
	return int64(e)
}

func (e WebhookSource) ToNullInt() null.Int {
	return null.NewInt(int64(e), true)
}

func WebhookSourceFromString(in string) WebhookSource {
	switch in {
	case "invalid":
		return WEBHOOK_SOURCE_INVALID
	case "stripe":
		return WEBHOOK_SOURCE_STRIPE
	case "scheduler":
		return WEBHOOK_SOURCE_SCHEDULER
	case "esign":
		return WEBHOOK_SOURCE_ESIGN
	}
	return WEBHOOK_SOURCE_INVALID
}

func WebhookSourceFromPointerString(in *string) WebhookSource {
	if in == nil {
		return WEBHOOK_SOURCE_INVALID
	}
	return WebhookSourceFromString(*in)
}

func (e WebhookSource) String() string {
	switch e {
	case WEBHOOK_SOURCE_INVALID:
		return "invalid"
	case WEBHOOK_SOURCE_STRIPE:
		return "stripe"
	case WEBHOOK_SOURCE_SCHEDULER:
		return "scheduler"
	case WEBHOOK_SOURCE_ESIGN:
		return "esign"
	}

	return "invalid"
}

func (e WebhookSource) StringPtr() *string {
	val := e.String()
	return &val
}

func WebhookSourceSliceToJSON(in []WebhookSource) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		log.Printf("error marshaling WebhookSource slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToWebhookSourceSlice(in json.RawMessage) []WebhookSource {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		log.Printf("error unmarshaling WebhookSource slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []WebhookSource{}
	for _, r := range res {
		finalRes = append(finalRes, WebhookSource(r))
	}
	return finalRes
}
