package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/GustavoCesarSantos/retro-board-api/internal/infra/http/middleware"
	boardApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	boardDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/memory"
	board "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation"
	identityApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/application"
	userDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/memory"
	identity "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/presentation"
	monitor "github.com/GustavoCesarSantos/retro-board-api/internal/modules/monitor/presentation"
	teamApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	teamDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/memory"
	team "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/presentation"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

var boardRepository = boardDb.NewBoardRepository()
var cardRepository = boardDb.NewCardRepository()
var columnRepository = boardDb.NewColumnRepository()
var teamRepository = teamDb.NewTeamRepository()
var teamMemberRepository = teamDb.NewTeamMemberRepository()
var userRepository = userDb.NewUserRepository()

var createAuthToken = identityApplication.NewCreateAuthToken()
var ensureBoardOwnership = boardApplication.NewEnsureBoardOwnership(boardRepository)
var ensureColumnOwnership = boardApplication.NewEnsureColumnOwnership(columnRepository)
var findAllColumns = boardApplication.NewFindAllColumns(columnRepository)
var findAllTeams = teamApplication.NewFindAllTeams(teamRepository)
var findTeam = teamApplication.NewFindTeam(teamRepository)
var findUserByEmail = identityApplication.NewFindUserByEmail(userRepository)
var getNextColumnPosition = boardApplication.NewGetNextColumnPosition(columnRepository)
var removeMember = teamApplication.NewRemoveMember(teamMemberRepository)
var saveBoard = boardApplication.NewSaveBoard(boardRepository)
var saveCard = boardApplication.NewSaveCard(cardRepository)
var saveColumn = boardApplication.NewSaveColumn(columnRepository)
var saveMember = teamApplication.NewSaveMember(teamMemberRepository)
var saveTeam = teamApplication.NewSaveTeam(teamRepository)
var saveUser = identityApplication.NewSaveUser(userRepository)
var updateRole = teamApplication.NewUpdateRole(teamMemberRepository)

var addMemberToTeam = team.NewAddMemberToTeam(saveMember)
var createBoard = board.NewCreateBoard(saveBoard)
var createCard = board.NewCreateCard(
	ensureBoardOwnership,
	ensureColumnOwnership,
	saveCard,
)
var createColumn = board.NewCreateColumn(
	ensureBoardOwnership,
	findAllColumns,
	getNextColumnPosition,
	saveColumn,
)
var createTeam = team.NewCreateTeam(saveTeam)
var healthcheck = monitor.NewHealthcheck()
var listAllTeams = team.NewListAllTeams(findAllTeams)
var removeMemberFromTeam = team.NewRemoveMemberFromTeam(removeMember)
var showTeam = team.NewShowTeam(findTeam)
var signinUser = identity.NewSigninUser(createAuthToken, findUserByEmail)
var signinWithGoogle = identity.NewSigninWithGoogle()
var signinWithGoogleCallback = identity.NewSigninWithGoogleCallback(findUserByEmail, saveUser)
var changeMemberRole = team.NewChangeMemberRole(updateRole)

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
    router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId", showTeam.Handle)

    router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/members", addMemberToTeam.Handle)
    router.HandlerFunc(http.MethodDelete, "/v1/teams/:teamId/members/:memberId", removeMemberFromTeam.Handle)
    router.HandlerFunc(http.MethodPut, "/v1/teams/:teamId/members/:memberId/change-role", changeMemberRole.Handle)
	
    router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/boards", createBoard.Handle)
    router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/boards/:boardId/columns", createCard.Handle)
    router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards", createColumn.Handle)

	return middleware.RecoverPanic(middleware.EnableCORS(router))
}
