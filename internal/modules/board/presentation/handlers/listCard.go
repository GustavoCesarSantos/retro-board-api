package board

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type ListCard struct {
	findCard application.IFindCard
}

func NewListCard(
	findCard application.IFindCard,
) *ListCard {
    return &ListCard {
		findCard,
    }
}

type ListCardEnvelop struct {
	Card *dtos.ListCardResponse `json:"card"`
}

// @Summary      Get a single card by ID
// @Description  Retrieves the details of a specific card based on its ID.
// @Tags         Board
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId     path      int               true  "Team ID"
// @Param        boardId    path      int               true  "Board ID"
// @Param        columnId   path      int               true  "Column ID"
// @Param        cardId     path      int               true  "Card ID"
// @Success      200        {object} board.ListCardEnvelop "Card details"
// @Failure      400        {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or invalid input data)"
// @Failure      404        {object} utils.ErrorEnvelope "Card not found"
// @Failure      500        {object} utils.ErrorEnvelope "Internal server error"
// @Router      /teams/:teamId/boards/:boardId/columns/:columnId/cards/:cardId [get]
func(lc *ListCard) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "listCard.go",
		"func": "listCard.Handle",
		"line": 0,
	}
	cardId, cardIdErr := utils.ReadIDParam(r, "cardId")
	if cardIdErr != nil {
		utils.BadRequestResponse(w, r, cardIdErr, metadataErr)
		return
	}
	card, findErr := lc.findCard.Execute(cardId)
	if findErr != nil {
		switch {
		case errors.Is(findErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r, metadataErr)
		default:
            utils.ServerErrorResponse(w, r, findErr, metadataErr)
		}
		return
    }
	response := dtos.NewListCardResponse(
		card.ID,
		card.ColumnId,
		card.MemberId,
		card.Text,
		card.CreatedAt,
		card.UpdatedAt,
	)
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"card": response}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
