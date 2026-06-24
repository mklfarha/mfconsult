package client

import (
	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"
	main_entity "github.com/mklfarha/mfconsult/entity/client"

	"github.com/mklfarha/mfconsult/entity/mapper"

	"github.com/guregu/null/v6"
)

func mapModelsToEntities(models []mfconsultdb.Client) []main_entity.Client {
	result := []main_entity.Client{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(m mfconsultdb.Client) main_entity.Client {
	return main_entity.Client{
		ID:        mapper.StringToUUID(m.ID),
		Name:      m.Name,
		Email:     m.Email,
		Timezone:  null.NewString(m.Timezone.String, m.Timezone.Valid),
		Notes:     null.NewString(m.Notes.String, m.Notes.Valid),
		CreatedAt: null.NewTime(m.CreatedAt.Time, m.CreatedAt.Valid),
		UpdatedAt: null.NewTime(m.UpdatedAt.Time, m.UpdatedAt.Valid),
	}
}
