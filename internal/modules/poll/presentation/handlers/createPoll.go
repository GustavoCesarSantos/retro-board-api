package poll

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/poll/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type createPoll struct {
    saveOption application.ISaveOption
    savePoll application.ISavePoll
}

func NewCreatePoll(
    saveOption application.ISaveOption, 
    savePoll application.ISavePoll,
) *createPoll {
    return &createPoll{
        saveOption,
        savePoll,
    }
}

// @Summary      Create a new poll
// @Description  Creates a new poll with options for a specific team.
// @Tags         Poll
// @Security	 BearerAuth
// @Accept       json
// @Produce      json
// @Param        teamId     path      int                        true  "Team ID"
// @Param        body       body      dtos.CreatePollRequest     true  "Poll creation data"
// @Success      204        "Poll created successfully"
// @Failure      400        {object}  utils.ErrorEnvelope "Invalid request (e.g., missing parameters or validation error)"
// @Failure      500        {object}  utils.ErrorEnvelope "Internal server error"
// @Router       /teams/:teamId/polls [post]
func(cp *createPoll) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "createPoll.go",
		"func": "createPoll.Handle",
		"line": 0,
	}
	teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
	if teamIdErr != nil {
        utils.BadRequestResponse(w, r, teamIdErr, metadataErr)
		return
	}
	var input dtos.CreatePollRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr, metadataErr)
		return
	}
    poll, savePollErr := cp.savePoll.Execute(teamId, input.Poll.Name)
	if savePollErr != nil {
		utils.ServerErrorResponse(w, r, savePollErr, metadataErr)
		return
	}
    for _, option := range input.Poll.Options {
        cp.saveOption.Execute(poll.ID, option.Text)
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
