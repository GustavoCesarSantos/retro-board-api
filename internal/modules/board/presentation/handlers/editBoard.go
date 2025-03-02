package board

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type EditBoard struct {
	updateBoard application.IUpdateBoard
}

func NewEditBoard(
	updateBoard application.IUpdateBoard,
) *EditBoard {
    return &EditBoard{
		updateBoard,
    }
}

// @Summary      Edit a board
// @Description  Updates the name or activation status of a board.
// @Tags         Board
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId   path      int                     true  "Team ID"
// @Param        boardId  path      int                     true  "Board ID"
// @Param        body     body      dtos.EditBoardRequest   true  "Board update details"
// @Success      204      "Board successfully updated"
// @Failure      400      {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      404      {object} utils.ErrorEnvelope          "Team or board not found"
// @Failure      500      {object} utils.ErrorEnvelope          "Internal server error"
// @Router       /teams/:teamId/boards/:boardId [put]
func(eb *EditBoard) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "editBoard.go",
		"func": "editBoard.Handle",
		"line": 0,
	}
	boardId, boardIdErr := utils.ReadIDParam(r, "boardId")
	if boardIdErr != nil {
		utils.NotFoundResponse(w, r, metadataErr)
		return
	}
	var input dtos.EditBoardRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr, metadataErr)
		return
	}
    updateErr := eb.updateBoard.Execute(boardId, struct {
		Name *string
		Active *bool
	}{ 
		Name: input.Name,
		Active: input.Active,
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
