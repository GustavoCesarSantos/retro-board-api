package board

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type listAllCards struct {
	findAllCards application.IFindAllCards
}

func NewListAllCards(
	findAllCards application.IFindAllCards,
) *listAllCards {
    return &listAllCards {
	    findAllCards,
    }
}

type ListAllCardsEnvelop struct {
	Cards []*dtos.ListAllCardsResponse `json:"cards"`
}

// @Summary      List all cards for a column on a board
// @Description  Retrieves all cards associated with a specific column on a board.
// @Tags         Board
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId     path      int               true  "Team ID"
// @Param        boardId    path      int               true  "Board ID"
// @Param        columnId   path      int               true  "Column ID"
// @Success      200        {object} board.ListAllCardsEnvelop "List of cards"
// @Failure      400        {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or invalid input data)"
// @Failure      404        {object} utils.ErrorEnvelope "Board or column not found"
// @Failure      500        {object} utils.ErrorEnvelope "Internal server error"
// @Router       /teams/:teamId/boards/:boardId/columns/:columnId/cards [get]
func(lc *listAllCards) Handle(w http.ResponseWriter, r *http.Request) {
	columnId, columnIdErr := utils.ReadIDParam(r, "columnId")
	if columnIdErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	cards, findErr := lc.findAllCards.Execute(columnId)
	if findErr != nil {
		utils.ServerErrorResponse(w, r, findErr)
		return
    }
	var response []*dtos.ListAllCardsResponse
    for _, card := range cards {
        response = append(response, dtos.NewListAllCardsResponse(
			card.ID,
			card.ColumnId,
			card.MemberId,
			card.Text,
			card.CreatedAt,
			card.UpdatedAt,
		))
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"cards": response}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
