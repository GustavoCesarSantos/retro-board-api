package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.uber.org/fx"

	"github.com/GustavoCesarSantos/retro-board-api/internal/infra/database"
	httpServer "github.com/GustavoCesarSantos/retro-board-api/internal/infra/http"
	"github.com/GustavoCesarSantos/retro-board-api/internal/infra/http/middleware"
	"github.com/GustavoCesarSantos/retro-board-api/internal/infra/oauth2"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/monitor"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/realtime"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/providers"
)

// @title Retro Board API
// @version 1.0
// @description Documentação da API.
// @termsOfService http://swagger.io/terms/

// @host localhost:9000
// @BasePath /v1

// @securityDefinitions.bearerAuth
// @securityDefinitions.apiKey
// @name Authorization
// @in header

func main() {
	if os.Getenv("LOAD_ENV_FILE") == "true" {
		if err := godotenv.Load(); err != nil {
			slog.Error("failed to load .env file", "error", err)
			os.Exit(1)
		}
	}

	app := fx.New(
		middleware.Module,
		providers.Module,
		board.Module,
		identity.Module,
		monitor.Module,
		poll.Module,
		realtime.Module,
		team.Module,

		fx.Provide(
			database.OpenDB,
			httpServer.NewRouter,
			httpServer.NewServer,
		),

		fx.Invoke(
			oauth2.SetProvider, 
			startServer,
		),
	)

	app.Run()
}

func startServer(lc fx.Lifecycle, server *httpServer.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				shutdownError := make(chan error)
				var wg sync.WaitGroup
				srv := server.GetHTTPServer()
				go server.HandleShutdown(srv, shutdownError, &wg)
				if err := server.Start(); err != nil {
					if !errors.Is(err, http.ErrServerClosed) {
						slog.Error(err.Error())
						os.Exit(1)
					}
					err = <-shutdownError
					if err != nil {
						slog.Error(err.Error())
						os.Exit(1)
					}
					slog.Info("stopped server", "addr", srv.Addr)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
