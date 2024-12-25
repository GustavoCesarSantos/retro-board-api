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

// @Summary      Delete an option from a poll
// @Description  Deletes an option from a specific poll, ensuring proper ownership validation.
// @Tags         Poll
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId     path      int    true  "Team ID"
// @Param        pollId     path      int    true  "Poll ID"
// @Param        optionId   path      int    true  "Option ID"
// @Success      204        "Option deleted successfully"
// @Failure      400        {object}  utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      500        {object}  utils.ErrorEnvelope  "Internal server error"
// @Router       /teams/:teamId/polls/:pollId/options/:optionId [delete]
func(do *deleteOption) Handle(w http.ResponseWriter, r *http.Request) {
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
	optionId, optionIdErr := utils.ReadIDParam(r, "optionId")
	if optionIdErr != nil {
		utils.BadRequestResponse(w, r, optionIdErr)
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
