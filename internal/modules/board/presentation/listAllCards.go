package board

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type listAllCards struct {
	ensureBoardOwnership application.IEnsureBoardOwnership
	ensureColumnOwnership application.IEnsureColumnOwnership
    findAllCards application.IFindAllCards
}

func NewListAllCards(
	ensureBoardOwnership application.IEnsureBoardOwnership,
	ensureColumnOwnership application.IEnsureColumnOwnership,
	findAllCards application.IFindAllCards,
) *listAllCards {
    return &listAllCards {
		ensureBoardOwnership,
		ensureColumnOwnership,
        findAllCards,
    }
}

func(lc *listAllCards) Handle(w http.ResponseWriter, r *http.Request) {
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
	columnId, columnIdErr := utils.ReadIDParam(r, "columnId")
	if columnIdErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	ensureBoardErr := lc.ensureBoardOwnership.Execute(teamId, boardId)
	if ensureBoardErr != nil {
		utils.BadRequestResponse(w, r, ensureBoardErr)
		return
	}
    ensureColumnErr := lc.ensureColumnOwnership.Execute(boardId, columnId)
	if ensureColumnErr != nil {
		utils.BadRequestResponse(w, r, ensureColumnErr)
		return
	}
	cards := lc.findAllCards.Execute(columnId)
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"cards": cards}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
