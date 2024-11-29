package team

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type createTeam struct {
    saveTeam application.ISaveTeam
}

func NewCreateTeam(saveTeam application.ISaveTeam) *createTeam {
    return &createTeam{
        saveTeam,
    }
}

func(ct *createTeam) Handle(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name   string       `json:"name"`
	}
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr)
		return
	}
    ct.saveTeam.Execute(input.Name, 1)
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
