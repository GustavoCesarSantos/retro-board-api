package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/GustavoCesarSantos/retro-board-api/docs"
	"github.com/GustavoCesarSantos/retro-board-api/internal/infra/http/middleware"
	board "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/handlers"
	identity "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/presentation/handlers"
	monitor "github.com/GustavoCesarSantos/retro-board-api/internal/modules/monitor/presentation/handlers"
	poll "github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/presentation/handlers"
	realtime "github.com/GustavoCesarSantos/retro-board-api/internal/modules/realtime/presentation/handlers"
	team "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/presentation/handlers"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

func NewRouter(
	userAuthenticator *middleware.UserAuthenticator,
	teamMemberValidator *middleware.TeamMemberValidator,
	boardValidator *middleware.BoardValidator,
	pollValidator *middleware.PollValidator,
	monitor *monitor.Handlers,
	identity *identity.Handlers,
	team *team.Handlers,
	board *board.Handlers,
	poll *poll.Handlers,
	realtime *realtime.Handlers,
) http.Handler {
	metadataErr := utils.Envelope{
		"file": "routes.go",
		"func": "routes.routes",
		"line": 0,
	}
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metadataErr["line"] = 42
		utils.NotFoundResponse(w, r, metadataErr)
	})
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metadataErr["line"] = 47
		utils.MethodNotAllowedResponse(w, r, metadataErr)
	})

    router.Handler(http.MethodGet, "/metrics", promhttp.Handler())

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", monitor.Healthcheck.Handle)

	router.HandlerFunc(http.MethodPost, "/v1/auth/refresh-token", identity.RefreshAuthToken.Handle)
	router.HandlerFunc(http.MethodPost, "/v1/auth/signin", identity.SigninUser.Handle)
	router.HandlerFunc(http.MethodGet, "/v1/auth/signin/google", identity.SigninWithGoogle.Handle)
	router.HandlerFunc(
		http.MethodGet,
		"/v1/auth/signin/google/callback",
		identity.SigninWithGoogleCallback.Handle,
	)
	router.HandlerFunc(
		http.MethodPost,
		"/v1/auth/signout",
		userAuthenticator.Authenticate(identity.SignoutUser.Handle),
	)

	router.HandlerFunc(
		http.MethodPost,
		"/v1/teams",
		userAuthenticator.Authenticate(team.CreateTeam.Handle),
	)
	router.HandlerFunc(
		http.MethodGet,
		"/v1/teams",
		userAuthenticator.Authenticate(team.ListAllTeams.Handle),
	)
	router.HandlerFunc(
		http.MethodGet,
		"/v1/teams/:teamId",
		userAuthenticator.Authenticate(team.ShowTeam.Handle),
	)
	router.HandlerFunc(
		http.MethodPatch,
		"/v1/teams/:teamId",
		userAuthenticator.Authenticate(team.EditTeam.Handle),
	)
	router.HandlerFunc(
		http.MethodDelete,
		"/v1/teams/:teamId",
		userAuthenticator.Authenticate(team.DeleteTeam.Handle),
	)

	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/members", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			team.ListAllMembers.Handle,
		),
	))
	router.HandlerFunc(
		http.MethodDelete,
		"/v1/teams/:teamId/members/:memberId",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				team.RemoveMemberFromTeam.Handle,
			),
		),
	)
	router.HandlerFunc(
		http.MethodPatch,
		"/v1/teams/:teamId/members/:memberId/roles",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				team.ChangeMemberRole.Handle,
			),
		),
	)
	router.HandlerFunc(
		http.MethodPost,
		"/v1/teams/:teamId/members/invite",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				team.AddMemberToTeam.Handle,
			),
		),
	)
	router.HandlerFunc(
		http.MethodPatch,
		"/v1/teams/:teamId/members/:memberId/accept-invite",
		userAuthenticator.Authenticate(team.EditMember.Handle),
	)

	router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/boards", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			board.CreateBoard.Handle,
		),
	))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/boards", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			board.ListAllBoards.Handle,
		),
	))
	router.HandlerFunc(
		http.MethodGet,
		"/v1/teams/:teamId/boards/:boardId",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				boardValidator.EnsureBoardOwnership(
					board.ListBoard.Handle,
				),
			),
		),
	)
	router.HandlerFunc(
		http.MethodPut,
		"/v1/teams/:teamId/boards/:boardId",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				boardValidator.EnsureBoardOwnership(
					board.EditBoard.Handle,
				),
			),
		),
	)
	router.HandlerFunc(
		http.MethodDelete,
		"/v1/teams/:teamId/boards/:boardId",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				boardValidator.EnsureBoardOwnership(
					board.DeleteBoard.Handle,
				),
			),
		),
	)

	router.HandlerFunc(
		http.MethodPost,
		"/v1/teams/:teamId/boards/:boardId/columns",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				boardValidator.EnsureBoardOwnership(
					board.CreateColumn.Handle,
				),
			),
		),
	)
	router.HandlerFunc(
		http.MethodGet,
		"/v1/teams/:teamId/boards/:boardId/columns",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				boardValidator.EnsureBoardOwnership(
					board.ListAllColumns.Handle,
				),
			),
		),
	)
	router.HandlerFunc(
		http.MethodGet,
		"/v1/teams/:teamId/boards/:boardId/columns/:columnId",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				boardValidator.EnsureBoardOwnership(
					board.ListColumn.Handle,
				),
			),
		),
	)
	router.HandlerFunc(
		http.MethodPut,
		"/v1/teams/:teamId/boards/:boardId/columns/:columnId",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				boardValidator.EnsureBoardOwnership(
					boardValidator.EnsureColumnOwnership(
						board.EditColumn.Handle,
					),
				),
			),
		),
	)
	router.HandlerFunc(
		http.MethodDelete,
		"/v1/teams/:teamId/boards/:boardId/columns/:columnId",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				boardValidator.EnsureBoardOwnership(
					boardValidator.EnsureColumnOwnership(
						board.DeleteColumn.Handle,
					),
				),
			),
		),
	)

	router.HandlerFunc(
		http.MethodPut,
		"/v1/teams/:teamId/boards/:boardId/columns/:columnId/move",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				boardValidator.EnsureBoardOwnership(
					boardValidator.EnsureColumnOwnership(
						board.MoveColumnToAnotherPosition.Handle,
					),
				),
			),
		),
	)

	router.HandlerFunc(
		http.MethodPost,
		"/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				boardValidator.EnsureBoardOwnership(
					boardValidator.EnsureColumnOwnership(
						board.CreateCard.Handle,
					),
				),
			),
		),
	)
	router.HandlerFunc(
		http.MethodGet,
		"/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				boardValidator.EnsureBoardOwnership(
					boardValidator.EnsureColumnOwnership(
						board.ListAllCards.Handle,
					),
				),
			),
		),
	)
	router.HandlerFunc(
		http.MethodGet,
		"/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				boardValidator.EnsureBoardOwnership(
					boardValidator.EnsureColumnOwnership(
						boardValidator.EnsureCardOwnership(
							board.ListCard.Handle,
						),
					),
				),
			),
		),
	)
	router.HandlerFunc(
		http.MethodPut,
		"/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				boardValidator.EnsureBoardOwnership(
					boardValidator.EnsureColumnOwnership(
						boardValidator.EnsureCardOwnership(
							board.EditCard.Handle,
						),
					),
				),
			),
		),
	)
	router.HandlerFunc(
		http.MethodDelete,
		"/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				boardValidator.EnsureBoardOwnership(
					boardValidator.EnsureColumnOwnership(
						boardValidator.EnsureCardOwnership(
							board.DeleteCard.Handle,
						),
					),
				),
			),
		),
	)
	router.HandlerFunc(
		http.MethodPut,
		"/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId/move",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				boardValidator.EnsureBoardOwnership(
					boardValidator.EnsureColumnOwnership(
						boardValidator.EnsureCardOwnership(
							board.MoveCardToAnotherColumn.Handle,
						),
					),
				),
			),
		),
	)

	router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/polls", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			poll.CreatePoll.Handle,
		),
	))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/polls", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			poll.ListAllPolls.Handle,
		),
	))
	router.HandlerFunc(
		http.MethodGet,
		"/v1/teams/:teamId/polls/:pollId",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				pollValidator.EnsurePollOwnership(
					poll.ListPoll.Handle,
				),
			),
		),
	)
	router.HandlerFunc(
		http.MethodGet,
		"/v1/teams/:teamId/polls/:pollId/result",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				pollValidator.EnsurePollOwnership(
					poll.ShowPollResult.Handle,
				),
			),
		),
	)

	router.HandlerFunc(
		http.MethodDelete,
		"/v1/teams/:teamId/polls/:pollId/options/:optionId",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				pollValidator.EnsurePollOwnership(
					pollValidator.EnsureOptionOwnership(
						poll.DeleteOption.Handle,
					),
				),
			),
		),
	)

	router.HandlerFunc(
		http.MethodPost,
		"/v1/teams/:teamId/polls/:pollId/options/:optionId/vote",
		userAuthenticator.Authenticate(
			teamMemberValidator.EnsureMemberAccess(
				pollValidator.EnsurePollOwnership(
					pollValidator.EnsureOptionOwnership(
						poll.Vote.Handle,
					),
				),
			),
		),
	)

	router.Handler(http.MethodGet, "/v1/docs/*filepath", httpSwagger.WrapHandler)

	router.HandlerFunc(
		http.MethodGet,
		"/v1/ws/teams/:teamId/boards/:boardId",
		userAuthenticator.AuthenticateWebSocket(
			teamMemberValidator.EnsureMemberAccess(
				realtime.ConnectToBoardRoom.Handle,
			),
		),
	)

	return middleware.RecoverPanic(middleware.EnableCORS(router))
}
