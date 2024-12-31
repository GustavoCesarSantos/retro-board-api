package board

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type listBoard struct {
	findBoard application.IFindBoard
}

func NewListBoard(
	findBoard application.IFindBoard,
) *listBoard {
    return &listBoard {
	    findBoard,
    }
}

type ListBoardEnvelop struct {
	Board *dtos.ListBoardResponse `json:"board"`
}

// @Summary      Get a single board by ID
// @Description  Retrieves the details of a specific board based on its ID.
// @Tags         Board
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId     path      int               true  "Team ID"
// @Param        boardId    path      int               true  "Board ID"
// @Success      200        {object} board.ListBoardEnvelop "Board details"
// @Failure      400        {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or invalid input data)"
// @Failure      404        {object} utils.ErrorEnvelope "Board not found"
// @Failure      500        {object} utils.ErrorEnvelope "Internal server error"
// @Router       /teams/:teamId/boards/:boardId [get]

func(lb *listBoard) Handle(w http.ResponseWriter, r *http.Request) {
	boardId, boardIdErr := utils.ReadIDParam(r, "boardId")
	if boardIdErr != nil {
		utils.BadRequestResponse(w, r, boardIdErr)
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
	response := dtos.NewListBoardResponse(
		board.ID,
		board.TeamId,
		board.Name,
		board.Active,
		board.CreatedAt,
		board.UpdatedAt,
	)
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"board": response}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
