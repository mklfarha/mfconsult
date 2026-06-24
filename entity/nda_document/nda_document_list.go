package nda_document

import (
	entitytypes "github.com/mklfarha/mfconsult/entity/types"
)

func (e NdaDocument) FieldIdentifierToTypeMap() map[string]entitytypes.FieldType {
	return map[string]entitytypes.FieldType{
		"id":              entitytypes.StringFieldType,
		"client_id":       entitytypes.StringFieldType,
		"url":             entitytypes.StringFieldType,
		"status":          entitytypes.SingleEnumFieldType,
		"signed_at":       entitytypes.TimestampFieldType,
		"created_at":      entitytypes.TimestampFieldType,
		"envelope_id":     entitytypes.StringFieldType,
		"certificate_url": entitytypes.StringFieldType,
	}
}

func (e NdaDocument) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "id")
	res = append(res, "client_id")
	res = append(res, "url")
	res = append(res, "status")
	res = append(res, "signed_at")
	res = append(res, "created_at")
	res = append(res, "envelope_id")
	res = append(res, "certificate_url")

	return res
}

func (e NdaDocument) DependantFieldIdentifierToTypeMap() map[string]map[string]entitytypes.FieldType {
	res := make(map[string]map[string]entitytypes.FieldType)

	return res
}

func (e NdaDocument) EntityIdentifier() string {
	return "nda_document"
}

func (e NdaDocument) PrimaryKeyIdentifiers() []string {
	return []string{
		"id",
	}
}

func (e NdaDocument) ArrayFieldIdentifierToType() map[string]entitytypes.FieldType {
	res := make(map[string]entitytypes.FieldType)

	return res
}
