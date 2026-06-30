package main

import (
	"go.uber.org/config"
	"net/http"

	mfconsultconfig "github.com/mklfarha/mfconsult/config"
	"github.com/mklfarha/mfconsult/core"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			zap.NewProduction,
			mfconsultconfig.New,
			core.New,
		),

		fx.Invoke(httpServer),
	).Run()
}

func httpServer(config config.Provider, logger *zap.Logger) {
	// http port from config
	httpPort := config.Get("ports.http").String()

	go http.ListenAndServe(":"+httpPort, nil)

	logger.Info(`Serving HTTP on PORT: %s`, zap.String("port", httpPort))
}
