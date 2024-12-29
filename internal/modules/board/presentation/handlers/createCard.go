package board

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type createCard struct {
	ensureBoardOwnership application.IEnsureBoardOwnership
    ensureColumnOwnership application.IEnsureColumnOwnership
    saveCard application.ISaveCard
}

func NewCreateCard(
	ensureBoardOwnership application.IEnsureBoardOwnership,
    ensureColumnOwnership application.IEnsureColumnOwnership,
    saveCard application.ISaveCard,
) *createCard {
    return &createCard{
		ensureBoardOwnership,
        ensureColumnOwnership,
        saveCard,
    }
}

type CreateCardEnvelop struct {
	Card dtos.CreateCardResponse `json:"card"`
}

// @Summary      Create a new card
// @Description  Creates a card associated with the specified column, board, and team.
// @Tags         Board
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId    path      int    true  "Team ID"
// @Param        boardId   path      int    true  "Board ID"
// @Param        columnId  path      int    true  "Column ID"
// @Param        input     body      dtos.CreateCardRequest true "Card creation data"
// @Success 201 {object}   board.CreateCardEnvelop "Card successfully created"
// @Failure      400       {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      404       {object} utils.ErrorEnvelope "Team, board, or column not found"
// @Failure      500       {object} utils.ErrorEnvelope "Internal server error"
// @Router       /teams/:teamId/boards/:boardId/columns/:columnId/cards [post]
func(cc *createCard) Handle(w http.ResponseWriter, r *http.Request) {
	//TO-DO criar caso de uso para verificar se o usuário pertence ao time
	user := utils.ContextGetUser(r)
    teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
	if teamIdErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
    boardId, boardIdErr := utils.ReadIDParam(r, "boardId")
	if boardIdErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	columnId, columnIdErr := utils.ReadIDParam(r, "columnId")
	if columnIdErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	var input dtos.CreateCardRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr)
		return
	}
    ensureBoardErr := cc.ensureBoardOwnership.Execute(teamId, boardId)
	if ensureBoardErr != nil {
		utils.BadRequestResponse(w, r, ensureBoardErr)
		return
	}
    ensureColumnErr := cc.ensureColumnOwnership.Execute(boardId, columnId)
	if ensureColumnErr != nil {
		utils.BadRequestResponse(w, r, ensureColumnErr)
		return
	}
    card, saveErr := cc.saveCard.Execute(columnId, user.ID, input.Text)
	if saveErr != nil {
		utils.ServerErrorResponse(w, r, saveErr)
		return
	}
	response := dtos.NewCreateCardResponse(card.ID, card.Text, card.CreatedAt)
    writeJsonErr := utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"card": response}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
