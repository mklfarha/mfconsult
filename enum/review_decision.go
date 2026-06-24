package enum

import (
	"encoding/json"
	"github.com/guregu/null/v6"
	"log"
)

//go:generate go run github.com/dmarkham/enumer -type=ReviewDecision -json
type ReviewDecision int64

const (
	REVIEW_DECISION_INVALID = iota
	REVIEW_DECISION_PENDING
	REVIEW_DECISION_APPROVED
	REVIEW_DECISION_DECLINED
)

func (e ReviewDecision) ToInt64() int64 {
	return int64(e)
}

func (e ReviewDecision) ToNullInt() null.Int {
	return null.NewInt(int64(e), true)
}

func ReviewDecisionFromString(in string) ReviewDecision {
	switch in {
	case "invalid":
		return REVIEW_DECISION_INVALID
	case "pending":
		return REVIEW_DECISION_PENDING
	case "approved":
		return REVIEW_DECISION_APPROVED
	case "declined":
		return REVIEW_DECISION_DECLINED
	}
	return REVIEW_DECISION_INVALID
}

func ReviewDecisionFromPointerString(in *string) ReviewDecision {
	if in == nil {
		return REVIEW_DECISION_INVALID
	}
	return ReviewDecisionFromString(*in)
}

func (e ReviewDecision) String() string {
	switch e {
	case REVIEW_DECISION_INVALID:
		return "invalid"
	case REVIEW_DECISION_PENDING:
		return "pending"
	case REVIEW_DECISION_APPROVED:
		return "approved"
	case REVIEW_DECISION_DECLINED:
		return "declined"
	}

	return "invalid"
}

func (e ReviewDecision) StringPtr() *string {
	val := e.String()
	return &val
}

func ReviewDecisionSliceToJSON(in []ReviewDecision) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		log.Printf("error marshaling ReviewDecision slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToReviewDecisionSlice(in json.RawMessage) []ReviewDecision {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		log.Printf("error unmarshaling ReviewDecision slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []ReviewDecision{}
	for _, r := range res {
		finalRes = append(finalRes, ReviewDecision(r))
	}
	return finalRes
}
