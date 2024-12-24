package board

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type createColumn struct {
    ensureBoardOwnership application.IEnsureBoardOwnership
    findAllColumns application.IFindAllColumns
	getNextColumnPosition application.IGetNextColumnPosition
    saveColumn application.ISaveColumn
}

func NewCreateColumn(
    ensureBoardOwnership application.IEnsureBoardOwnership, 
    findAllColumns application.IFindAllColumns,
	getNextColumnPosition application.IGetNextColumnPosition,
    saveColumn application.ISaveColumn,
) *createColumn {
    return &createColumn{
        ensureBoardOwnership,
        findAllColumns,
		getNextColumnPosition,
        saveColumn,
    }
}

// @Summary      Create a new column
// @Description  Creates a column associated with the specified board and team.
// @Tags         Board
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId    path      int    true  "Team ID"
// @Param        boardId   path      int    true  "Board ID"
// @Param        input     body      dtos.CreateColumnRequest true "Column creation data"
// @Success      204       "Column successfully created"
// @Failure      400       {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      404       {object} utils.ErrorEnvelope "Team or board not found"
// @Failure      500       {object} utils.ErrorEnvelope "Internal server error"
// @Router      /teams/:teamId/boards/:boardId/columns [post]
func(cc *createColumn) Handle(w http.ResponseWriter, r *http.Request) {
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
	var input dtos.CreateColumnRequest
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
    position, getErr := cc.getNextColumnPosition.Execute(boardId)
	if getErr != nil {
		utils.ServerErrorResponse(w, r, getErr)
		return
	}
    saveErr := cc.saveColumn.Execute(boardId, input.Name, input.Color, position)
	if saveErr != nil {
		utils.ServerErrorResponse(w, r, saveErr)
		return
	}
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
