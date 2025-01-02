package middleware

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type boardValidator struct {
    provider interfaces.IBoardApi
}

func NewBoardValidator(provider interfaces.IBoardApi) *boardValidator {
    return &boardValidator{
        provider,
    }
}


func (bv *boardValidator) EnsureBoardOwnership(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
		if teamIdErr != nil {
			utils.BadRequestResponse(w, r, teamIdErr)
			return
		}
		boardId, boardIdErr := utils.ReadIDParam(r, "boardId")
		if boardIdErr != nil {
			utils.BadRequestResponse(w, r, boardIdErr)
			return
		}
		boards, findErr := bv.provider.FindAllBoardsByTeamId(teamId)
		if findErr != nil {
			utils.NotFoundResponse(w, r)
			return 
		}
		found := false
		for _, board := range boards {
			if board.ID == boardId {
				found = true
				break
			}
		}
		if !found {
			utils.ForbiddenResponse(w, r, utils.ErrBoardNotInTeam)
			return 
		}
		next.ServeHTTP(w, r)
	})
}

func (bv *boardValidator) EnsureColumnOwnership(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		boardId, boardIdErr := utils.ReadIDParam(r, "boardId")
		if boardIdErr != nil {
			utils.BadRequestResponse(w, r, boardIdErr)
			return
		}
		columnId, columnIdErr := utils.ReadIDParam(r, "columnId")
		if columnIdErr != nil {
			utils.BadRequestResponse(w, r, columnIdErr)
			return
		}
		columns, findErr := bv.provider.FindAllColumnsByBoardId(boardId)
		if findErr != nil {
			utils.NotFoundResponse(w, r)
			return 
		}
		found := false
		for _, column := range columns {
			if column.ID == columnId {
				found = true
				break
			}
		}
		if !found {
			utils.ForbiddenResponse(w, r, utils.ErrColumnNotInBoard)
			return 
		}
		next.ServeHTTP(w, r)
	})
}

func (bv *boardValidator) EnsureCardOwnership(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		cards, findErr := bv.provider.FindAllCardsByColumnId(columnId)
		if findErr != nil {
			utils.NotFoundResponse(w, r)
			return 
		}
		found := false
		for _, card := range cards {
			if card.ID == cardId {
				found = true
				break
			}
		}
		if !found {
			utils.ForbiddenResponse(w, r, utils.ErrCardNotInColumn)
			return 
		}
		next.ServeHTTP(w, r)
	})
}
