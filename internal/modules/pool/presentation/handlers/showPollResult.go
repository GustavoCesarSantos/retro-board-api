package poll

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type showPollResult struct {
	ensurePollOwnership application.IEnsurePollOwnership
    countVotesByPollId application.ICountVotesByPollId
}

func NewShowPollResult(
	ensurePollOwnership application.IEnsurePollOwnership,
    countVotesByPollId application.ICountVotesByPollId,
) *showPollResult {
    return &showPollResult {
		ensurePollOwnership,
        countVotesByPollId,
    }
}

type ShowPollResultEnvelop struct {
	Result *dtos.ShowPollResultResponse `json:"result"`
}

// @Summary      Show poll result
// @Description  Retrieves the result of a poll, including the total votes, votes per option, and the winning option.
// @Tags         Poll
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId     path      int    true  "Team ID"
// @Param        pollId     path      int    true  "Poll ID"
// @Success      200        {object}  poll.ShowPollResultEnvelop  "Poll results"
// @Failure      400        {object}  utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      404        {object}  utils.ErrorEnvelope  "Poll not found"
// @Failure      500        {object}  utils.ErrorEnvelope  "Internal server error"
// @Router      /teams/:teamId/polls/:pollId/result [get]
func(spr *showPollResult) Handle(w http.ResponseWriter, r *http.Request) {
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
	ensurePollErr := spr.ensurePollOwnership.Execute(teamId, pollId)
	if ensurePollErr != nil {
		utils.BadRequestResponse(w, r, ensurePollErr)
		return
	}
	result, resultErr := spr.countVotesByPollId.Execute(pollId)
	if resultErr == nil {
		utils.BadRequestResponse(w, r, resultErr)
		return
	}
	response := dtos.NewShowPollResultResponse(
		result.Options,
		result.Winner,
		result.Total,
	)
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"result": response}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
