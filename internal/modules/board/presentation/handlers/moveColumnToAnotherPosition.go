package board

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type MoveColumnToAnotherPosition struct {
	moveColumn application.IMoveColumn
}

func NewMoveColumnToAnotherPosition(
	moveColumn application.IMoveColumn,
) *MoveColumnToAnotherPosition {
    return &MoveColumnToAnotherPosition{
        moveColumn,
    }
}

func(mc *MoveColumnToAnotherPosition) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "moveColumnToAnotherPosition.go",
		"func": "moveColumnToAnotherPosition.Handle",
		"line": 0,
	}
	columnId, columnIdErr := utils.ReadIDParam(r, "columnId")
	if columnIdErr != nil {
		utils.BadRequestResponse(w, r, columnIdErr, metadataErr)
		return
	}
	var input dtos.MoveColumnToAnotherPositionRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr, metadataErr)
		return
	}
    moveErr := mc.moveColumn.Execute(columnId, input.NewPosition)
	if moveErr != nil {
		switch {
		case errors.Is(moveErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r, metadataErr)
		default:
            utils.ServerErrorResponse(w, r, moveErr, metadataErr)
		}
		return
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
