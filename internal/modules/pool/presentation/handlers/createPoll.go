package poll

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/presentation/dtos"
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
	//TO-DO criar caso de uso para verificar se o usuário pertence ao time
    teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
	if teamIdErr != nil {
        utils.BadRequestResponse(w, r, teamIdErr)
		return
	}
	var input dtos.CreatePollRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr)
		return
	}
    pollId := cp.savePoll.Execute(teamId, input.Poll.Name)
    for _, option := range input.Poll.Options {
        cp.saveOption.Execute(pollId, option.Text)
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
