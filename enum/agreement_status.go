package enum

import (
	"encoding/json"
	"github.com/guregu/null/v6"
	"log"
)

//go:generate go run github.com/dmarkham/enumer -type=AgreementStatus -json
type AgreementStatus int64

const (
	AGREEMENT_STATUS_INVALID = iota
	AGREEMENT_STATUS_REQUESTED
	AGREEMENT_STATUS_SENT
	AGREEMENT_STATUS_SIGNED
)

func (e AgreementStatus) ToInt64() int64 {
	return int64(e)
}

func (e AgreementStatus) ToNullInt() null.Int {
	return null.NewInt(int64(e), true)
}

func AgreementStatusFromString(in string) AgreementStatus {
	switch in {
	case "invalid":
		return AGREEMENT_STATUS_INVALID
	case "requested":
		return AGREEMENT_STATUS_REQUESTED
	case "sent":
		return AGREEMENT_STATUS_SENT
	case "signed":
		return AGREEMENT_STATUS_SIGNED
	}
	return AGREEMENT_STATUS_INVALID
}

func AgreementStatusFromPointerString(in *string) AgreementStatus {
	if in == nil {
		return AGREEMENT_STATUS_INVALID
	}
	return AgreementStatusFromString(*in)
}

func (e AgreementStatus) String() string {
	switch e {
	case AGREEMENT_STATUS_INVALID:
		return "invalid"
	case AGREEMENT_STATUS_REQUESTED:
		return "requested"
	case AGREEMENT_STATUS_SENT:
		return "sent"
	case AGREEMENT_STATUS_SIGNED:
		return "signed"
	}

	return "invalid"
}

func (e AgreementStatus) StringPtr() *string {
	val := e.String()
	return &val
}

func AgreementStatusSliceToJSON(in []AgreementStatus) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		log.Printf("error marshaling AgreementStatus slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToAgreementStatusSlice(in json.RawMessage) []AgreementStatus {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		log.Printf("error unmarshaling AgreementStatus slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []AgreementStatus{}
	for _, r := range res {
		finalRes = append(finalRes, AgreementStatus(r))
	}
	return finalRes
}
