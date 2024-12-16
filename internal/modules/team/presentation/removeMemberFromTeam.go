package team

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type removeMemberFromTeam struct {
	ensureAdminMembership application.IEnsureAdminMembership
    removeMember application.IRemoveMember
}

func NewRemoveMemberFromTeam(
	ensureAdminMembership application.IEnsureAdminMembership,
    removeMember application.IRemoveMember,
) *removeMemberFromTeam {
    return &removeMemberFromTeam{
        ensureAdminMembership,
        removeMember,
    }
}

func(rmt *removeMemberFromTeam) Handle(w http.ResponseWriter, r *http.Request) {
    user := utils.ContextGetUser(r)
    teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
	if teamIdErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
    memberId, memberIdErr := utils.ReadIDParam(r, "memberId")
	if memberIdErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	ensureAdminErr := rmt.ensureAdminMembership.Execute(teamId, user.ID)
	if ensureAdminErr != nil {
		utils.BadRequestResponse(w, r, ensureAdminErr)
		return
	}
    removeErr := rmt.removeMember.Execute(teamId, memberId)
	if removeErr != nil {
		switch {
		case errors.Is(removeErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r)
		default:
            utils.ServerErrorResponse(w, r, removeErr)
		}
		return
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
