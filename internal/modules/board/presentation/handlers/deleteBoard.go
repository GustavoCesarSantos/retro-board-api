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
