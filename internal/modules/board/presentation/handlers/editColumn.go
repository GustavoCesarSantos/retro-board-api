package board

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type editColumn struct {
	ensureBoardOwnership application.IEnsureBoardOwnership
	ensureColumnOwnership application.IEnsureColumnOwnership
    updateColumn application.IUpdateColumn
}

func NewEditColumn(
	ensureBoardOwnership application.IEnsureBoardOwnership,
	ensureColumnOwnership application.IEnsureColumnOwnership, 
	updateColumn application.IUpdateColumn,
) *editColumn {
    return &editColumn{
		ensureBoardOwnership,
		ensureColumnOwnership,
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
	//TO-DO criar caso de uso para verificar se o usuário pertence ao time
    teamId, err := utils.ReadIDParam(r, "teamId")
	if err != nil {
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
	var input dtos.EditColumnRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr)
		return
	}
	ensureBoardErr := ec.ensureBoardOwnership.Execute(teamId, boardId)
	if ensureBoardErr != nil {
		utils.BadRequestResponse(w, r, ensureBoardErr)
		return
	}
    ensureColumnErr := ec.ensureColumnOwnership.Execute(boardId, columnId)
	if ensureColumnErr != nil {
		utils.BadRequestResponse(w, r, ensureColumnErr)
		return
	}
    updateErr := ec.updateColumn.Execute(columnId, struct {
		Name *string
		Color *string
	}{ 
		Name: input.Name,
		Color: input.Color,
	})
	if updateErr != nil {
		switch {
		case errors.Is(updateErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r)
		default:
            utils.ServerErrorResponse(w, r, updateErr)
		}
		return
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
