package board

import (
	"math"
	"net/http"
	"strconv"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type ListAllColumns struct {
	findAllColumns application.IFindAllColumns
}

func NewListAllColumns(
	findAllColumns application.IFindAllColumns,
) *ListAllColumns {
    return &ListAllColumns {
		findAllColumns,
    }
}

type ListAllColumnsEnvelop struct {
	Columns dtos.ListAllColumnsResponsePaginated `json:"columns"`
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
func(lc *ListAllColumns) Handle(w http.ResponseWriter, r *http.Request) {
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
	limitStr := r.URL.Query().Get("limit")
	lastIDStr := r.URL.Query().Get("lastId")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		metadataErr["line"] = 47
		utils.BadRequestResponse(w, r, utils.ErrMissingOrInvalidLimitQueryParam, metadataErr)
		return
	}
	lastID := math.MaxInt64
	if lastIDStr != "" {
		lastID, err = strconv.Atoi(lastIDStr)
		if err != nil {
			metadataErr["line"] = 55
			utils.BadRequestResponse(w, r, utils.ErrInvalidLimitQueryParam, metadataErr)
			return
		}
	}
	columns, findErr := lc.findAllColumns.Execute(boardId, limit, lastID)
	if findErr != nil {
		utils.ServerErrorResponse(w, r, findErr, metadataErr)
		return
    }
	var response dtos.ListAllColumnsResponsePaginated
    for _, column := range columns.Items {
        response.Items = append(response.Items, dtos.NewListAllColumnsResponse(
			column.ID,
			column.BoardId,
			column.Name,
			column.Color,
			column.Position,
			column.CreatedAt,
			column.UpdatedAt,
		))
    }
	response.NextCursor = columns.NextCursor
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"columns": response}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
