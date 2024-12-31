package http

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/GustavoCesarSantos/retro-board-api/docs"
	"github.com/GustavoCesarSantos/retro-board-api/internal/infra/http/middleware"
	boardApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	boardDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/memory"
	boardProvider "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/integration/provider"
	board "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/handlers"
	identityApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/application"
	userDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/nativeSql"
	identityProvider "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/integration/provider"
	identity "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/presentation/handlers"
	monitor "github.com/GustavoCesarSantos/retro-board-api/internal/modules/monitor/presentation/handlers"
	pollApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/application"
	pollDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/external/db/memory"
	poll "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/presentation/handlers"
	teamApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	teamDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/nativeSql"
	teamMemberProvider "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/integration/provider"
	team "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/presentation/handlers"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

func routes(db *sql.DB) http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(utils.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(utils.MethodNotAllowedResponse)

	//Init Repos
	boardRepository := boardDb.NewBoardRepository()
	cardRepository := boardDb.NewCardRepository()
	columnRepository := boardDb.NewColumnRepository()
	optionRepository := pollDb.NewOptionRepository()
	pollRepository := pollDb.NewPollRepository()
	teamRepository := teamDb.NewTeamRepository(db)
	teamMemberRepository := teamDb.NewTeamMemberRepository(db)
	userRepository := userDb.NewUserRepository(db)
	voteRepository := pollDb.NewVoteRepository()

	boardPublicApiProvider := boardProvider.NewBoardPublicApiProvider(
		boardRepository,
		cardRepository,
		columnRepository,
	)
	teamMemberPublicApiProvider := teamMemberProvider.NewTeamMemberPublicApiProvider(teamMemberRepository)
	userPublicApiProvider := identityProvider.NewUserPublicApiProvider(userRepository)

	//Init Middlewares
	boardValidator := middleware.NewBoardValidator(boardPublicApiProvider)
	teamMemberValidator := middleware.NewTeamMemberValidator(teamMemberPublicApiProvider)
	userAuthenticator := middleware.NewUserAuthenticator(userPublicApiProvider)

	//Init Use Cases
	countVotesByPollId := pollApplication.NewCountVotesByPollId(
		optionRepository, 
		voteRepository,
	)
	createAuthToken := identityApplication.NewCreateAuthToken()
	decodeAuthToken := identityApplication.NewDecodeAuthToken()
	ensureAdminMembership := teamApplication.NewEnsureAdminMembership(teamMemberRepository)
	ensureOptionOwnership := pollApplication.NewEnsureOptionOwnership(optionRepository)
	ensurePollOwnership := pollApplication.NewEnsurePollOwnership(pollRepository)
	findAllBoards := boardApplication.NewFindAllBoards(boardRepository)
	findAllCards := boardApplication.NewFindAllCards(cardRepository)
	findAllColumns := boardApplication.NewFindAllColumns(columnRepository)
	findAllPolls := pollApplication.NewFindAllPolls(pollRepository)
	findAllTeams := teamApplication.NewFindAllTeams(teamRepository)
	findBoard := boardApplication.NewFindBoard(boardRepository)
	findCard := boardApplication.NewFindCard(cardRepository)
	findMemberInfosByEmail := teamApplication.NewFindMemberInfoByEmail(userPublicApiProvider)
	findPoll := pollApplication.NewFindPoll(pollRepository)
	findTeam := teamApplication.NewFindTeam(teamRepository)
	findUserByEmail := identityApplication.NewFindUserByEmail(userRepository)
	getNextColumnPosition := boardApplication.NewGetNextColumnPosition(
		columnRepository,
	)
	incrementVersion := identityApplication.NewIncrementVersion(userRepository)
	moveCardBetweenColumns := boardApplication.NewMoveCardBetweenColumns(
		cardRepository,
	)
	removeBoard := boardApplication.NewRemoveBoard(boardRepository)
	removeCard := boardApplication.NewRemoveCard(cardRepository)
	removeColumn := boardApplication.NewRemoveColumn(columnRepository)
	removeMember := teamApplication.NewRemoveMember(teamMemberRepository)
	removeOption := pollApplication.NewRemoveOption(optionRepository)
	removeTeam := teamApplication.NewRemoveTeam(teamRepository)
	saveBoard := boardApplication.NewSaveBoard(boardRepository)
	saveCard := boardApplication.NewSaveCard(cardRepository)
	saveColumn := boardApplication.NewSaveColumn(columnRepository)
	saveMember := teamApplication.NewSaveMember(teamMemberRepository)
	saveOption := pollApplication.NewSaveOption(optionRepository)
	savePoll := pollApplication.NewSavePoll(pollRepository)
	saveTeam := teamApplication.NewSaveTeam(teamRepository)
	saveUser := identityApplication.NewSaveUser(userRepository)
	saveVote := pollApplication.NewSaveVote(voteRepository)
	updateBoard := boardApplication.NewUpdateBoard(boardRepository)
	updateCard := boardApplication.NewUpdateCard(cardRepository)
	updateColumn := boardApplication.NewUpdateColumn(columnRepository)
	updateRole := teamApplication.NewUpdateRole(teamMemberRepository)

	//Init Handlers
	addMemberToTeam := team.NewAddMemberToTeam(
		ensureAdminMembership, 
		findMemberInfosByEmail, 
		saveMember,
	)
	changeMemberRole := team.NewChangeMemberRole(ensureAdminMembership, updateRole)
	createBoard := board.NewCreateBoard(saveBoard)
	createCard := board.NewCreateCard(
		saveCard,
	)
	createColumn := board.NewCreateColumn(
		findAllColumns,
		getNextColumnPosition,
		saveColumn,
	)
	createPoll := poll.NewCreatePoll(saveOption, savePoll)
	createTeam := team.NewCreateTeam(removeTeam, saveMember, saveTeam)
	deleteBoard := board.NewDeleteBoard(removeBoard)
	deleteCard := board.NewDeleteCard(
		removeCard,
	)
	deleteColumn := board.NewDeleteColumn(
		removeColumn,
	)
	deleteOption := poll.NewDeleteOption(
		ensurePollOwnership, 
		ensureOptionOwnership, 
		removeOption,
	)
	editBoard := board.NewEditBoard(updateBoard)
	editCard := board.NewEditCard(updateCard)
	editColumn := board.NewEditColumn(updateColumn)
	healthcheck := monitor.NewHealthcheck()
	listAllBoards := board.NewListAllBoards(findAllBoards)
	listAllCards := board.NewListAllCards(findAllCards)
	listAllColumns := board.NewListAllColumns(findAllColumns)
	listAllPolls := poll.NewListAllPolls(findAllPolls)
	listAllTeams := team.NewListAllTeams(findAllTeams)
	listBoard := board.NewListBoard(findBoard)
	listCard := board.NewListCard(findCard)
	listPoll := poll.NewListPoll(ensurePollOwnership, findPoll)
	moveCardtoAnotherColumn := board.NewMoveCardtoAnotherColumn(
		moveCardBetweenColumns,
	)
	refreshAuthToken := identity.NewRefreshAuthToken(
		createAuthToken, 
		decodeAuthToken,
		findUserByEmail,
	)
	removeMemberFromTeam := team.NewRemoveMemberFromTeam(
		ensureAdminMembership, 
		removeMember,
	)
	showPollResult := poll.NewShowPollResult(
		ensurePollOwnership, 
		countVotesByPollId,
	)
	showTeam := team.NewShowTeam(findTeam)
	signinUser := identity.NewSigninUser(
		createAuthToken, 
		findUserByEmail, 
		incrementVersion,
	)
	signinWithGoogle := identity.NewSigninWithGoogle()
	signinWithGoogleCallback := identity.NewSigninWithGoogleCallback(
		findUserByEmail, 
		saveUser,
	)
	signoutUser := identity.NewSignoutUser(incrementVersion)
	vote := poll.NewVote(ensurePollOwnership, ensureOptionOwnership, saveVote)


	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", healthcheck.Handle)

	router.HandlerFunc(http.MethodPost, "/v1/auth/refresh-token", refreshAuthToken.Handle)
	router.HandlerFunc(http.MethodPost, "/v1/auth/signin", signinUser.Handle)
	router.HandlerFunc(http.MethodGet, "/v1/auth/signin/google", signinWithGoogle.Handle)
	router.HandlerFunc(http.MethodGet, "/v1/auth/signin/google/callback", signinWithGoogleCallback.Handle)
	router.HandlerFunc(http.MethodPost, "/v1/auth/signout", userAuthenticator.Authenticate(signoutUser.Handle))

	router.HandlerFunc(http.MethodPost, "/v1/teams", userAuthenticator.Authenticate(createTeam.Handle))
	router.HandlerFunc(http.MethodGet, "/v1/teams", userAuthenticator.Authenticate(listAllTeams.Handle))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId", userAuthenticator.Authenticate(showTeam.Handle))

	router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/members", userAuthenticator.Authenticate(addMemberToTeam.Handle))
	router.HandlerFunc(http.MethodDelete, "/v1/teams/:teamId/members/:memberId", userAuthenticator.Authenticate(removeMemberFromTeam.Handle))
	router.HandlerFunc(http.MethodPatch, "/v1/teams/:teamId/members/:memberId/roles", userAuthenticator.Authenticate(changeMemberRole.Handle))

	router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/boards", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			createBoard.Handle,
		),
	))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/boards", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			listAllBoards.Handle,
		),
	))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/boards/:boardId", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			boardValidator.EnsureBoardOwnership(
				listBoard.Handle,
			),
		),
	))
	router.HandlerFunc(http.MethodPut, "/v1/teams/:teamId/boards/:boardId", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			boardValidator.EnsureBoardOwnership(
				editBoard.Handle,
			),
		),
	))
	router.HandlerFunc(http.MethodDelete, "/v1/teams/:teamId/boards/:boardId", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			boardValidator.EnsureBoardOwnership(
				deleteBoard.Handle,
			),
		),
	))

	router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/boards/:boardId/columns", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			boardValidator.EnsureBoardOwnership(
				createColumn.Handle,
			),
		),
	))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/boards/:boardId/columns", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			boardValidator.EnsureBoardOwnership(
				listAllColumns.Handle,
			),
		),
	))
	router.HandlerFunc(http.MethodPut, "/v1/teams/:teamId/boards/:boardId/columns/:columnId", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			boardValidator.EnsureBoardOwnership(
				boardValidator.EnsureColumnOwnership(
					editColumn.Handle,
				),
			),
		),
	))
	router.HandlerFunc(http.MethodDelete, "/v1/teams/:teamId/boards/:boardId/columns/:columnId", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			boardValidator.EnsureBoardOwnership(
				boardValidator.EnsureColumnOwnership(
					deleteColumn.Handle,
				),
			),
		),
	))

	router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			boardValidator.EnsureBoardOwnership(
				boardValidator.EnsureColumnOwnership(
					createCard.Handle,
				),
			),
		),
	))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			boardValidator.EnsureBoardOwnership(
				boardValidator.EnsureColumnOwnership(
					listAllCards.Handle,
				),
			),
		),
	))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			boardValidator.EnsureBoardOwnership(
				boardValidator.EnsureColumnOwnership(
					boardValidator.EnsureCardOwnership(
						listCard.Handle,
					),
				),
			),
		),
	))
	router.HandlerFunc(http.MethodPut, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			boardValidator.EnsureBoardOwnership(
				boardValidator.EnsureColumnOwnership(
					boardValidator.EnsureCardOwnership(
						editCard.Handle,
					),
				),
			),
		),
	))
	router.HandlerFunc(http.MethodDelete, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			boardValidator.EnsureBoardOwnership(
				boardValidator.EnsureColumnOwnership(
					boardValidator.EnsureCardOwnership(
						deleteCard.Handle,
					),
				),
			),
		),
	))
	router.HandlerFunc(http.MethodPut, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId/move", userAuthenticator.Authenticate(
		teamMemberValidator.EnsureMemberAccess(
			boardValidator.EnsureBoardOwnership(
				boardValidator.EnsureColumnOwnership(
					boardValidator.EnsureCardOwnership(
						moveCardtoAnotherColumn.Handle,
					),
				),
			),
		),
	))

	router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/polls", userAuthenticator.Authenticate(createPoll.Handle))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/polls", userAuthenticator.Authenticate(listAllPolls.Handle))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/polls/:pollId", userAuthenticator.Authenticate(listPoll.Handle))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/polls/:pollId/result", userAuthenticator.Authenticate(showPollResult.Handle))

	router.HandlerFunc(http.MethodDelete, "/v1/teams/:teamId/polls/:pollId/options/:optionId", userAuthenticator.Authenticate(deleteOption.Handle))

	router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/polls/:pollId/options/:optionId/vote", userAuthenticator.Authenticate(vote.Handle))

	router.Handler(http.MethodGet, "/v1/docs/*filepath", httpSwagger.WrapHandler)

	return middleware.RecoverPanic(middleware.EnableCORS(router))
}
