package team

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type addMemberToTeam struct {
	ensureAdminMembership application.IEnsureAdminMembership
	findMemberInfoByEmail application.IFindMemberInfoByEmail
    saveMember application.ISaveMember
}

func NewAddMemberToTeam(
	ensureAdminMembership application.IEnsureAdminMembership,
	findMemberInfoByEmail application.IFindMemberInfoByEmail,
	saveMember application.ISaveMember,
) *addMemberToTeam {
    return &addMemberToTeam{
		ensureAdminMembership,
		findMemberInfoByEmail,
        saveMember,
    }
}

// Handle adds a member to a team.
// @Summary Adds a member to a team
// @Description Adds a new member to a team with a specific role, provided the authenticated user is an admin of the team.
// @Tags Team
// @Param teamId path int true "Team ID"
// @Param input body dtos.AddMemberToTeamRequest true "Member details (ID and role)"
// @Security BearerAuth
// @Produce json
// @Success 204 "Member added successfully"
// @Failure 400 {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or insufficient permissions)"
// @Failure 404 {object} utils.ErrorEnvelope "Team not found"
// @Failure 500 {object} utils.ErrorEnvelope "Internal server error"
// @Router /teams/:teamId/members [post]
func(amt *addMemberToTeam) Handle(w http.ResponseWriter, r *http.Request) {
    user := utils.ContextGetUser(r)
	teamId, readIDErr := utils.ReadIDParam(r, "teamId")
	if readIDErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	var input dtos.AddMemberToTeamRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr)
		return
	}
	ensureAdminErr := amt.ensureAdminMembership.Execute(teamId, user.ID)
	if ensureAdminErr != nil {
		utils.BadRequestResponse(w, r, ensureAdminErr)
		return
	}
	memberInfo, findErr := amt.findMemberInfoByEmail.Execute(input.Email)
	if findErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
    saveErr := amt.saveMember.Execute(teamId, memberInfo.ID, input.RoleId)
    if saveErr != nil {
		utils.ServerErrorResponse(w, r, saveErr)
		return
	}
	writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
