package core

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

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

	"github.com/mklfarha/mfconsult/core/module/engagement_agreement"

	"github.com/mklfarha/mfconsult/core/module/webhook_event"

	"github.com/mklfarha/mfconsult/core/module/engagement_inquiry"

	"github.com/mklfarha/mfconsult/core/module/magic_link"

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

	engagement_agreement engagement_agreement.Module

	webhook_event webhook_event.Module

	engagement_inquiry engagement_inquiry.Module

	magic_link magic_link.Module
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

	logger := params.Logger
	if logger == nil {
		logger = zap.NewNop()
	}

	if params.Lifecycle != nil {
		params.Lifecycle.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				// Warm the connection pool and fail fast with a clear error if the
				// database is unreachable or credentials/permissions are wrong,
				// instead of surfacing it later as a request-time timeout.
				pingCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
				defer cancel()
				if err := db.PingContext(pingCtx); err != nil {
					logger.Error("failed to connect to database on startup", zap.Error(err))
					return fmt.Errorf("error connecting to DB: %w", err)
				}
				logger.Info("connected to database")
				return nil
			},
			OnStop: func(ctx context.Context) error {
				db.Close()
				return nil
			},
		})
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

func (i Implementation) EngagementAgreement() engagement_agreement.Module {
	if i.engagement_agreement == nil {
		i.engagement_agreement = engagement_agreement.New(coretypes.ModuleParams{
			Repository: i.repository,
			Logger:     i.logger,
		})
	}
	return i.engagement_agreement
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

func (i Implementation) EngagementInquiry() engagement_inquiry.Module {
	if i.engagement_inquiry == nil {
		i.engagement_inquiry = engagement_inquiry.New(coretypes.ModuleParams{
			Repository: i.repository,
			Logger:     i.logger,
		})
	}
	return i.engagement_inquiry
}

func (i Implementation) MagicLink() magic_link.Module {
	if i.magic_link == nil {
		i.magic_link = magic_link.New(coretypes.ModuleParams{
			Repository: i.repository,
			Logger:     i.logger,
		})
	}
	return i.magic_link
}
