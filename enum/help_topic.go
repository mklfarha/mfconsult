package enum

import (
	"encoding/json"
	"github.com/guregu/null/v6"
	"log"
)

//go:generate go run github.com/dmarkham/enumer -type=HelpTopic -json
type HelpTopic int64

const (
	HELP_TOPIC_INVALID = iota
	HELP_TOPIC_GREENFIELD_ARCHITECTURE
	HELP_TOPIC_IMPROVING_EXISTING_SYSTEM
	HELP_TOPIC_AUDITING_AI_GENERATED_CODE
	HELP_TOPIC_SCALING_PERFORMANCE
	HELP_TOPIC_TECHNICAL_LEADERSHIP
	HELP_TOPIC_OTHER
)

func (e HelpTopic) ToInt64() int64 {
	return int64(e)
}

func (e HelpTopic) ToNullInt() null.Int {
	return null.NewInt(int64(e), true)
}

func HelpTopicFromString(in string) HelpTopic {
	switch in {
	case "invalid":
		return HELP_TOPIC_INVALID
	case "greenfield_architecture":
		return HELP_TOPIC_GREENFIELD_ARCHITECTURE
	case "improving_existing_system":
		return HELP_TOPIC_IMPROVING_EXISTING_SYSTEM
	case "auditing_ai_generated_code":
		return HELP_TOPIC_AUDITING_AI_GENERATED_CODE
	case "scaling_performance":
		return HELP_TOPIC_SCALING_PERFORMANCE
	case "technical_leadership":
		return HELP_TOPIC_TECHNICAL_LEADERSHIP
	case "other":
		return HELP_TOPIC_OTHER
	}
	return HELP_TOPIC_INVALID
}

func HelpTopicFromPointerString(in *string) HelpTopic {
	if in == nil {
		return HELP_TOPIC_INVALID
	}
	return HelpTopicFromString(*in)
}

func (e HelpTopic) String() string {
	switch e {
	case HELP_TOPIC_INVALID:
		return "invalid"
	case HELP_TOPIC_GREENFIELD_ARCHITECTURE:
		return "greenfield_architecture"
	case HELP_TOPIC_IMPROVING_EXISTING_SYSTEM:
		return "improving_existing_system"
	case HELP_TOPIC_AUDITING_AI_GENERATED_CODE:
		return "auditing_ai_generated_code"
	case HELP_TOPIC_SCALING_PERFORMANCE:
		return "scaling_performance"
	case HELP_TOPIC_TECHNICAL_LEADERSHIP:
		return "technical_leadership"
	case HELP_TOPIC_OTHER:
		return "other"
	}

	return "invalid"
}

func (e HelpTopic) StringPtr() *string {
	val := e.String()
	return &val
}

func HelpTopicSliceToJSON(in []HelpTopic) json.RawMessage {
	res := make([]int64, len(in))
	for i, e := range in {
		res[i] = int64(e)
	}
	jr, err := json.Marshal(res)
	if err != nil {
		log.Printf("error marshaling HelpTopic slice to json: %v", err)
		return json.RawMessage{}
	}
	return jr
}

func JSONToHelpTopicSlice(in json.RawMessage) []HelpTopic {
	res := []int64{}
	err := json.Unmarshal(in, &res)
	if err != nil {
		log.Printf("error unmarshaling HelpTopic slice to int slice: %v", err)
		return nil
	}
	if len(res) == 0 {
		return nil
	}
	finalRes := []HelpTopic{}
	for _, r := range res {
		finalRes = append(finalRes, HelpTopic(r))
	}
	return finalRes
}
