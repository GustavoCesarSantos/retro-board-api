package team

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type DeleteTeam struct {
	ensureAdminMembership application.IEnsureAdminMembership
    removeTeam application.IRemoveTeam
}

func NewDeleteTeam(
	ensureAdminMembership application.IEnsureAdminMembership,
    removeTeam application.IRemoveTeam,
) *DeleteTeam {
    return &DeleteTeam{
        ensureAdminMembership,
        removeTeam,
    }
}

func(dt *DeleteTeam) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "deleteTeam.go",
		"func": "deleteTeam.Handle",
		"line": 0,
	}
    user := utils.ContextGetUser(r)
    teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
	if teamIdErr != nil {
		metadataErr["line"] = 35
		utils.BadRequestResponse(w, r, teamIdErr, metadataErr)
		return
	}
	ensureAdminErr := dt.ensureAdminMembership.Execute(teamId, user.ID)
	if ensureAdminErr != nil {
		metadataErr["line"] = 41
		utils.BadRequestResponse(w, r, ensureAdminErr, metadataErr)
		return
	}
    removeErr := dt.removeTeam.Execute(teamId, user.ID)
	if removeErr != nil {
		switch {
		case errors.Is(removeErr, utils.ErrRecordNotFound):
			metadataErr["line"] = 49
            utils.NotFoundResponse(w, r, metadataErr)
		default:
			metadataErr["line"] =  52
            utils.ServerErrorResponse(w, r, removeErr, metadataErr)
		}
		return
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		metadataErr["line"] = 59
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
