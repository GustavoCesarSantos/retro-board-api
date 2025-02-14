package board

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type listColumn struct {
	findColumn application.IFindColumn
}

func NewListColumn(
	findColumn application.IFindColumn,
) *listColumn {
    return &listColumn {
		findColumn,
    }
}

type ListColumnEnvelop struct {
	Column *dtos.ListColumnResponse `json:"column"`
}

func(lc *listColumn) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "listColumn.go",
		"func": "listColumn.Handle",
		"line": 0,
	}
	columnId, columnIdErr := utils.ReadIDParam(r, "columnId")
	if columnIdErr != nil {
		utils.BadRequestResponse(w, r, columnIdErr, metadataErr)
		return
	}
	column, findErr := lc.findColumn.Execute(columnId)
	if findErr != nil {
		switch {
		case errors.Is(findErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r, metadataErr)
		default:
            utils.ServerErrorResponse(w, r, findErr, metadataErr)
		}
		return
    }
	response := dtos.NewListColumnResponse(
		column.ID,
		column.BoardId,
		column.Name,
		column.Color,
		column.Position,
		column.CreatedAt,
		column.UpdatedAt,
	)
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"column": response}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
