package team

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type addMemberToTeam struct {
	ensureAdminMembership application.IEnsureAdminMembership
    saveMember application.ISaveMember
}

func NewAddMemberToTeam(
	ensureAdminMembership application.IEnsureAdminMembership,
	saveMember application.ISaveMember,
) *addMemberToTeam {
    return &addMemberToTeam{
		ensureAdminMembership,
        saveMember,
    }
}

func(amt *addMemberToTeam) Handle(w http.ResponseWriter, r *http.Request) {
    user := utils.ContextGetUser(r)
	teamId, readIDErr := utils.ReadIDParam(r, "teamId")
	if readIDErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	var input struct {
        MemberId int64 `json:"memberId"`
        RoleId int64 `json:"roleId"`
	}
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
    saveErr := amt.saveMember.Execute(teamId, input.MemberId, input.RoleId)
    if saveErr != nil {
		utils.ServerErrorResponse(w, r, saveErr)
		return
	}
	writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
