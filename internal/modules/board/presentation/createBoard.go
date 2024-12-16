package board

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
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

func(cb *createBoard) Handle(w http.ResponseWriter, r *http.Request) {
    teamId, err := utils.ReadIDParam(r, "teamId")
	if err != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	var input struct {
		Name   string       `json:"name"`
	}
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
