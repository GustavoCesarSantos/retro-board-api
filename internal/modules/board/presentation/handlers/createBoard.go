package board

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type createBoard struct {
    saveBoard application.ISaveBoard
}

func NewCreateBoard(saveBoard application.ISaveBoard) *createBoard {
    return &createBoard{
        saveBoard,
    }
}

// @Summary      Create a new board
// @Description  Creates a board associated with the specified team.
// @Tags         Board
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId   path      int    true  "Team ID"
// @Param        input    body      dtos.CreateBoardRequest true "Board creation data"
// @Success      204      "Board successfully created"
// @Failure      400      {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      404      {object} utils.ErrorEnvelope "Team not found"
// @Failure      500      {object} utils.ErrorEnvelope "Internal server error"
// @Router      /teams/:teamId/boards [post]
func(cb *createBoard) Handle(w http.ResponseWriter, r *http.Request) {
    teamId, err := utils.ReadIDParam(r, "teamId")
	if err != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	var input dtos.CreateBoardRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr)
		return
	}
    saveErr := cb.saveBoard.Execute(teamId, input.Name)
	if saveErr != nil {
		utils.ServerErrorResponse(w, r, saveErr)
		return
	}
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}