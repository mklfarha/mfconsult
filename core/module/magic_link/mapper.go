package magic_link

import (
	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"
	main_entity "github.com/mklfarha/mfconsult/entity/magic_link"

	"github.com/mklfarha/mfconsult/entity/mapper"

	"github.com/guregu/null/v6"
	"github.com/mklfarha/mfconsult/enum"
)

func mapModelsToEntities(models []mfconsultdb.MagicLink) []main_entity.MagicLink {
	result := []main_entity.MagicLink{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(m mfconsultdb.MagicLink) main_entity.MagicLink {
	return main_entity.MagicLink{
		ID:         mapper.StringToUUID(m.ID),
		ClientId:   mapper.StringToUUID(m.ClientId),
		Email:      null.NewString(m.Email.String, m.Email.Valid),
		Token:      m.Token,
		Purpose:    enum.MagicLinkPurpose(m.Purpose.Int64),
		ExpiresAt:  null.NewTime(m.ExpiresAt.Time, m.ExpiresAt.Valid),
		ConsumedAt: null.NewTime(m.ConsumedAt.Time, m.ConsumedAt.Valid),
		CreatedAt:  null.NewTime(m.CreatedAt.Time, m.CreatedAt.Valid),
		CreatedIp:  null.NewString(m.CreatedIp.String, m.CreatedIp.Valid),
	}
}
