package board

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type createColumn struct {
    findAllColumns application.IFindAllColumns
	getNextColumnPosition application.IGetNextColumnPosition
    saveColumn application.ISaveColumn
}

func NewCreateColumn(
    findAllColumns application.IFindAllColumns,
	getNextColumnPosition application.IGetNextColumnPosition,
    saveColumn application.ISaveColumn,
) *createColumn {
    return &createColumn{
        findAllColumns,
		getNextColumnPosition,
        saveColumn,
    }
}

type CreateColumnEnvelop struct {
	Column dtos.CreateColumnResponse `json:"column"`
}

// @Summary      Create a new column
// @Description  Creates a column associated with the specified board and team.
// @Tags         Board
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId    path      int    true  "Team ID"
// @Param        boardId   path      int    true  "Board ID"
// @Param        input     body      dtos.CreateColumnRequest true "Column creation data"
// @Success 201 {object}   board.CreateColumnEnvelop "Column successfully created"
// @Failure      400       {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      404       {object} utils.ErrorEnvelope "Team or board not found"
// @Failure      500       {object} utils.ErrorEnvelope "Internal server error"
// @Router      /teams/:teamId/boards/:boardId/columns [post]
func(cc *createColumn) Handle(w http.ResponseWriter, r *http.Request) {
    metadataErr := utils.Envelope{
		"file": "createColumn.go",
		"func": "createColumn.Handle",
		"line": 0,
	}
	boardId, boardIdErr := utils.ReadIDParam(r, "boardId")
	if boardIdErr != nil {
		utils.NotFoundResponse(w, r, metadataErr)
		return
	}
	var input dtos.CreateColumnRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr, metadataErr)
		return
	}
    position, getErr := cc.getNextColumnPosition.Execute(boardId)
	if getErr != nil {
		utils.ServerErrorResponse(w, r, getErr, metadataErr)
		return
	}
    column, saveErr := cc.saveColumn.Execute(boardId, input.Name, input.Color, position)
	if saveErr != nil {
		utils.ServerErrorResponse(w, r, saveErr, metadataErr)
		return
	}
	response := dtos.NewCreateColumnResponse(column.ID, column.Name, column.Color, column.CreatedAt)
    writeJsonErr := utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"column": response}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
