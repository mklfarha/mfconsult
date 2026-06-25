package engagement_inquiry

import (
	entitytypes "github.com/mklfarha/mfconsult/entity/types"
)

func (e EngagementInquiry) FieldIdentifierToTypeMap() map[string]entitytypes.FieldType {
	return map[string]entitytypes.FieldType{
		"id":                    entitytypes.StringFieldType,
		"client_id":             entitytypes.StringFieldType,
		"name":                  entitytypes.StringFieldType,
		"email":                 entitytypes.StringFieldType,
		"phone":                 entitytypes.StringFieldType,
		"company":               entitytypes.StringFieldType,
		"project_summary":       entitytypes.StringFieldType,
		"why_more_than_session": entitytypes.StringFieldType,
		"scope_details":         entitytypes.StringFieldType,
		"budget_range":          entitytypes.StringFieldType,
		"timeline":              entitytypes.StringFieldType,
		"status":                entitytypes.SingleEnumFieldType,
		"review_notes":          entitytypes.StringFieldType,
		"created_at":            entitytypes.TimestampFieldType,
		"updated_at":            entitytypes.TimestampFieldType,
	}
}

func (e EngagementInquiry) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "id")
	res = append(res, "client_id")
	res = append(res, "name")
	res = append(res, "email")
	res = append(res, "phone")
	res = append(res, "company")
	res = append(res, "project_summary")
	res = append(res, "why_more_than_session")
	res = append(res, "scope_details")
	res = append(res, "budget_range")
	res = append(res, "timeline")
	res = append(res, "status")
	res = append(res, "review_notes")
	res = append(res, "created_at")
	res = append(res, "updated_at")

	return res
}

func (e EngagementInquiry) DependantFieldIdentifierToTypeMap() map[string]map[string]entitytypes.FieldType {
	res := make(map[string]map[string]entitytypes.FieldType)

	return res
}

func (e EngagementInquiry) EntityIdentifier() string {
	return "engagement_inquiry"
}

func (e EngagementInquiry) PrimaryKeyIdentifiers() []string {
	return []string{
		"id",
	}
}

func (e EngagementInquiry) ArrayFieldIdentifierToType() map[string]entitytypes.FieldType {
	res := make(map[string]entitytypes.FieldType)

	return res
}
