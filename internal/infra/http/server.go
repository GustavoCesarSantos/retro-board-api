package http

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/configs"
)

type Server struct {
	server *http.Server
}

func NewServer(router http.Handler) *Server {
	serverConfigs := configs.LoadServerConfig()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return &Server{
		server: &http.Server{
			Addr:         fmt.Sprintf(":%d", serverConfigs.Port),
			Handler:      router,
			IdleTimeout:  time.Minute,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			MaxHeaderBytes: 1 << 20,
			ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		},
	}
}

func (svr *Server) GetHTTPServer() *http.Server {
	return svr.server
}

func (svr *Server) HandleShutdown(srv *http.Server, shutdownError chan error, wg *sync.WaitGroup) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit
	slog.Info("shutting down server", "signal", s.String())
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		shutdownError <- err
	}
	slog.Info("completing background tasks", "addr", srv.Addr)
	wg.Wait()
	shutdownError <- nil
}

func (svr *Server) Shutdown(ctx context.Context) error {
	slog.Info("shutting down server")
	return svr.server.Shutdown(ctx)
}

func (svr *Server) Start() error {
	slog.Info(fmt.Sprintf("starting server on :%s", svr.server.Addr))
	return svr.server.ListenAndServe()
}
