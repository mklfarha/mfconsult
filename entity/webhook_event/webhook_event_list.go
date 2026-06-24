package webhook_event

import (
	entitytypes "github.com/mklfarha/mfconsult/entity/types"
)

func (e WebhookEvent) FieldIdentifierToTypeMap() map[string]entitytypes.FieldType {
	return map[string]entitytypes.FieldType{
		"id":           entitytypes.StringFieldType,
		"source":       entitytypes.SingleEnumFieldType,
		"event_type":   entitytypes.StringFieldType,
		"payload":      entitytypes.StringFieldType,
		"processed_at": entitytypes.TimestampFieldType,
		"created_at":   entitytypes.TimestampFieldType,
	}
}

func (e WebhookEvent) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "id")
	res = append(res, "source")
	res = append(res, "event_type")
	res = append(res, "payload")
	res = append(res, "processed_at")
	res = append(res, "created_at")

	return res
}

func (e WebhookEvent) DependantFieldIdentifierToTypeMap() map[string]map[string]entitytypes.FieldType {
	res := make(map[string]map[string]entitytypes.FieldType)

	return res
}

func (e WebhookEvent) EntityIdentifier() string {
	return "webhook_event"
}

func (e WebhookEvent) PrimaryKeyIdentifiers() []string {
	return []string{
		"id",
	}
}

func (e WebhookEvent) ArrayFieldIdentifierToType() map[string]entitytypes.FieldType {
	res := make(map[string]entitytypes.FieldType)

	return res
}
