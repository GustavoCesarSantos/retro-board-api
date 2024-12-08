package poll

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/application"
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

func(lp *listPoll) Handle(w http.ResponseWriter, r *http.Request) {
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
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"poll": poll}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
