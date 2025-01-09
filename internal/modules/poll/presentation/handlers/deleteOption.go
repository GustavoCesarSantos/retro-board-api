package poll

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type deleteOption struct {
	removeOption application.IRemoveOption
}

func NewDeleteOption(
	removeOption application.IRemoveOption,
) *deleteOption {
    return &deleteOption{
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
	metadataErr := utils.Envelope{
		"file": "deleteOption.go",
		"func": "deleteOption.Handle",
		"line": 0,
	}
	optionId, optionIdErr := utils.ReadIDParam(r, "optionId")
	if optionIdErr != nil {
		utils.BadRequestResponse(w, r, optionIdErr, metadataErr)
		return
	}
    do.removeOption.Execute(optionId)
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
