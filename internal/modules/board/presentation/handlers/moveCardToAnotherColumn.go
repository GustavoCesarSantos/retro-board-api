package board

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type moveCardtoAnotherColumn struct {
	moveCardBetweenColumn application.IMoveCardBetweenColumns
}

func NewMoveCardtoAnotherColumn(
	moveCardBetweenColumn application.IMoveCardBetweenColumns,
) *moveCardtoAnotherColumn {
    return &moveCardtoAnotherColumn{
	    moveCardBetweenColumn,
    }
}

// @Summary      Move a card to another column
// @Description  Moves a card from one column to another in a specific board.
// @Tags         Board
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId     path      int               true  "Team ID"
// @Param        boardId    path      int               true  "Board ID"
// @Param        columnId   path      int               true  "Column ID"
// @Param        cardId     path      int               true  "Card ID"
// @Param        body       body      dtos.MoveCardtoAnotherColumnRequest true "New Column ID"
// @Success      204        "Card moved successfully"
// @Failure      400        {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or invalid input data)"
// @Failure      404        {object} utils.ErrorEnvelope "Card or column not found"
// @Failure      500        {object} utils.ErrorEnvelope "Internal server error"
// @Router       /teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId/move [put]
func(mc *moveCardtoAnotherColumn) Handle(w http.ResponseWriter, r *http.Request) {
	cardId, cardIdErr := utils.ReadIDParam(r, "cardId")
	if cardIdErr != nil {
		utils.BadRequestResponse(w, r, cardIdErr)
		return
	}
	var input dtos.MoveCardtoAnotherColumnRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr)
		return
	}
    moveErr := mc.moveCardBetweenColumn.Execute(cardId, input.NewColumnId)
	if moveErr != nil {
		switch {
		case errors.Is(moveErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r)
		default:
            utils.ServerErrorResponse(w, r, moveErr)
		}
		return
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
