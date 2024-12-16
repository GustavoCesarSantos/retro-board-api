package board

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type listCard struct {
	ensureBoardOwnership application.IEnsureBoardOwnership
	ensureColumnOwnership application.IEnsureColumnOwnership
	ensureCardOwnership application.IEnsureCardOwnership
    findCard application.IFindCard
}

func NewListCard(
	ensureBoardOwnership application.IEnsureBoardOwnership,
	ensureColumnOwnership application.IEnsureColumnOwnership,
	ensureCardOwnership application.IEnsureCardOwnership, 
	findCard application.IFindCard,
) *listCard {
    return &listCard {
		ensureBoardOwnership,
		ensureColumnOwnership,
		ensureCardOwnership,
        findCard,
    }
}

func(lc *listCard) Handle(w http.ResponseWriter, r *http.Request) {
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
	cardId, cardIdErr := utils.ReadIDParam(r, "cardId")
	if cardIdErr != nil {
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
    ensureCardErr := lc.ensureCardOwnership.Execute(columnId, cardId)
	if ensureCardErr != nil {
		utils.BadRequestResponse(w, r, ensureCardErr)
		return
	}
	card, findErr := lc.findCard.Execute(cardId)
	if findErr != nil {
		switch {
		case errors.Is(findErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r)
		default:
            utils.ServerErrorResponse(w, r, findErr)
		}
		return
    }
	if card == nil {
		utils.NotFoundResponse(w, r)
		return
	}
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"card": card}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
