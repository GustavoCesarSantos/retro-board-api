package board

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type listAllBoards struct {
	findAllBoards application.IFindAllBoards
}

func NewListAllBoards(
	findAllBoards application.IFindAllBoards,
) *listAllBoards {
    return &listAllBoards {
		findAllBoards,
    }
}

type ListAllBoardsEnvelop struct {
	Boards []*dtos.ListAllBoardsResponse `json:"boards"`
}

// @Summary      List all boards for a team
// @Description  Retrieves all boards associated with a specific team.
// @Tags         Board
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId     path      int               true  "Team ID"
// @Success      200        {object} board.ListAllBoardsEnvelop "List of boards"
// @Failure      400        {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      404        {object} utils.ErrorEnvelope "Team not found"
// @Failure      500        {object} utils.ErrorEnvelope "Internal server error"
// @Router       /teams/:teamId/boards [get]
func(lb *listAllBoards) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "listAllBoards.go",
		"func": "listAllBoards.Handle",
		"line": 0,
	}
	teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
	if teamIdErr != nil {
		utils.NotFoundResponse(w, r, metadataErr)
		return
	}
	boards, findErr := lb.findAllBoards.Execute(teamId)
	if findErr != nil {
		utils.ServerErrorResponse(w, r, findErr, metadataErr)
		return
    }
	var response []*dtos.ListAllBoardsResponse
    for _, board := range boards {
        response = append(response, dtos.NewListAllBoardsResponse(
			board.ID,
			board.TeamId,
			board.Name,
			board.Active,
			board.CreatedAt,
			board.UpdatedAt,
		))
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"boards": response}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
