package board

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type listAllColumns struct {
	findAllColumns application.IFindAllColumns
}

func NewListAllColumns(
	findAllColumns application.IFindAllColumns,
) *listAllColumns {
    return &listAllColumns {
		findAllColumns,
    }
}

type ListAllColumnsEnvelop struct {
	Columns []*dtos.ListAllColumnsResponse `json:"columns"`
}

// @Summary      List all columns for a board
// @Description  Retrieves all columns associated with a specific board.
// @Tags         Board
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId     path      int               true  "Team ID"
// @Param        boardId    path      int               true  "Board ID"
// @Success      200        {object} board.ListAllColumnsEnvelop "List of columns"
// @Failure      400        {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or invalid input data)"
// @Failure      500        {object} utils.ErrorEnvelope "Internal server error"
// @Router       /teams/:teamId/boards/:boardId/columns [get]
func(lc *listAllColumns) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "listAllColumns.go",
		"func": "listAllColumns.Handle",
		"line": 0,
	}
	boardId, boardIdErr := utils.ReadIDParam(r, "boardId")
	if boardIdErr != nil {
		utils.BadRequestResponse(w, r, boardIdErr, metadataErr)
		return
	}
	columns, findErr := lc.findAllColumns.Execute(boardId)
	if findErr != nil {
		utils.ServerErrorResponse(w, r, findErr, metadataErr)
		return
    }
	var response []*dtos.ListAllColumnsResponse
    for _, column := range columns {
        response = append(response, dtos.NewListAllColumnsResponse(
			column.ID,
			column.BoardId,
			column.Name,
			column.Color,
			column.Position,
			column.CreatedAt,
			column.UpdatedAt,
		))
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"columns": response}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
