package http

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/GustavoCesarSantos/retro-board-api/internal/infra/http/middleware"
	boardApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	boardDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/external/db/memory"
	board "github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation"
	identityApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/application"
	userDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/external/db/nativeSql"
	identity "github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/presentation"
	monitor "github.com/GustavoCesarSantos/retro-board-api/internal/modules/monitor/presentation"
	pollApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/application"
	pollDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/external/db/memory"
	poll "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/presentation"
	teamApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	teamDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/nativeSql"
	team "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/presentation"
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

	//Init Middlewares
	userAuthenticator := middleware.NewUserAuthenticator(userRepository)

	//Init Use Cases
	countVotesByPollId := pollApplication.NewCountVotesByPollId(
		optionRepository, 
		voteRepository,
	)
	createAuthToken := identityApplication.NewCreateAuthToken()
	decodeAuthToken := identityApplication.NewDecodeAuthToken()
	ensureAdminMembership := teamApplication.NewEnsureAdminMembership(teamMemberRepository)
	ensureBoardOwnership := boardApplication.NewEnsureBoardOwnership(boardRepository)
	ensureCardOwnership := boardApplication.NewEnsureCardOwnership(cardRepository)
	ensureColumnOwnership := boardApplication.NewEnsureColumnOwnership(
		columnRepository,
	)
	ensureOptionOwnership := pollApplication.NewEnsureOptionOwnership(optionRepository)
	ensurePollOwnership := pollApplication.NewEnsurePollOwnership(pollRepository)
	findAllBoards := boardApplication.NewFindAllBoards(boardRepository)
	findAllCards := boardApplication.NewFindAllCards(cardRepository)
	findAllColumns := boardApplication.NewFindAllColumns(columnRepository)
	findAllPolls := pollApplication.NewFindAllPolls(pollRepository)
	findAllTeams := teamApplication.NewFindAllTeams(teamRepository)
	findBoard := boardApplication.NewFindBoard(boardRepository)
	findCard := boardApplication.NewFindCard(cardRepository)
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
	addMemberToTeam := team.NewAddMemberToTeam(ensureAdminMembership, saveMember)
	changeMemberRole := team.NewChangeMemberRole(ensureAdminMembership, updateRole)
	createBoard := board.NewCreateBoard(saveBoard)
	createCard := board.NewCreateCard(
		ensureBoardOwnership,
		ensureColumnOwnership,
		saveCard,
	)
	createColumn := board.NewCreateColumn(
		ensureBoardOwnership,
		findAllColumns,
		getNextColumnPosition,
		saveColumn,
	)
	createPoll := poll.NewCreatePoll(saveOption, savePoll)
	createTeam := team.NewCreateTeam(removeTeam, saveMember, saveTeam)
	deleteBoard := board.NewDeleteBoard(ensureBoardOwnership, removeBoard)
	deleteCard := board.NewDeleteCard(
		ensureBoardOwnership,
		ensureColumnOwnership,
		ensureCardOwnership,
		removeCard,
	)
	deleteColumn := board.NewDeleteColumn(
		ensureBoardOwnership,
		ensureColumnOwnership,
		removeColumn,
	)
	deleteOption := poll.NewDeleteOption(
		ensurePollOwnership, 
		ensureOptionOwnership, 
		removeOption,
	)
	editBoard := board.NewEditBoard(ensureBoardOwnership, updateBoard)
	editCard := board.NewEditCard(
		ensureBoardOwnership,
		ensureColumnOwnership,
		ensureCardOwnership,
		updateCard,
	)
	editColumn := board.NewEditColumn(
		ensureBoardOwnership,
		ensureColumnOwnership,
		updateColumn,
	)
	healthcheck := monitor.NewHealthcheck()
	listAllBoards := board.NewListAllBoards(findAllBoards)
	listAllCards := board.NewListAllCards(
		ensureBoardOwnership,
		ensureColumnOwnership,
		findAllCards,
	)
	listAllColumns := board.NewListAllColumns(ensureBoardOwnership, findAllColumns)
	listAllPolls := poll.NewListAllPolls(findAllPolls)
	listAllTeams := team.NewListAllTeams(findAllTeams)
	listBoard := board.NewListBoard(ensureBoardOwnership, findBoard)
	listCard := board.NewListCard(
		ensureBoardOwnership,
		ensureColumnOwnership,
		ensureCardOwnership,
		findCard,
	)
	listPoll := poll.NewListPoll(ensurePollOwnership, findPoll)
	moveCardtoAnotherColumn := board.NewMoveCardtoAnotherColumn(
		ensureBoardOwnership,
		ensureColumnOwnership,
		ensureCardOwnership,
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
	router.HandlerFunc(http.MethodPut, "/v1/teams/:teamId/members/:memberId/change-role", userAuthenticator.Authenticate(changeMemberRole.Handle))

	router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/boards", userAuthenticator.Authenticate(createBoard.Handle))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/boards", userAuthenticator.Authenticate(listAllBoards.Handle))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/boards/:boardId", userAuthenticator.Authenticate(listBoard.Handle))
	router.HandlerFunc(http.MethodPut, "/v1/teams/:teamId/boards/:boardId", userAuthenticator.Authenticate(editBoard.Handle))
	router.HandlerFunc(http.MethodDelete, "/v1/teams/:teamId/boards/:boardId", userAuthenticator.Authenticate(deleteBoard.Handle))

	router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/boards/:boardId/columns", userAuthenticator.Authenticate(createColumn.Handle))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/boards/:boardId/columns", userAuthenticator.Authenticate(listAllColumns.Handle))
	router.HandlerFunc(http.MethodPut, "/v1/teams/:teamId/boards/:boardId/columns/:columnId", userAuthenticator.Authenticate(editColumn.Handle))
	router.HandlerFunc(http.MethodDelete, "/v1/teams/:teamId/boards/:boardId/columns/:columnId", userAuthenticator.Authenticate(deleteColumn.Handle))

	router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards", userAuthenticator.Authenticate(createCard.Handle))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards", userAuthenticator.Authenticate(listAllCards.Handle))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId", userAuthenticator.Authenticate(listCard.Handle))
	router.HandlerFunc(http.MethodPut, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId", userAuthenticator.Authenticate(editCard.Handle))
	router.HandlerFunc(http.MethodDelete, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId", userAuthenticator.Authenticate(deleteCard.Handle))
	router.HandlerFunc(http.MethodPut, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId/move", userAuthenticator.Authenticate(moveCardtoAnotherColumn.Handle))

	router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/polls", userAuthenticator.Authenticate(createPoll.Handle))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/polls", userAuthenticator.Authenticate(listAllPolls.Handle))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/polls/:pollId", userAuthenticator.Authenticate(listPoll.Handle))
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/polls/:pollId/result", userAuthenticator.Authenticate(showPollResult.Handle))

	router.HandlerFunc(http.MethodDelete, "/v1/teams/:teamId/polls/:pollId/options/:optionId", userAuthenticator.Authenticate(deleteOption.Handle))

	router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/polls/:pollId/options/:optionId/vote", userAuthenticator.Authenticate(vote.Handle))

	return middleware.RecoverPanic(middleware.EnableCORS(router))
}
