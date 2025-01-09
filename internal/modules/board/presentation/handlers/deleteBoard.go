package board

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type deleteBoard struct {
	removeBoard application.IRemoveBoard
}

func NewDeleteBoard(
	removeBoard application.IRemoveBoard,
) *deleteBoard {
    return &deleteBoard{
	    removeBoard,
    }
}

// @Summary      Delete a board
// @Description  Deletes a board associated with the specified team and board ID.
// @Tags         Board
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId    path      int    true  "Team ID"
// @Param        boardId   path      int    true  "Board ID"
// @Success      204       "Board successfully deleted"
// @Failure      400       {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      404       {object} utils.ErrorEnvelope "Team or board not found"
// @Failure      500       {object} utils.ErrorEnvelope "Internal server error"
// @Router       /teams/:teamId/boards/:boardId [delete]
func(db *deleteBoard) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "deleteBoard.go",
		"func": "deleteBoard.Handle",
		"line": 0,
	}
	boardId, boardIdErr := utils.ReadIDParam(r, "boardId")
	if boardIdErr != nil {
		utils.NotFoundResponse(w, r, metadataErr)
		return
	}
    removeErr := db.removeBoard.Execute(boardId)
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
