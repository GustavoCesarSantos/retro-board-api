package board

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type moveCardtoAnotherColumn struct {
	ensureBoardOwnership application.IEnsureBoardOwnership
	ensureColumnOwnership application.IEnsureColumnOwnership
	ensureCardOwnership application.IEnsureCardOwnership
    moveCardBetweenColumn application.IMoveCardBetweenColumns
}

func NewMoveCardtoAnotherColumn(
	ensureBoardOwnership application.IEnsureBoardOwnership,
	ensureColumnOwnership application.IEnsureColumnOwnership,
	ensureCardOwnership application.IEnsureCardOwnership,
	moveCardBetweenColumn application.IMoveCardBetweenColumns,
) *moveCardtoAnotherColumn {
    return &moveCardtoAnotherColumn{
		ensureBoardOwnership,
		ensureColumnOwnership,
		ensureCardOwnership,
        moveCardBetweenColumn,
    }
}

func(mc *moveCardtoAnotherColumn) Handle(w http.ResponseWriter, r *http.Request) {
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
	cardId, cardIdErr := utils.ReadIDParam(r, "cardId")
	if cardIdErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	var input struct {
		NewColumnId   *int64       `json:"newColumnId"`
	}
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr)
		return
	}
	ensureBoardErr := mc.ensureBoardOwnership.Execute(teamId, boardId)
	if ensureBoardErr != nil {
		utils.BadRequestResponse(w, r, ensureBoardErr)
		return
	}
    ensureColumnErr := mc.ensureColumnOwnership.Execute(boardId, columnId)
	if ensureColumnErr != nil {
		utils.BadRequestResponse(w, r, ensureColumnErr)
		return
	}
    ensureCardErr := mc.ensureCardOwnership.Execute(columnId, cardId)
	if ensureCardErr != nil {
		utils.BadRequestResponse(w, r, ensureCardErr)
		return
	}
    mc.moveCardBetweenColumn.Execute(cardId, *input.NewColumnId)
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
