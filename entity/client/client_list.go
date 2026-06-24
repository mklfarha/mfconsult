package client

import (
	entitytypes "github.com/mklfarha/mfconsult/entity/types"
)

func (e Client) FieldIdentifierToTypeMap() map[string]entitytypes.FieldType {
	return map[string]entitytypes.FieldType{
		"id":         entitytypes.StringFieldType,
		"name":       entitytypes.StringFieldType,
		"email":      entitytypes.StringFieldType,
		"timezone":   entitytypes.StringFieldType,
		"notes":      entitytypes.StringFieldType,
		"created_at": entitytypes.TimestampFieldType,
		"updated_at": entitytypes.TimestampFieldType,
	}
}

func (e Client) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "id")
	res = append(res, "name")
	res = append(res, "email")
	res = append(res, "timezone")
	res = append(res, "notes")
	res = append(res, "created_at")
	res = append(res, "updated_at")

	return res
}

func (e Client) DependantFieldIdentifierToTypeMap() map[string]map[string]entitytypes.FieldType {
	res := make(map[string]map[string]entitytypes.FieldType)

	return res
}

func (e Client) EntityIdentifier() string {
	return "client"
}

func (e Client) PrimaryKeyIdentifiers() []string {
	return []string{
		"id",
	}
}

func (e Client) ArrayFieldIdentifierToType() map[string]entitytypes.FieldType {
	res := make(map[string]entitytypes.FieldType)

	return res
}
