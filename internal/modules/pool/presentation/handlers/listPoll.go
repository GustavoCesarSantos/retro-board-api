package poll

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type listPoll struct {
	ensurePollOwnership application.IEnsurePollOwnership
    findPoll application.IFindPoll
}

func NewListPoll(
	ensurePollOwnership application.IEnsurePollOwnership,
	findPoll application.IFindPoll,
) *listPoll {
    return &listPoll {
		ensurePollOwnership,
        findPoll,
    }
}

type ListPollEnvelop struct {
	Poll *dtos.ListPollResponse `json:"poll"`
}

// @Summary      Get poll details
// @Description  Retrieves detailed information about a specific poll.
// @Tags         Poll
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId     path      int    true  "Team ID"
// @Param        pollId     path      int    true  "Poll ID"
// @Success      200        {object}  poll.ListPollEnvelop  "Poll details"
// @Failure      400        {object}  utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      404        {object}  utils.ErrorEnvelope  "Poll not found"
// @Failure      500        {object}  utils.ErrorEnvelope  "Internal server error"
// @Router       /teams/:teamId/polls/:pollId [get]
func(lp *listPoll) Handle(w http.ResponseWriter, r *http.Request) {
	//TO-DO criar caso de uso para verificar se o usuário pertence ao time
    teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
	if teamIdErr != nil {
		utils.BadRequestResponse(w, r, teamIdErr)
		return
	}
	pollId, pollIdErr := utils.ReadIDParam(r, "pollId")
	if pollIdErr != nil {
		utils.BadRequestResponse(w, r, pollIdErr)
		return
	}
	ensurePollErr := lp.ensurePollOwnership.Execute(teamId, pollId)
	if ensurePollErr != nil {
		utils.BadRequestResponse(w, r, ensurePollErr)
		return
	}
	poll := lp.findPoll.Execute(pollId)
	if poll == nil {
		utils.NotFoundResponse(w, r)
		return
	}
	response := dtos.NewListPollResponse(
		poll.ID,
		poll.TeamId,
		poll.Name,
		poll.CreatedAt,
		poll.UpdatedAt,
	)
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"poll": response}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
