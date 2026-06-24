package booking_document

import (
	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"
	main_entity "github.com/mklfarha/mfconsult/entity/booking_document"

	"github.com/mklfarha/mfconsult/entity/mapper"

	"github.com/guregu/null/v6"
	"github.com/mklfarha/mfconsult/enum"
)

func mapModelsToEntities(models []mfconsultdb.BookingDocument) []main_entity.BookingDocument {
	result := []main_entity.BookingDocument{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(m mfconsultdb.BookingDocument) main_entity.BookingDocument {
	return main_entity.BookingDocument{
		ID:         mapper.StringToUUID(m.ID),
		BookingId:  mapper.StringToUUID(m.BookingId),
		Kind:       enum.DocumentKind(m.Kind.Int64),
		URL:        null.NewString(m.URL.String, m.URL.Valid),
		Label:      null.NewString(m.Label.String, m.Label.Valid),
		PurgeAfter: null.NewTime(m.PurgeAfter.Time, m.PurgeAfter.Valid),
		CreatedAt:  null.NewTime(m.CreatedAt.Time, m.CreatedAt.Valid),
	}
}
