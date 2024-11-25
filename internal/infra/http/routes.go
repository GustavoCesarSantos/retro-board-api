package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/GustavoCesarSantos/retro-board-api/internal/infra/http/middleware"
	identityApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/application"
	userDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/memory"
	identity "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/presentation"
	monitor "github.com/GustavoCesarSantos/retro-board-api/internal/modules/monitor/presentation"
	teamApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	teamDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/memory"
	team "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/presentation"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

var teamRepository = teamDb.NewTeamRepository()
var userRepository = userDb.NewUserRepository()

var createAuthToken = identityApplication.NewCreateAuthToken()
var findAllTeams = teamApplication.NewFindAllTeams(teamRepository)
var findTeam = teamApplication.NewFindTeam(teamRepository)
var findUserByEmail = identityApplication.NewFindUserByEmail(userRepository)
var saveTeam = teamApplication.NewSaveTeam(teamRepository)
var saveUser = identityApplication.NewSaveUser(userRepository)

var createTeam = team.NewCreateTeam(saveTeam)
var healthcheck = monitor.NewHealthcheck()
var listAllTeams = team.NewListAllTeams(findAllTeams)
var showTeam = team.NewShowTeam(findTeam)
var signinUser = identity.NewSigninUser(createAuthToken, findUserByEmail)
var signinWithGoogle = identity.NewSigninWithGoogle()
var signinWithGoogleCallback = identity.NewSigninWithGoogleCallback(findUserByEmail, saveUser)

func routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(utils.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(utils.MethodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", healthcheck.Handle)

	router.HandlerFunc(http.MethodPost, "/v1/auth/signin", signinUser.Handle)
	router.HandlerFunc(http.MethodGet, "/v1/auth/signin/google", signinWithGoogle.Handle)
	router.HandlerFunc(http.MethodGet, "/v1/auth/signin/google/callback", signinWithGoogleCallback.Handle)

	router.HandlerFunc(http.MethodPost, "/v1/teams", createTeam.Handle)
	router.HandlerFunc(http.MethodGet, "/v1/teams", listAllTeams.Handle)
    router.HandlerFunc(http.MethodGet, "/v1/teams/:id", showTeam.Handle)

	return middleware.RecoverPanic(middleware.EnableCORS(router))
}
