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
	pollApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/application"
	pollDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/external/db/memory"
	poll "github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/presentation"
	teamApplication "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	teamDb "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/external/db/memory"
	team "github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/presentation"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

var boardRepository = boardDb.NewBoardRepository()
var cardRepository = boardDb.NewCardRepository()
var columnRepository = boardDb.NewColumnRepository()
var optionRepository = pollDb.NewOptionRepository()
var pollRepository = pollDb.NewPollRepository()
var teamRepository = teamDb.NewTeamRepository()
var teamMemberRepository = teamDb.NewTeamMemberRepository()
var userRepository = userDb.NewUserRepository()
var voteRepository = pollDb.NewVoteRepository()

var countVotesByPollId = pollApplication.NewCountVotesByPollId(
    optionRepository, 
    voteRepository,
)
var createAuthToken = identityApplication.NewCreateAuthToken()
var ensureBoardOwnership = boardApplication.NewEnsureBoardOwnership(boardRepository)
var ensureCardOwnership = boardApplication.NewEnsureCardOwnership(cardRepository)
var ensureColumnOwnership = boardApplication.NewEnsureColumnOwnership(
    columnRepository,
)
var ensureOptionOwnership = pollApplication.NewEnsureOptionOwnership(optionRepository)
var ensurePollOwnership = pollApplication.NewEnsurePollOwnership(pollRepository)
var findAllBoards = boardApplication.NewFindAllBoards(boardRepository)
var findAllCards = boardApplication.NewFindAllCards(cardRepository)
var findAllColumns = boardApplication.NewFindAllColumns(columnRepository)
var findAllPolls = pollApplication.NewFindAllPolls(pollRepository)
var findAllTeams = teamApplication.NewFindAllTeams(teamRepository)
var findBoard = boardApplication.NewFindBoard(boardRepository)
var findCard = boardApplication.NewFindCard(cardRepository)
var findPoll = pollApplication.NewFindPoll(pollRepository)
var findTeam = teamApplication.NewFindTeam(teamRepository)
var findUserByEmail = identityApplication.NewFindUserByEmail(userRepository)
var getNextColumnPosition = boardApplication.NewGetNextColumnPosition(
    columnRepository,
)
var moveCardBetweenColumns = boardApplication.NewMoveCardBetweenColumns(
    cardRepository,
)
var removeBoard = boardApplication.NewRemoveBoard(boardRepository)
var removeCard = boardApplication.NewRemoveCard(cardRepository)
var removeColumn = boardApplication.NewRemoveColumn(columnRepository)
var removeMember = teamApplication.NewRemoveMember(teamMemberRepository)
var removeOption = pollApplication.NewRemoveOption(optionRepository)
var saveBoard = boardApplication.NewSaveBoard(boardRepository)
var saveCard = boardApplication.NewSaveCard(cardRepository)
var saveColumn = boardApplication.NewSaveColumn(columnRepository)
var saveMember = teamApplication.NewSaveMember(teamMemberRepository)
var saveOption = pollApplication.NewSaveOption(optionRepository)
var savePoll = pollApplication.NewSavePoll(pollRepository)
var saveTeam = teamApplication.NewSaveTeam(teamRepository)
var saveUser = identityApplication.NewSaveUser(userRepository)
var saveVote = pollApplication.NewSaveVote(voteRepository)
var updateBoard = boardApplication.NewUpdateBoard(boardRepository)
var updateCard = boardApplication.NewUpdateCard(cardRepository)
var updateColumn = boardApplication.NewUpdateColumn(columnRepository)
var updateRole = teamApplication.NewUpdateRole(teamMemberRepository)

var addMemberToTeam = team.NewAddMemberToTeam(saveMember)
var changeMemberRole = team.NewChangeMemberRole(updateRole)
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
var createPoll = poll.NewCreatePoll(saveOption, savePoll)
var createTeam = team.NewCreateTeam(saveTeam)
var deleteBoard = board.NewDeleteBoard(ensureBoardOwnership, removeBoard)
var deleteCard = board.NewDeleteCard(
	ensureBoardOwnership,
	ensureColumnOwnership,
	ensureCardOwnership,
	removeCard,
)
var deleteColumn = board.NewDeleteColumn(
	ensureBoardOwnership,
	ensureColumnOwnership,
	removeColumn,
)
var deleteOption = poll.NewDeleteOption(
    ensurePollOwnership, 
    ensureOptionOwnership, 
    removeOption,
)
var editBoard = board.NewEditBoard(ensureBoardOwnership, updateBoard)
var editCard = board.NewEditCard(
	ensureBoardOwnership,
	ensureColumnOwnership,
	ensureCardOwnership,
	updateCard,
)
var editColumn = board.NewEditColumn(
	ensureBoardOwnership,
	ensureColumnOwnership,
	updateColumn,
)
var healthcheck = monitor.NewHealthcheck()
var listAllBoards = board.NewListAllBoards(findAllBoards)
var listAllCards = board.NewListAllCards(
	ensureBoardOwnership,
	ensureColumnOwnership,
	findAllCards,
)
var listAllColumns = board.NewListAllColumns(ensureBoardOwnership, findAllColumns)
var listAllPolls = poll.NewListAllPolls(findAllPolls)
var listAllTeams = team.NewListAllTeams(findAllTeams)
var listBoard = board.NewListBoard(ensureBoardOwnership, findBoard)
var listCard = board.NewListCard(
	ensureBoardOwnership,
	ensureColumnOwnership,
	ensureCardOwnership,
	findCard,
)
var listPoll = poll.NewListPoll(ensurePollOwnership, findPoll)
var moveCardtoAnotherColumn = board.NewMoveCardtoAnotherColumn(
	ensureBoardOwnership,
	ensureColumnOwnership,
	ensureCardOwnership,
	moveCardBetweenColumns,
)
var removeMemberFromTeam = team.NewRemoveMemberFromTeam(removeMember)
var showPollResult = poll.NewShowPollResult(ensurePollOwnership, countVotesByPollId)
var showTeam = team.NewShowTeam(findTeam)
var signinUser = identity.NewSigninUser(createAuthToken, findUserByEmail)
var signinWithGoogle = identity.NewSigninWithGoogle()
var signinWithGoogleCallback = identity.NewSigninWithGoogleCallback(
    findUserByEmail, 
    saveUser,
)
var vote = poll.NewVote(ensurePollOwnership, ensureOptionOwnership, saveVote)

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
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/boards", listAllBoards.Handle)
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/boards/:boardId", listBoard.Handle)
	router.HandlerFunc(http.MethodPut, "/v1/teams/:teamId/boards/:boardId", editBoard.Handle)
	router.HandlerFunc(http.MethodDelete, "/v1/teams/:teamId/boards/:boardId", deleteBoard.Handle)

	router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/boards/:boardId/columns", createColumn.Handle)
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/boards/:boardId/columns", listAllColumns.Handle)
	router.HandlerFunc(http.MethodPut, "/v1/teams/:teamId/boards/:boardId/columns/:columnId", editColumn.Handle)
	router.HandlerFunc(http.MethodDelete, "/v1/teams/:teamId/boards/:boardId/columns/:columnId", deleteColumn.Handle)

	router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards", createCard.Handle)
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards", listAllCards.Handle)
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId", listCard.Handle)
	router.HandlerFunc(http.MethodPut, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId", editCard.Handle)
	router.HandlerFunc(http.MethodDelete, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId", deleteCard.Handle)
	router.HandlerFunc(http.MethodPut, "/v1/teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId/move", moveCardtoAnotherColumn.Handle)

	router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/polls", createPoll.Handle)
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/polls", listAllPolls.Handle)
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/polls/:pollId", listPoll.Handle)
	router.HandlerFunc(http.MethodGet, "/v1/teams/:teamId/polls/:pollId/result", showPollResult.Handle)

	router.HandlerFunc(http.MethodDelete, "/v1/teams/:teamId/polls/:pollId/options/:optionId", deleteOption.Handle)

	router.HandlerFunc(http.MethodPost, "/v1/teams/:teamId/polls/:pollId/options/:optionId/vote", vote.Handle)

	return middleware.RecoverPanic(middleware.EnableCORS(router))
}
