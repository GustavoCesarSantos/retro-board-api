package team

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type changeMemberRole struct {
    uploadRole application.IUpdateRole
}

func NewChangeMemberRole(uploadRole application.IUpdateRole) *changeMemberRole {
    return &changeMemberRole{
        uploadRole,
    }
}

func(cmp *changeMemberRole) Handle(w http.ResponseWriter, r *http.Request) {
	/*
	TO-DO: Adicionar uma verificação da role do usuário que está executando a
	ação. Somente usuários com a role: Admin podem trocar a role de outros usuá
	rios
	*/
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
    cmp.uploadRole.Execute(teamId, memberId, input.Role)
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
