package booking_document

import (
	entitytypes "github.com/mklfarha/mfconsult/entity/types"
)

func (e BookingDocument) FieldIdentifierToTypeMap() map[string]entitytypes.FieldType {
	return map[string]entitytypes.FieldType{
		"id":          entitytypes.StringFieldType,
		"booking_id":  entitytypes.StringFieldType,
		"kind":        entitytypes.SingleEnumFieldType,
		"url":         entitytypes.StringFieldType,
		"label":       entitytypes.StringFieldType,
		"purge_after": entitytypes.TimestampFieldType,
		"created_at":  entitytypes.TimestampFieldType,
	}
}

func (e BookingDocument) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "id")
	res = append(res, "booking_id")
	res = append(res, "kind")
	res = append(res, "url")
	res = append(res, "label")
	res = append(res, "purge_after")
	res = append(res, "created_at")

	return res
}

func (e BookingDocument) DependantFieldIdentifierToTypeMap() map[string]map[string]entitytypes.FieldType {
	res := make(map[string]map[string]entitytypes.FieldType)

	return res
}

func (e BookingDocument) EntityIdentifier() string {
	return "booking_document"
}

func (e BookingDocument) PrimaryKeyIdentifiers() []string {
	return []string{
		"id",
	}
}

func (e BookingDocument) ArrayFieldIdentifierToType() map[string]entitytypes.FieldType {
	res := make(map[string]entitytypes.FieldType)

	return res
}
