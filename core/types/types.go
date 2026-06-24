package types

import (
	"github.com/mklfarha/mfconsult/core/repository"
	"go.uber.org/zap"
)

type ModuleParams struct {
	Repository *repository.Implementation
	Logger     *zap.Logger
}
