package http

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/infra/http/middleware"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/application"
	db "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/memory"
	identity "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/presentation"
	monitor "github.com/GustavoCesarSantos/retro-board-api/internal/modules/monitor/presentation"
)

var userRepository = db.NewUserRepository()

var createAuthToken = application.NewCreateAuthToken()
var findUserByEmail = application.NewFindUserByEmail(userRepository)
var saveUser = application.NewSaveUser(userRepository)

var healthcheck = monitor.NewHealthcheck()
var signinUser = identity.NewSigninUser(createAuthToken, findUserByEmail)
var signinWithGoogle = identity.NewSigninWithGoogle()
var signinWithGoogleCallback = identity.NewSigninWithGoogleCallback(findUserByEmail, saveUser)

func routes(mux *http.ServeMux) http.Handler {
	mux.Handle("GET /v1/healthcheck", http.HandlerFunc(healthcheck.Handle))
	mux.Handle("POST /v1/auth/signin", http.HandlerFunc(signinUser.Handle))
	mux.Handle("GET /v1/auth/signin/google", http.HandlerFunc(signinWithGoogle.Handle))
	mux.Handle("GET /v1/auth/signin/google/callback", http.HandlerFunc(signinWithGoogleCallback.Handle))
	return middleware.RecoverPanic(middleware.EnableCORS(mux))
}
