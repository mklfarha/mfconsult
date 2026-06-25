package engagement_agreement

import (
	entitytypes "github.com/mklfarha/mfconsult/entity/types"
)

func (e EngagementAgreement) FieldIdentifierToTypeMap() map[string]entitytypes.FieldType {
	return map[string]entitytypes.FieldType{
		"id":                    entitytypes.StringFieldType,
		"client_id":             entitytypes.StringFieldType,
		"nda_url":               entitytypes.StringFieldType,
		"status":                entitytypes.SingleEnumFieldType,
		"signed_at":             entitytypes.TimestampFieldType,
		"created_at":            entitytypes.TimestampFieldType,
		"envelope_id":           entitytypes.StringFieldType,
		"certificate_url":       entitytypes.StringFieldType,
		"contract_url":          entitytypes.StringFieldType,
		"engagement_inquiry_id": entitytypes.StringFieldType,
		"updated_at":            entitytypes.TimestampFieldType,
	}
}

func (e EngagementAgreement) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "id")
	res = append(res, "client_id")
	res = append(res, "nda_url")
	res = append(res, "status")
	res = append(res, "signed_at")
	res = append(res, "created_at")
	res = append(res, "envelope_id")
	res = append(res, "certificate_url")
	res = append(res, "contract_url")
	res = append(res, "engagement_inquiry_id")
	res = append(res, "updated_at")

	return res
}

func (e EngagementAgreement) DependantFieldIdentifierToTypeMap() map[string]map[string]entitytypes.FieldType {
	res := make(map[string]map[string]entitytypes.FieldType)

	return res
}

func (e EngagementAgreement) EntityIdentifier() string {
	return "engagement_agreement"
}

func (e EngagementAgreement) PrimaryKeyIdentifiers() []string {
	return []string{
		"id",
	}
}

func (e EngagementAgreement) ArrayFieldIdentifierToType() map[string]entitytypes.FieldType {
	res := make(map[string]entitytypes.FieldType)

	return res
}
