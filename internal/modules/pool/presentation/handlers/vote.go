package poll

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type vote struct {
	ensurePollOwnership application.IEnsurePollOwnership
	ensureOptionOwnership application.IEnsureOptionOwnership
    saveVote application.ISaveVote
}

func NewVote(
	ensurePollOwnership application.IEnsurePollOwnership,
    ensureOptionOwnership application.IEnsureOptionOwnership,
    saveVote application.ISaveVote,
) *vote {
    return &vote{
        ensurePollOwnership,
		ensureOptionOwnership,
        saveVote,
    }
}

// @Summary      Vote in a poll
// @Description  Cast a vote for a specific option in a poll.
// @Tags         Poll
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId     path      int    true  "Team ID"
// @Param        pollId     path      int    true  "Poll ID"
// @Param        optionId   path      int    true  "Option ID"
// @Success      204        "Vote successfully recorded"
// @Failure      400        {object}  utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      403        {object}  utils.ErrorEnvelope  "User is not allowed to vote on this poll"
// @Failure      404        {object}  utils.ErrorEnvelope  "Poll or option not found"
// @Failure      500        {object}  utils.ErrorEnvelope  "Internal server error"
// @Router       /teams/:teamId/polls/:pollId/options/:optionId/vote [post]
func(v *vote) Handle(w http.ResponseWriter, r *http.Request) {
	//TO-DO criar caso de uso para verificar se o usuário pertence ao time
    user := utils.ContextGetUser(r)
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
	optionId, optionIdErr := utils.ReadIDParam(r, "optionId")
	if optionIdErr != nil {
		utils.BadRequestResponse(w, r, optionIdErr)
		return
	}
	ensurePollErr := v.ensurePollOwnership.Execute(teamId, pollId)
	if ensurePollErr != nil {
		utils.BadRequestResponse(w, r, ensurePollErr)
		return
	}
	ensureOptionErr := v.ensureOptionOwnership.Execute(pollId, optionId)
	if ensureOptionErr != nil {
		utils.BadRequestResponse(w, r, ensureOptionErr)
		return
	}
    v.saveVote.Execute(user.ID, optionId)
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
