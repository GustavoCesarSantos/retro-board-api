package team

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type addMemberToTeam struct {
    saveMember application.ISaveMember
}

func NewAddMemberToTeam(saveMember application.ISaveMember) *addMemberToTeam {
    return &addMemberToTeam{
        saveMember,
    }
}

func(amt *addMemberToTeam) Handle(w http.ResponseWriter, r *http.Request) {
    teamId, readIDErr := utils.ReadIDParam(r, "teamId")
	if readIDErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	var input struct {
        MemberId int64 `json:"memberId"`
        Role int64 `json:"role"`
	}
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr)
		return
	}
    amt.saveMember.Execute(teamId, input.MemberId, input.Role)
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
