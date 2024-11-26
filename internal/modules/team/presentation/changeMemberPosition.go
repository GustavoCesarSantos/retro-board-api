package team

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type changeMemberPosition struct {
    uploadRole application.IUpdateRole
}

func NewChangeMemberPosition(uploadRole application.IUpdateRole) *changeMemberPosition {
    return &changeMemberPosition{
        uploadRole,
    }
}

func(cmp *changeMemberPosition) Handle(w http.ResponseWriter, r *http.Request) {
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
    writeJsonErr := utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"status": "success"}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
