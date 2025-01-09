package board

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type deleteCard struct {
	removeCard application.IRemoveCard
}

func NewDeleteCard(
	removeCard application.IRemoveCard,
) *deleteCard {
    return &deleteCard{
		removeCard,
    }
}

// @Summary      Delete a card
// @Description  Deletes a card associated with the specified team, board, column, and card ID.
// @Tags         Board
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId    path      int    true  "Team ID"
// @Param        boardId   path      int    true  "Board ID"
// @Param        columnId  path      int    true  "Column ID"
// @Param        cardId    path      int    true  "Card ID"
// @Success      204       "Card successfully deleted"
// @Failure      400       {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      404       {object} utils.ErrorEnvelope "Team, board, column, or card not found"
// @Failure      500       {object} utils.ErrorEnvelope "Internal server error"
// @Router       /teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId [delete]
func(dc *deleteCard) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "deleteCard.go",
		"func": "deleteCard.Handle",
		"line": 0,
	}
	cardId, cardIdErr := utils.ReadIDParam(r, "cardId")
	if cardIdErr != nil {
		utils.NotFoundResponse(w, r, metadataErr)
		return
	}
    removeErr := dc.removeCard.Execute(cardId)
    if removeErr != nil {
		switch {
		case errors.Is(removeErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r, metadataErr)
		default:
            utils.ServerErrorResponse(w, r, removeErr, metadataErr)
		}
		return
    }
	writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
