package board

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
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
	var input struct {
		Name   string       `json:"name"`
		Color   string       `json:"color"`
	}
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
