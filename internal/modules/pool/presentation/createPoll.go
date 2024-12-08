package poll

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type Option struct {
    Text string `json:"text"`
}

type Poll struct {
    Name string `json:"name"`
    Options []Option `json:"options"`
}

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

func(cp *createPoll) Handle(w http.ResponseWriter, r *http.Request) {
    teamId, err := utils.ReadIDParam(r, "teamId")
	if err != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	var input struct {
        Poll Poll `json:"poll"`
	}
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
