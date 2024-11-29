package board

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
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

func(ec *editColumn) Handle(w http.ResponseWriter, r *http.Request) {
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
	var input struct {
		Name   *string       `json:"name"`
		Color *string       `json:"color"`
	}
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
    ec.updateColumn.Execute(boardId, struct {
		Name *string
		Color *string
	}{ 
		Name: input.Name,
		Color: input.Color,
	})
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
