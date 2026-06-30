package booking

import (
	"context"
	"github.com/mklfarha/mfconsult/core/module/booking/types"
	"github.com/mklfarha/mfconsult/core/repository"
	coretypes "github.com/mklfarha/mfconsult/core/types"
	gocache "github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"sync"
	"time"
)

type Module interface {
	FetchBookingById(ctx context.Context, req types.FetchBookingByIdRequest, opts ...Option) (types.FetchBookingByIdResponse, error)

	List(ctx context.Context, req types.ListRequest, opts ...Option) (types.ListResponse, error)

	Upsert(ctx context.Context, req types.UpsertRequest, opts ...Option) (types.UpsertResponse, error)
	Insert(ctx context.Context, req types.UpsertRequest, opts ...Option) (types.UpsertResponse, error)
	Update(ctx context.Context, req types.UpsertRequest, opts ...Option) (types.UpsertResponse, error)
	Delete(ctx context.Context, req types.DeleteRequest, opts ...Option) error
}

type module struct {
	mu         sync.Mutex
	sg         singleflight.Group
	cache      *gocache.Cache
	repository *repository.Implementation
	logger     *zap.Logger
}

func New(params coretypes.ModuleParams) Module {
	return &module{
		repository: params.Repository,
		logger:     params.Logger,
		cache:      gocache.New(30*time.Second, 5*time.Minute),
	}
}
