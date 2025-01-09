package team

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type changeMemberRole struct {
	ensureAdminMembership application.IEnsureAdminMembership
    updateRole application.IUpdateRole
}

func NewChangeMemberRole(
	ensureAdminMembership application.IEnsureAdminMembership,
	updateRole application.IUpdateRole,
) *changeMemberRole {
    return &changeMemberRole{
		ensureAdminMembership,
        updateRole,
    }
}

// ChangeMemberRole updates a team member's role.
// @Summary Change a member's role in a team
// @Description This endpoint allows an admin of a team to change the role of another member within the team.
// @Tags Team
// @Security BearerAuth
// @Param teamId path string true "ID of the team"
// @Param memberId path string true "ID of the member whose role is being changed"
// @Param input body dtos.ChangeMemberRoleRequest true "New role for the member"
// @Produce json
// @Success 204 "Role updated successfully"
// @Failure 400 {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or insufficient permissions)"
// @Failure 404 {object} utils.ErrorEnvelope "Team or member not found"
// @Failure 500 {object} utils.ErrorEnvelope "Internal server error"
// @Router /teams/:teamId/members/:memberId/roles [patch]
func(cmp *changeMemberRole) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "changeMemberRole.go",
		"func": "changeMemberRole.Handle",
		"line": 0,
	}
	user := utils.ContextGetUser(r)
	teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
	if teamIdErr != nil {
		utils.NotFoundResponse(w, r, metadataErr)
		return
	}
    memberId, memberIdErr := utils.ReadIDParam(r, "memberId")
	if memberIdErr != nil {
		utils.NotFoundResponse(w, r, metadataErr)
		return
	}
	var input dtos.ChangeMemberRoleRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr, metadataErr)
		return
	}
	ensureAdminErr := cmp.ensureAdminMembership.Execute(teamId, user.ID)
	if ensureAdminErr != nil {
		utils.BadRequestResponse(w, r, ensureAdminErr, metadataErr)
		return
	}
    updateErr := cmp.updateRole.Execute(teamId, memberId, input.Role)
	if updateErr != nil {
		switch {
		case errors.Is(updateErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r, metadataErr)
		default:
            utils.ServerErrorResponse(w, r, updateErr, metadataErr)
		}
		return
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
