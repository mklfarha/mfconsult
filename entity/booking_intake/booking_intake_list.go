package booking_intake

import (
	entitytypes "github.com/mklfarha/mfconsult/entity/types"
)

func (e BookingIntake) FieldIdentifierToTypeMap() map[string]entitytypes.FieldType {
	return map[string]entitytypes.FieldType{
		"reason":        entitytypes.StringFieldType,
		"help_topic":    entitytypes.SingleEnumFieldType,
		"help_details":  entitytypes.StringFieldType,
		"stack_details": entitytypes.StringFieldType,
		"prep_notes":    entitytypes.StringFieldType,
	}
}

func (e BookingIntake) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "reason")
	res = append(res, "help_topic")
	res = append(res, "help_details")
	res = append(res, "stack_details")
	res = append(res, "prep_notes")

	return res
}

func (e BookingIntake) DependantFieldIdentifierToTypeMap() map[string]map[string]entitytypes.FieldType {
	res := make(map[string]map[string]entitytypes.FieldType)

	return res
}

func (e BookingIntake) EntityIdentifier() string {
	return "booking_intake"
}

func (e BookingIntake) PrimaryKeyIdentifiers() []string {
	return []string{}
}

func (e BookingIntake) ArrayFieldIdentifierToType() map[string]entitytypes.FieldType {
	res := make(map[string]entitytypes.FieldType)

	return res
}
