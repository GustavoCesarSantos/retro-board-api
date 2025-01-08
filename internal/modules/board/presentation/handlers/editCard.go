package board

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type editCard struct {
    notifyUpdateCard application.INotifyUpdateCard
	updateCard application.IUpdateCard
}

func NewEditCard(
    notifyUpdateCard application.INotifyUpdateCard,
	updateCard application.IUpdateCard,
) *editCard {
    return &editCard{
        notifyUpdateCard,
	    updateCard,
    }
}

// @Summary      Edit a card
// @Description  Updates the text content of a card.
// @Tags         Board
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId    path      int                      true  "Team ID"
// @Param        boardId   path      int                      true  "Board ID"
// @Param        columnId  path      int                      true  "Column ID"
// @Param        cardId    path      int                      true  "Card ID"
// @Param        body      body      dtos.EditCardRequest     true  "Card update details"
// @Success      204       "Card successfully updated"
// @Failure      400       {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      404       {object} utils.ErrorEnvelope           "Team, board, column, or card not found"
// @Failure      500       {object} utils.ErrorEnvelope           "Internal server error"
// @Router       /teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId [put]
func(ec *editCard) Handle(w http.ResponseWriter, r *http.Request) {
	boardId, boardIdErr := utils.ReadIDParam(r, "boardId")
	if boardIdErr != nil {
		utils.BadRequestResponse(w, r, boardIdErr)
		return
	}
	columnId, columnIdErr := utils.ReadIDParam(r, "columnId")
	if columnIdErr != nil {
		utils.BadRequestResponse(w, r, columnIdErr)
		return
	}
	cardId, cardIdErr := utils.ReadIDParam(r, "cardId")
	if cardIdErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	var input dtos.EditCardRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr)
		return
	}
    updateErr := ec.updateCard.Execute(cardId, input.Text)
	if updateErr != nil {
		switch {
		case errors.Is(updateErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r)
		default:
            utils.ServerErrorResponse(w, r, updateErr)
		}
		return
    }
    ec.notifyUpdateCard.Execute(boardId, columnId, *input.Text)
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
