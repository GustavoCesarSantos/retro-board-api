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
	metadataErr := utils.Envelope{
		"file": "addMemberToTeam.go",
		"func": "addMemberToTeam.Handle",
		"line": 0,
	}
    user := utils.ContextGetUser(r)
	teamId, readIDErr := utils.ReadIDParam(r, "teamId")
	if readIDErr != nil {
		metadataErr["line"] = 51
		utils.BadRequestResponse(w, r, readIDErr, metadataErr)
		return
	}
	var input dtos.AddMemberToTeamRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		metadataErr["line"] = 58
		utils.BadRequestResponse(w, r, readErr, metadataErr)
		return
	}
	ensureAdminErr := amt.ensureAdminMembership.Execute(teamId, user.ID)
	if ensureAdminErr != nil {
		metadataErr["line"] = 64
		utils.BadRequestResponse(w, r, ensureAdminErr, metadataErr)
		return
	}
	memberInfo, findErr := amt.findMemberInfoByEmail.Execute(input.Email)
	if findErr != nil {
		metadataErr["line"] = 70
		utils.NotFoundResponse(w, r, metadataErr)
		return
	}
    saveErr := amt.saveMember.Execute(teamId, memberInfo.ID, input.RoleId)
    if saveErr != nil {
		metadataErr["line"] = 76
		utils.ServerErrorResponse(w, r, saveErr, metadataErr)
		return
	}
	writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		metadataErr["line"] = 82
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
