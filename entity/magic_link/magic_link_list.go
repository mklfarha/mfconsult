package magic_link

import (
	entitytypes "github.com/mklfarha/mfconsult/entity/types"
)

func (e MagicLink) FieldIdentifierToTypeMap() map[string]entitytypes.FieldType {
	return map[string]entitytypes.FieldType{
		"id":          entitytypes.StringFieldType,
		"client_id":   entitytypes.StringFieldType,
		"email":       entitytypes.StringFieldType,
		"token":       entitytypes.StringFieldType,
		"purpose":     entitytypes.SingleEnumFieldType,
		"expires_at":  entitytypes.TimestampFieldType,
		"consumed_at": entitytypes.TimestampFieldType,
		"created_at":  entitytypes.TimestampFieldType,
		"created_ip":  entitytypes.StringFieldType,
	}
}

func (e MagicLink) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "id")
	res = append(res, "client_id")
	res = append(res, "email")
	res = append(res, "token")
	res = append(res, "purpose")
	res = append(res, "expires_at")
	res = append(res, "consumed_at")
	res = append(res, "created_at")
	res = append(res, "created_ip")

	return res
}

func (e MagicLink) DependantFieldIdentifierToTypeMap() map[string]map[string]entitytypes.FieldType {
	res := make(map[string]map[string]entitytypes.FieldType)

	return res
}

func (e MagicLink) EntityIdentifier() string {
	return "magic_link"
}

func (e MagicLink) PrimaryKeyIdentifiers() []string {
	return []string{
		"id",
	}
}

func (e MagicLink) ArrayFieldIdentifierToType() map[string]entitytypes.FieldType {
	res := make(map[string]entitytypes.FieldType)

	return res
}
