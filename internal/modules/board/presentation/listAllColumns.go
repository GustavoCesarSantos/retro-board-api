package board

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type listAllColumns struct {
	ensureBoardOwnership application.IEnsureBoardOwnership
    findAllColumns application.IFindAllColumns
}

func NewListAllColumns(
	ensureBoardOwnership application.IEnsureBoardOwnership, 
	findAllColumns application.IFindAllColumns,
) *listAllColumns {
    return &listAllColumns {
		ensureBoardOwnership,
        findAllColumns,
    }
}

func(lc *listAllColumns) Handle(w http.ResponseWriter, r *http.Request) {
	//TO-DO criar caso de uso para verificar se o usuário pertence ao time
    teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
	if teamIdErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	boardId, boardIdErr := utils.ReadIDParam(r, "boardId")
	if boardIdErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	ensureBoardErr := lc.ensureBoardOwnership.Execute(teamId, boardId)
	if ensureBoardErr != nil {
		utils.BadRequestResponse(w, r, ensureBoardErr)
		return
	}
	columns, findErr := lc.findAllColumns.Execute(boardId)
	if findErr != nil {
		utils.ServerErrorResponse(w, r, findErr)
		return
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"columns": columns}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
