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

func NewCreateBoard(
	saveBoard application.ISaveBoard,
) *createBoard {
    return &createBoard{
		saveBoard,
    }
}

type CreateBoardEnvelop struct {
	Board dtos.CreateBoardResponse `json:"board"`
}

// @Summary      Create a new board
// @Description  Creates a board associated with the specified team.
// @Tags         Board
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId   path      int    true  "Team ID"
// @Param        input    body      dtos.CreateBoardRequest true "Board creation data"
// @Success 201 {object}  board.CreateBoardEnvelop "Board successfully created"
// @Failure      400      {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      404      {object} utils.ErrorEnvelope "Team not found"
// @Failure      500      {object} utils.ErrorEnvelope "Internal server error"
// @Router      /teams/:teamId/boards [post]
func(cb *createBoard) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "createBoard.go",
		"func": "createBoard.Handle",
		"line": 0,
	}
	teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
	if teamIdErr != nil {
		utils.BadRequestResponse(w, r, teamIdErr, metadataErr)
		return
	}
	var input dtos.CreateBoardRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr, metadataErr)
		return
	}
	board, saveErr := cb.saveBoard.Execute(teamId, input.Name)
	if saveErr != nil {
		utils.ServerErrorResponse(w, r, saveErr, metadataErr)
		return
	}
	response := dtos.NewCreateBoardResponse(board.ID, board.Name, board.CreatedAt)
    writeJsonErr := utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"board": response}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
