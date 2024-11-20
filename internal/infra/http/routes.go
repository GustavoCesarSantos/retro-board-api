package http

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/infra/http/middleware"
	monitor "github.com/GustavoCesarSantos/retro-board-api/internal/modules/monitor/presentation"
)

func routes(mux *http.ServeMux) http.Handler {
	mux.Handle("GET /v1/healthcheck", http.HandlerFunc(monitor.Healthcheck))
	return middleware.EnableCORS(mux)
}