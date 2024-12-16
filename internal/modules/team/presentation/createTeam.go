package team

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type createTeam struct {
	removeTeam application.IRemoveTeam
    saveMember application.ISaveMember
    saveTeam application.ISaveTeam
}

func NewCreateTeam(
	removeTeam application.IRemoveTeam, 
	saveMember application.ISaveMember, 
	saveTeam application.ISaveTeam,
) *createTeam {
    return &createTeam{
		removeTeam,
		saveMember,
        saveTeam,
    }
}

func(ct *createTeam) Handle(w http.ResponseWriter, r *http.Request) {
	user := utils.ContextGetUser(r)
	var input struct {
		Name   string       `json:"name"`
	}
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr)
		return
	}
    team, saveTeamErr := ct.saveTeam.Execute(input.Name, user.ID)
	if saveTeamErr != nil {
		utils.ServerErrorResponse(w, r, saveTeamErr)
		return
	}
	adminRoleId := int64(1)
	saveAdminErr := ct.saveMember.Execute(team.ID, team.AdminId, adminRoleId)
	if saveAdminErr != nil {
		removeErr := ct.removeTeam.Execute(team.ID, team.AdminId)
		if removeErr != nil {
			switch {
			case errors.Is(removeErr, utils.ErrRecordNotFound):
				utils.NotFoundResponse(w, r)
			default:
				utils.ServerErrorResponse(w, r, removeErr)
			}
			return
		}
		utils.ServerErrorResponse(w, r, saveTeamErr)
		return
	}
	writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"team": team}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
