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
