package core

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	mfconsultconfig "github.com/mklfarha/mfconsult/config"
	coretypes "github.com/mklfarha/mfconsult/core/types"
	"go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/mklfarha/mfconsult/core/module/client"

	"github.com/mklfarha/mfconsult/core/module/booking"

	"github.com/mklfarha/mfconsult/core/module/booking_document"

	"github.com/mklfarha/mfconsult/core/module/booking_recap"

	"github.com/mklfarha/mfconsult/core/module/nda_document"

	"github.com/mklfarha/mfconsult/core/module/webhook_event"

	"github.com/mklfarha/mfconsult/core/repository"
)

type Implementation struct {
	db         *sql.DB
	repository *repository.Implementation
	logger     *zap.Logger

	client client.Module

	booking booking.Module

	booking_document booking_document.Module

	booking_recap booking_recap.Module

	nda_document nda_document.Module

	webhook_event webhook_event.Module
}

type Params struct {
	fx.In
	Provider  config.Provider
	Lifecycle fx.Lifecycle
	Logger    *zap.Logger `optional:"true"`
}

func New(params Params) (*Implementation, error) {

	var dbs mfconsultconfig.DBs
	if err := params.Provider.Get("db").Populate(&dbs); err != nil {
		return nil, err
	}

	if len(dbs) == 0 {
		return nil, errors.New("db configuration not found")
	}

	dbconfig := dbs[0]
	db, err := sql.Open(dbconfig.Driver, dbconfig.Path())
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %v", err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(0)
	repository := repository.New(db)

	if params.Lifecycle != nil {
		params.Lifecycle.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				db.Close()
				return nil
			},
		})
	}

	logger := params.Logger
	if logger == nil {
		logger = zap.NewNop()
	}

	return &Implementation{
		db:         db,
		repository: repository,
		logger:     logger,
	}, nil
}

func (i *Implementation) Destroy() {
	i.db.Close()
}

func (i *Implementation) DB() *sql.DB {
	return i.db
}

func (i Implementation) Client() client.Module {
	if i.client == nil {
		i.client = client.New(coretypes.ModuleParams{
			Repository: i.repository,
			Logger:     i.logger,
		})
	}
	return i.client
}

func (i Implementation) Booking() booking.Module {
	if i.booking == nil {
		i.booking = booking.New(coretypes.ModuleParams{
			Repository: i.repository,
			Logger:     i.logger,
		})
	}
	return i.booking
}

func (i Implementation) BookingDocument() booking_document.Module {
	if i.booking_document == nil {
		i.booking_document = booking_document.New(coretypes.ModuleParams{
			Repository: i.repository,
			Logger:     i.logger,
		})
	}
	return i.booking_document
}

func (i Implementation) BookingRecap() booking_recap.Module {
	if i.booking_recap == nil {
		i.booking_recap = booking_recap.New(coretypes.ModuleParams{
			Repository: i.repository,
			Logger:     i.logger,
		})
	}
	return i.booking_recap
}

func (i Implementation) NdaDocument() nda_document.Module {
	if i.nda_document == nil {
		i.nda_document = nda_document.New(coretypes.ModuleParams{
			Repository: i.repository,
			Logger:     i.logger,
		})
	}
	return i.nda_document
}

func (i Implementation) WebhookEvent() webhook_event.Module {
	if i.webhook_event == nil {
		i.webhook_event = webhook_event.New(coretypes.ModuleParams{
			Repository: i.repository,
			Logger:     i.logger,
		})
	}
	return i.webhook_event
}
