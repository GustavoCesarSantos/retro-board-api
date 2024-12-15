package team

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
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

func(cmp *changeMemberRole) Handle(w http.ResponseWriter, r *http.Request) {
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
	var input struct {
        Role int64 `json:"role"`
	}
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr)
		return
	}
	ensureAdminErr := cmp.ensureAdminMembership.Execute(teamId, user.ID)
	if ensureAdminErr != nil {
		utils.BadRequestResponse(w, r, ensureAdminErr)
		return
	}
    updateErr := cmp.updateRole.Execute(teamId, memberId, input.Role)
	if updateErr != nil {
		switch {
		case errors.Is(updateErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r)
		default:
            utils.ServerErrorResponse(w, r, updateErr)
		}
		return
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
