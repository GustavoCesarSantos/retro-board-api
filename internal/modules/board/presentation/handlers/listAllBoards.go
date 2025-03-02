package board

import (
	"math"
	"net/http"
	"strconv"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type ListAllBoards struct {
	findAllBoards application.IFindAllBoards
}

func NewListAllBoards(
	findAllBoards application.IFindAllBoards,
) *ListAllBoards {
    return &ListAllBoards {
		findAllBoards,
    }
}

type ListAllBoardsEnvelop struct {
	Boards dtos.ListAllBoardsResponsePaginated `json:"boards"`
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
func(lb *ListAllBoards) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "listAllBoards.go",
		"func": "listAllBoards.Handle",
		"line": 0,
	}
	teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
	if teamIdErr != nil {
		utils.BadRequestResponse(w, r, teamIdErr, metadataErr)
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
	boards, findErr := lb.findAllBoards.Execute(teamId, limit, lastID)
	if findErr != nil {
		utils.ServerErrorResponse(w, r, findErr, metadataErr)
		return
    }
	var response dtos.ListAllBoardsResponsePaginated
    for _, board := range boards.Items {
        response.Items = append(response.Items, dtos.NewListAllBoardsResponse(
			board.ID,
			board.TeamId,
			board.Name,
			board.Active,
			board.CreatedAt,
			board.UpdatedAt,
		))
    }
	response.NextCursor = boards.NextCursor
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"boards": response}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
