package booking_recap

import (
	entitytypes "github.com/mklfarha/mfconsult/entity/types"
)

func (e BookingRecap) FieldIdentifierToTypeMap() map[string]entitytypes.FieldType {
	return map[string]entitytypes.FieldType{
		"id":           entitytypes.StringFieldType,
		"booking_id":   entitytypes.StringFieldType,
		"body":         entitytypes.StringFieldType,
		"published_at": entitytypes.TimestampFieldType,
		"created_at":   entitytypes.TimestampFieldType,
	}
}

func (e BookingRecap) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "id")
	res = append(res, "booking_id")
	res = append(res, "body")
	res = append(res, "published_at")
	res = append(res, "created_at")

	return res
}

func (e BookingRecap) DependantFieldIdentifierToTypeMap() map[string]map[string]entitytypes.FieldType {
	res := make(map[string]map[string]entitytypes.FieldType)

	return res
}

func (e BookingRecap) EntityIdentifier() string {
	return "booking_recap"
}

func (e BookingRecap) PrimaryKeyIdentifiers() []string {
	return []string{
		"id",
	}
}

func (e BookingRecap) ArrayFieldIdentifierToType() map[string]entitytypes.FieldType {
	res := make(map[string]entitytypes.FieldType)

	return res
}
