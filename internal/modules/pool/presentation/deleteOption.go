package poll

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type deleteOption struct {
	ensurePollOwnership application.IEnsurePollOwnership
	ensureOptionOwnership application.IEnsureOptionOwnership
    removeOption application.IRemoveOption
}

func NewDeleteOption(
	ensurePollOwnership application.IEnsurePollOwnership,
    ensureOptionOwnership application.IEnsureOptionOwnership,
    removeOption application.IRemoveOption,
) *deleteOption {
    return &deleteOption{
        ensurePollOwnership,
		ensureOptionOwnership,
        removeOption,
    }
}

func(do *deleteOption) Handle(w http.ResponseWriter, r *http.Request) {
    teamId, err := utils.ReadIDParam(r, "teamId")
	if err != nil {
		utils.NotFoundResponse(w, r)
		return
	}
    pollId, pollIdErr := utils.ReadIDParam(r, "pollId")
	if pollIdErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	optionId, optionIdErr := utils.ReadIDParam(r, "optionId")
	if optionIdErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	ensurePollErr := do.ensurePollOwnership.Execute(teamId, pollId)
	if ensurePollErr != nil {
		utils.BadRequestResponse(w, r, ensurePollErr)
		return
	}
	ensureOptionErr := do.ensureOptionOwnership.Execute(pollId, optionId)
	if ensureOptionErr != nil {
		utils.BadRequestResponse(w, r, ensureOptionErr)
		return
	}
    do.removeOption.Execute(optionId)
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
