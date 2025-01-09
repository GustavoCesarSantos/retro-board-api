package board

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type deleteColumn struct {
	removeColumn application.IRemoveColumn
}

func NewDeleteColumn(
	removeColumn application.IRemoveColumn,
) *deleteColumn {
    return &deleteColumn{
		removeColumn,
    }
}

// @Summary      Delete a column
// @Description  Deletes a column associated with the specified team, board, and column ID.
// @Tags         Board
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId    path      int    true  "Team ID"
// @Param        boardId   path      int    true  "Board ID"
// @Param        columnId  path      int    true  "Column ID"
// @Success      204       "Column successfully deleted"
// @Failure      400       {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      404       {object} utils.ErrorEnvelope "Team, board, or column not found"
// @Failure      500       {object} utils.ErrorEnvelope "Internal server error"
// @Router       /teams/:teamId/boards/:boardId/columns/:columnId [delete]
func(dc *deleteColumn) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "deleteColumn.go",
		"func": "deleteColumn.Handle",
		"line": 0,
	}
	columnId, columnIdErr := utils.ReadIDParam(r, "columnId")
	if columnIdErr != nil {
		utils.NotFoundResponse(w, r, metadataErr)
		return
	}
    removeErr := dc.removeColumn.Execute(columnId)
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
