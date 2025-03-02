package poll

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type ListAllPolls struct {
    findAllPolls application.IFindAllPolls
}

func NewListAllPolls(findAllPolls application.IFindAllPolls) *ListAllPolls {
    return &ListAllPolls {
        findAllPolls,
    }
}

type ListAllPollsEnvelop struct {
	Polls []*dtos.ListAllPollsResponse `json:"polls"`
}

// @Summary      List all polls
// @Description  Retrieves all polls associated with a specific team.
// @Tags         Poll
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId     path      int    true  "Team ID"
// @Success      200        {object}  poll.ListAllPollsEnvelop  "List of polls"
// @Failure      400        {object}  utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      500        {object}  utils.ErrorEnvelope  "Internal server error"
// @Router       /teams/:teamId/polls [get]
func(lp *ListAllPolls) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "listAllPolls.go",
		"func": "listAllPolls.Handle",
		"line": 0,
	}
	teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
	if teamIdErr != nil {
		utils.BadRequestResponse(w, r, teamIdErr, metadataErr)
		return
	}
	polls, findErr := lp.findAllPolls.Execute(teamId)
	if findErr != nil {
		utils.ServerErrorResponse(w, r, findErr, metadataErr)
		return
	}
	var response []*dtos.ListAllPollsResponse
    for _, poll := range polls {
        response = append(response, dtos.NewListAllPollsResponse(
			poll.ID,
			poll.TeamId,
			poll.Name,
			poll.CreatedAt,
			poll.UpdatedAt,
		))
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"polls": response}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
