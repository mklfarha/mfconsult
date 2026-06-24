package booking_recap

import (
	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"
	main_entity "github.com/mklfarha/mfconsult/entity/booking_recap"

	"github.com/mklfarha/mfconsult/entity/mapper"

	"github.com/guregu/null/v6"
)

func mapModelsToEntities(models []mfconsultdb.BookingRecap) []main_entity.BookingRecap {
	result := []main_entity.BookingRecap{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(m mfconsultdb.BookingRecap) main_entity.BookingRecap {
	return main_entity.BookingRecap{
		ID:          mapper.StringToUUID(m.ID),
		BookingId:   mapper.StringToUUID(m.BookingId),
		Body:        null.NewString(m.Body.String, m.Body.Valid),
		PublishedAt: null.NewTime(m.PublishedAt.Time, m.PublishedAt.Valid),
		CreatedAt:   null.NewTime(m.CreatedAt.Time, m.CreatedAt.Valid),
	}
}
