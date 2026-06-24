package enum

import (
	"encoding/json"
	"github.com/guregu/null/v6"
	"log"
)

//go:generate go run github.com/dmarkham/enumer -type=NdaStatus -json
type NdaStatus int64

const (
	NDA_STATUS_INVALID = iota
	NDA_STATUS_REQUESTED
	NDA_STATUS_SENT
	NDA_STATUS_SIGNED
)

func (e NdaStatus) ToInt64() int64 {
	return int64(e)
}

func (e NdaStatus) ToNullInt() null.Int {
	return null.NewInt(int64(e), true)
}

func NdaStatusFromString(in string) NdaStatus {
	switch in {
	case "invalid":
		return NDA_STATUS_INVALID
	case "requested":
		return NDA_STATUS_REQUESTED
	case "sent":
		return NDA_STATUS_SENT
	case "signed":
		return NDA_STATUS_SIGNED
	}
	return NDA_STATUS_INVALID
}

func NdaStatusFromPointerString(in *string) NdaStatus {
	if in == nil {
		return NDA_STATUS_INVALID
	}
	return NdaStatusFromString(*in)
}

func (e NdaStatus) String() string {
	switch e {
	case NDA_STATUS_INVALID:
		return "invalid"
	case NDA_STATUS_REQUESTED:
		return "requested"
	case NDA_STATUS_SENT:
		return "sent"
	case NDA_STATUS_SIGNED:
		return "signed"
	}

	return "invalid"
}

func (e NdaStatus) StringPtr() *string {
	val := e.String()
	return &val
}

func NdaStatusSliceToJSON(in []NdaStatus) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		log.Printf("error marshaling NdaStatus slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToNdaStatusSlice(in json.RawMessage) []NdaStatus {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		log.Printf("error unmarshaling NdaStatus slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []NdaStatus{}
	for _, r := range res {
		finalRes = append(finalRes, NdaStatus(r))
	}
	return finalRes
}
