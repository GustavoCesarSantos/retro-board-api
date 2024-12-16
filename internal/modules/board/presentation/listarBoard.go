package board

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type listBoard struct {
	ensureBoardOwnership application.IEnsureBoardOwnership
    findBoard application.IFindBoard
}

func NewListBoard(ensureBoardOwnership application.IEnsureBoardOwnership, findBoard application.IFindBoard) *listBoard {
    return &listBoard {
		ensureBoardOwnership,
        findBoard,
    }
}

func(lb *listBoard) Handle(w http.ResponseWriter, r *http.Request) {
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
	ensureBoardErr := lb.ensureBoardOwnership.Execute(teamId, boardId)
	if ensureBoardErr != nil {
		utils.BadRequestResponse(w, r, ensureBoardErr)
		return
	}
	board, findErr := lb.findBoard.Execute(boardId)
	if findErr != nil {
		switch {
		case errors.Is(findErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r)
		default:
            utils.ServerErrorResponse(w, r, findErr)
		}
		return
    }
	if board == nil {
		utils.NotFoundResponse(w, r)
		return
	}
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"board": board}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
