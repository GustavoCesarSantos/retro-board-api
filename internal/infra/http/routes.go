package http

import (
	"net/http"

	monitor "github.com/GustavoCesarSantos/retro-board-api/internal/modules/monitor/presentation"
)

func Routes(mux *http.ServeMux) {
	mux.Handle("GET /health", http.HandlerFunc(monitor.Healthcheck))
}