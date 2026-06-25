package enum

import (
	"encoding/json"
	"github.com/guregu/null/v6"
	"log"
)

//go:generate go run github.com/dmarkham/enumer -type=MagicLinkPurpose -json
type MagicLinkPurpose int64

const (
	MAGIC_LINK_PURPOSE_INVALID = iota
	MAGIC_LINK_PURPOSE_BOOKING_ACCESS
	MAGIC_LINK_PURPOSE_PORTAL_ACCESS
)

func (e MagicLinkPurpose) ToInt64() int64 {
	return int64(e)
}

func (e MagicLinkPurpose) ToNullInt() null.Int {
	return null.NewInt(int64(e), true)
}

func MagicLinkPurposeFromString(in string) MagicLinkPurpose {
	switch in {
	case "invalid":
		return MAGIC_LINK_PURPOSE_INVALID
	case "booking_access":
		return MAGIC_LINK_PURPOSE_BOOKING_ACCESS
	case "portal_access":
		return MAGIC_LINK_PURPOSE_PORTAL_ACCESS
	}
	return MAGIC_LINK_PURPOSE_INVALID
}

func MagicLinkPurposeFromPointerString(in *string) MagicLinkPurpose {
	if in == nil {
		return MAGIC_LINK_PURPOSE_INVALID
	}
	return MagicLinkPurposeFromString(*in)
}

func (e MagicLinkPurpose) String() string {
	switch e {
	case MAGIC_LINK_PURPOSE_INVALID:
		return "invalid"
	case MAGIC_LINK_PURPOSE_BOOKING_ACCESS:
		return "booking_access"
	case MAGIC_LINK_PURPOSE_PORTAL_ACCESS:
		return "portal_access"
	}

	return "invalid"
}

func (e MagicLinkPurpose) StringPtr() *string {
	val := e.String()
	return &val
}

func MagicLinkPurposeSliceToJSON(in []MagicLinkPurpose) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		log.Printf("error marshaling MagicLinkPurpose slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToMagicLinkPurposeSlice(in json.RawMessage) []MagicLinkPurpose {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		log.Printf("error unmarshaling MagicLinkPurpose slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []MagicLinkPurpose{}
	for _, r := range res {
		finalRes = append(finalRes, MagicLinkPurpose(r))
	}
	return finalRes
}
