package team

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type RemoveMemberFromTeam struct {
	ensureAdminMembership application.IEnsureAdminMembership
    removeMember application.IRemoveMember
}

// RemoveMemberFromTeam removes a member from a team.
// @Summary Remove a member from a team
// @Description This endpoint removes a specific member from a team. Only admins are allowed to perform this operation.
// @Tags Team
// @Security BearerAuth
// @Param teamId path int true "Team ID"
// @Param memberId path int true "Member ID"
// @Produce json
// @Success 204 "Member removed successfully"
// @Failure 400 {object} utils.ErrorEnvelope "Invalid request (e.g., Invalid input or unauthorized operation)"
// @Failure 404 {object} utils.ErrorEnvelope "Not Found - Team or member not found"
// @Failure 500 {object} utils.ErrorEnvelope "Internal server error"
// @Router /teams/:teamId/members/:memberId [delete]
func NewRemoveMemberFromTeam(
	ensureAdminMembership application.IEnsureAdminMembership,
    removeMember application.IRemoveMember,
) *RemoveMemberFromTeam {
    return &RemoveMemberFromTeam{
        ensureAdminMembership,
        removeMember,
    }
}

func(rmt *RemoveMemberFromTeam) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "removeMemberFromTeam.go",
		"func": "removeMemberFromTeam.Handle",
		"line": 0,
	}
    user := utils.ContextGetUser(r)
    teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
	if teamIdErr != nil {
		metadataErr["line"] = 48
		utils.BadRequestResponse(w, r, teamIdErr, metadataErr)
		return
	}
    memberId, memberIdErr := utils.ReadIDParam(r, "memberId")
	if memberIdErr != nil {
		metadataErr["line"] = 54
		utils.BadRequestResponse(w, r, memberIdErr, metadataErr)
		return
	}
    if memberId == user.ID {
		metadataErr["line"] = 59
		utils.BadRequestResponse(w, r, errors.New("CAN'T REMOVE THIS USER."), metadataErr)
		return
    }
	ensureAdminErr := rmt.ensureAdminMembership.Execute(teamId, user.ID)
	if ensureAdminErr != nil {
		metadataErr["line"] = 65
		utils.BadRequestResponse(w, r, ensureAdminErr, metadataErr)
		return
	}
    removeErr := rmt.removeMember.Execute(teamId, memberId)
	if removeErr != nil {
		switch {
		case errors.Is(removeErr, utils.ErrRecordNotFound):
			metadataErr["line"] = 73
            utils.NotFoundResponse(w, r, metadataErr)
		default:
			metadataErr["line"] = 71
            utils.ServerErrorResponse(w, r, removeErr, metadataErr)
		}
		return
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		metadataErr["line"] = 78
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
