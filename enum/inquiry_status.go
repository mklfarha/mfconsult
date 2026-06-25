package enum

import (
	"encoding/json"
	"github.com/guregu/null/v6"
	"log"
)

//go:generate go run github.com/dmarkham/enumer -type=InquiryStatus -json
type InquiryStatus int64

const (
	INQUIRY_STATUS_INVALID = iota
	INQUIRY_STATUS_NEW
	INQUIRY_STATUS_REVIEWING
	INQUIRY_STATUS_CONTACTED
	INQUIRY_STATUS_QUALIFIED
	INQUIRY_STATUS_DECLINED
)

func (e InquiryStatus) ToInt64() int64 {
	return int64(e)
}

func (e InquiryStatus) ToNullInt() null.Int {
	return null.NewInt(int64(e), true)
}

func InquiryStatusFromString(in string) InquiryStatus {
	switch in {
	case "invalid":
		return INQUIRY_STATUS_INVALID
	case "new":
		return INQUIRY_STATUS_NEW
	case "reviewing":
		return INQUIRY_STATUS_REVIEWING
	case "contacted":
		return INQUIRY_STATUS_CONTACTED
	case "qualified":
		return INQUIRY_STATUS_QUALIFIED
	case "declined":
		return INQUIRY_STATUS_DECLINED
	}
	return INQUIRY_STATUS_INVALID
}

func InquiryStatusFromPointerString(in *string) InquiryStatus {
	if in == nil {
		return INQUIRY_STATUS_INVALID
	}
	return InquiryStatusFromString(*in)
}

func (e InquiryStatus) String() string {
	switch e {
	case INQUIRY_STATUS_INVALID:
		return "invalid"
	case INQUIRY_STATUS_NEW:
		return "new"
	case INQUIRY_STATUS_REVIEWING:
		return "reviewing"
	case INQUIRY_STATUS_CONTACTED:
		return "contacted"
	case INQUIRY_STATUS_QUALIFIED:
		return "qualified"
	case INQUIRY_STATUS_DECLINED:
		return "declined"
	}

	return "invalid"
}

func (e InquiryStatus) StringPtr() *string {
	val := e.String()
	return &val
}

func InquiryStatusSliceToJSON(in []InquiryStatus) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		log.Printf("error marshaling InquiryStatus slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToInquiryStatusSlice(in json.RawMessage) []InquiryStatus {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		log.Printf("error unmarshaling InquiryStatus slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []InquiryStatus{}
	for _, r := range res {
		finalRes = append(finalRes, InquiryStatus(r))
	}
	return finalRes
}
