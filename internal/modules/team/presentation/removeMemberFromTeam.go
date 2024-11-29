package team

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type removeMemberFromTeam struct {
    removeMember application.IRemoveMember
}

func NewRemoveMemberFromTeam(removeMember application.IRemoveMember) *removeMemberFromTeam {
    return &removeMemberFromTeam{
        removeMember,
    }
}

func(rmt *removeMemberFromTeam) Handle(w http.ResponseWriter, r *http.Request) {
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
    rmt.removeMember.Execute(teamId, memberId)
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
