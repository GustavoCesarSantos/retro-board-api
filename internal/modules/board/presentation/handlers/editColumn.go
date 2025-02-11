package board

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type editColumn struct {
	updateColumn application.IUpdateColumn
}

func NewEditColumn(
	updateColumn application.IUpdateColumn,
) *editColumn {
    return &editColumn{
		updateColumn,
    }
}

// @Summary      Edit a column
// @Description  Updates the details of a column, such as name and color.
// @Tags         Board
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId     path      int                        true  "Team ID"
// @Param        boardId    path      int                        true  "Board ID"
// @Param        columnId   path      int                        true  "Column ID"
// @Param        body       body      dtos.EditColumnRequest     true  "Column update details"
// @Success      204        "Column successfully updated"
// @Failure      400        {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      404        {object} utils.ErrorEnvelope           "Team, board, or column not found"
// @Failure      500        {object} utils.ErrorEnvelope           "Internal server error"
// @Router       /teams/:teamId/boards/:boardId/columns/:columnId [put]
func(ec *editColumn) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "editColumn.go",
		"func": "editColumn.Handle",
		"line": 0,
	}
	columnId, columnIdErr := utils.ReadIDParam(r, "columnId")
	if columnIdErr != nil {
		utils.NotFoundResponse(w, r, metadataErr)
		return
	}
	var input dtos.EditColumnRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr, metadataErr)
		return
	}
    updateErr := ec.updateColumn.Execute(columnId, struct {
		Name *string
		Color *string
        Position *int
	}{ 
		Name: input.Name,
		Color: input.Color,
            Position: input.Position, 
	})
	if updateErr != nil {
		switch {
		case errors.Is(updateErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r, metadataErr)
		default:
            utils.ServerErrorResponse(w, r, updateErr, metadataErr)
		}
		return
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
