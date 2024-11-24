package http

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/infra/http/middleware"
	identity "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/presentation"
	monitor "github.com/GustavoCesarSantos/retro-board-api/internal/modules/monitor/presentation"
)

var healthcheck = monitor.NewHealthcheck()
var signinWithGoogle = identity.NewSigninWithGoogle()
var signinWithGoogleCallback = identity.NewSigninWithGoogleCallback()

func routes(mux *http.ServeMux) http.Handler {
	mux.Handle("GET /v1/healthcheck", http.HandlerFunc(healthcheck.Handle))
	mux.Handle("GET /v1/auth/signin/google", http.HandlerFunc(signinWithGoogle.Handle))
	mux.Handle("GET /v1/auth/signin/google/callback", http.HandlerFunc(signinWithGoogleCallback.Handle))
	return middleware.RecoverPanic(middleware.EnableCORS(mux))
}
