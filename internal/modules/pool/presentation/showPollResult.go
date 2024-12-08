package poll

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/application"
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

func(spr *showPollResult) Handle(w http.ResponseWriter, r *http.Request) {
    teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
	if teamIdErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	pollId, pollIdErr := utils.ReadIDParam(r, "pollId")
	if pollIdErr != nil {
		utils.NotFoundResponse(w, r)
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
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"result": result}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
