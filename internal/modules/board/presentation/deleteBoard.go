package board

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type deleteBoard struct {
	ensureBoardOwnership application.IEnsureBoardOwnership
    removeBoard application.IRemoveBoard
}

func NewDeleteBoard(ensureBoardOwnership application.IEnsureBoardOwnership, removeBoard application.IRemoveBoard) *deleteBoard {
    return &deleteBoard{
		ensureBoardOwnership,
        removeBoard,
    }
}

func(db *deleteBoard) Handle(w http.ResponseWriter, r *http.Request) {
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
	ensureBoardErr := db.ensureBoardOwnership.Execute(teamId, boardId)
	if ensureBoardErr != nil {
		utils.BadRequestResponse(w, r, ensureBoardErr)
		return
	}
    removeErr := db.removeBoard.Execute(boardId)
	if removeErr != nil {
		switch {
		case errors.Is(removeErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r)
		default:
            utils.ServerErrorResponse(w, r, removeErr)
		}
		return
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
