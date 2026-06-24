package booking_payment

import (
	entitytypes "github.com/mklfarha/mfconsult/entity/types"
)

func (e BookingPayment) FieldIdentifierToTypeMap() map[string]entitytypes.FieldType {
	return map[string]entitytypes.FieldType{
		"amount_cents":   entitytypes.IntFieldType,
		"currency":       entitytypes.StringFieldType,
		"payment_status": entitytypes.SingleEnumFieldType,
		"stripe_ref":     entitytypes.StringFieldType,
	}
}

func (e BookingPayment) OrderedFieldIdentifiers() []string {
	res := []string{}
	res = append(res, "amount_cents")
	res = append(res, "currency")
	res = append(res, "payment_status")
	res = append(res, "stripe_ref")

	return res
}

func (e BookingPayment) DependantFieldIdentifierToTypeMap() map[string]map[string]entitytypes.FieldType {
	res := make(map[string]map[string]entitytypes.FieldType)

	return res
}

func (e BookingPayment) EntityIdentifier() string {
	return "booking_payment"
}

func (e BookingPayment) PrimaryKeyIdentifiers() []string {
	return []string{}
}

func (e BookingPayment) ArrayFieldIdentifierToType() map[string]entitytypes.FieldType {
	res := make(map[string]entitytypes.FieldType)

	return res
}
