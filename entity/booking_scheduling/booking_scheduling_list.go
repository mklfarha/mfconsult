package booking_scheduling

import (
	entitytypes "github.com/mklfarha/mfconsult/entity/types"
)

func (e BookingScheduling) FieldIdentifierToTypeMap() map[string]entitytypes.FieldType {
	return map[string]entitytypes.FieldType{
		"slot_start":           entitytypes.TimestampFieldType,
		"slot_end":             entitytypes.TimestampFieldType,
		"scheduler_booking_id": entitytypes.StringFieldType,
		"video_url":            entitytypes.StringFieldType,
	}
}

func (e BookingScheduling) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "slot_start")
	res = append(res, "slot_end")
	res = append(res, "scheduler_booking_id")
	res = append(res, "video_url")

	return res
}

func (e BookingScheduling) DependantFieldIdentifierToTypeMap() map[string]map[string]entitytypes.FieldType {
	res := make(map[string]map[string]entitytypes.FieldType)

	return res
}

func (e BookingScheduling) EntityIdentifier() string {
	return "booking_scheduling"
}

func (e BookingScheduling) PrimaryKeyIdentifiers() []string {
	return []string{}
}

func (e BookingScheduling) ArrayFieldIdentifierToType() map[string]entitytypes.FieldType {
	res := make(map[string]entitytypes.FieldType)

	return res
}
