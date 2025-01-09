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
		metadataErr := utils.Envelope{
			"file": "boardValidator.go",
			"func": "boardValidator.EnsureBoardOwnership",
			"line": 0,
		}
		teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
		if teamIdErr != nil {
			metadataErr["line"] = 30
			utils.BadRequestResponse(w, r, teamIdErr, metadataErr)
			return
		}
		boardId, boardIdErr := utils.ReadIDParam(r, "boardId")
		if boardIdErr != nil {
			metadataErr["line"] = 36
			utils.BadRequestResponse(w, r, boardIdErr, metadataErr)
			return
		}
		boards, findErr := bv.provider.FindAllBoardsByTeamId(teamId)
		if findErr != nil {
			metadataErr["line"] = 42
			utils.NotFoundResponse(w, r, metadataErr)
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
			metadataErr["line"] = 54
			utils.ForbiddenResponse(w, r, utils.ErrBoardNotInTeam, metadataErr)
			return 
		}
		next.ServeHTTP(w, r)
	})
}

func (bv *boardValidator) EnsureColumnOwnership(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metadataErr := utils.Envelope{
			"file": "boardValidator.go",
			"func": "boardValidator.EnsureColumnOwnership",
			"line": 0,
		}
		boardId, boardIdErr := utils.ReadIDParam(r, "boardId")
		if boardIdErr != nil {
			metadataErr["line"] = 71
			utils.BadRequestResponse(w, r, boardIdErr, metadataErr)
			return
		}
		columnId, columnIdErr := utils.ReadIDParam(r, "columnId")
		if columnIdErr != nil {
			metadataErr["line"] = 77
			utils.BadRequestResponse(w, r, columnIdErr, metadataErr)
			return
		}
		columns, findErr := bv.provider.FindAllColumnsByBoardId(boardId)
		if findErr != nil {
			metadataErr["line"] = 83
			utils.NotFoundResponse(w, r, metadataErr)
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
			metadataErr["line"] = 95
			utils.ForbiddenResponse(w, r, utils.ErrColumnNotInBoard, metadataErr)
			return 
		}
		next.ServeHTTP(w, r)
	})
}

func (bv *boardValidator) EnsureCardOwnership(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metadataErr := utils.Envelope{
			"file": "boardValidator.go",
			"func": "boardValidator.EnsureCardOwnership",
			"line": 0,
		}
		columnId, columnIdErr := utils.ReadIDParam(r, "columnId")
		if columnIdErr != nil {
			metadataErr["line"] = 112
			utils.NotFoundResponse(w, r, metadataErr)
			return
		}
		cardId, cardIdErr := utils.ReadIDParam(r, "cardId")
		if cardIdErr != nil {
			metadataErr["line"] = 118
			utils.NotFoundResponse(w, r, metadataErr)
			return
		}
		cards, findErr := bv.provider.FindAllCardsByColumnId(columnId)
		if findErr != nil {
			metadataErr["line"] = 124
			utils.NotFoundResponse(w, r, metadataErr)
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
			metadataErr["line"] = 136
			utils.ForbiddenResponse(w, r, utils.ErrCardNotInColumn, metadataErr)
			return 
		}
		next.ServeHTTP(w, r)
	})
}
