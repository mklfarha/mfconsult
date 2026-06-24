package repository

import (
	"context"

	"database/sql"

	mfconsultdb "github.com/mklfarha/mfconsult/core/repository/gen"
	"github.com/mklfarha/mfconsult/core/repository/list"
)

type Implementation struct {
	Queries *mfconsultdb.Queries
	DB      *sql.DB
	List    *list.Implementation
}

func New(db *sql.DB) *Implementation {
	queries := mfconsultdb.New(db)
	return &Implementation{
		Queries: queries,
		DB:      db,
		List:    list.New(),
	}
}

func (i *Implementation) BuildListEntityQuery(ctx context.Context, request list.ListRequest, entity list.ListEntity, onlyCount bool) (string, error) {
	return i.List.BuildListEntityQuery(ctx, request, entity, onlyCount)
}
