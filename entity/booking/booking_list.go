package booking

import (
	entitytypes "github.com/mklfarha/mfconsult/entity/types"

	"github.com/mklfarha/mfconsult/entity/booking_intake"
	"github.com/mklfarha/mfconsult/entity/booking_payment"
	"github.com/mklfarha/mfconsult/entity/booking_scheduling"
)

func (e Booking) FieldIdentifierToTypeMap() map[string]entitytypes.FieldType {
	return map[string]entitytypes.FieldType{
		"id":                  entitytypes.StringFieldType,
		"client_id":           entitytypes.StringFieldType,
		"status":              entitytypes.SingleEnumFieldType,
		"review_decision":     entitytypes.SingleEnumFieldType,
		"reviewed_at":         entitytypes.TimestampFieldType,
		"decline_reason":      entitytypes.StringFieldType,
		"pay_link_token":      entitytypes.StringFieldType,
		"pay_link_expires_at": entitytypes.TimestampFieldType,
		"portal_token":        entitytypes.StringFieldType,
		"intake":              entitytypes.SingleDependantEntityFieldType,
		"payment":             entitytypes.SingleDependantEntityFieldType,
		"scheduling":          entitytypes.SingleDependantEntityFieldType,
		"terms_version":       entitytypes.StringFieldType,
		"terms_accepted_at":   entitytypes.TimestampFieldType,
		"terms_accepted_ip":   entitytypes.StringFieldType,
		"created_at":          entitytypes.TimestampFieldType,
		"updated_at":          entitytypes.TimestampFieldType,
	}
}

func (e Booking) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "id")
	res = append(res, "client_id")
	res = append(res, "status")
	res = append(res, "review_decision")
	res = append(res, "reviewed_at")
	res = append(res, "decline_reason")
	res = append(res, "pay_link_token")
	res = append(res, "pay_link_expires_at")
	res = append(res, "portal_token")
	res = append(res, "intake")
	res = append(res, "payment")
	res = append(res, "scheduling")
	res = append(res, "terms_version")
	res = append(res, "terms_accepted_at")
	res = append(res, "terms_accepted_ip")
	res = append(res, "created_at")
	res = append(res, "updated_at")

	return res
}

func (e Booking) DependantFieldIdentifierToTypeMap() map[string]map[string]entitytypes.FieldType {
	res := make(map[string]map[string]entitytypes.FieldType)

	res["intake"] = booking_intake.BookingIntake{}.FieldIdentifierToTypeMap()
	res["payment"] = booking_payment.BookingPayment{}.FieldIdentifierToTypeMap()
	res["scheduling"] = booking_scheduling.BookingScheduling{}.FieldIdentifierToTypeMap()
	return res
}

func (e Booking) EntityIdentifier() string {
	return "booking"
}

func (e Booking) PrimaryKeyIdentifiers() []string {
	return []string{
		"id",
	}
}

func (e Booking) ArrayFieldIdentifierToType() map[string]entitytypes.FieldType {
	res := make(map[string]entitytypes.FieldType)

	return res
}
