package team

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type listAllMembers struct {
	findAllMembers application.IFindAllMembers
}

func NewListAllMembers(findAllMembers application.IFindAllMembers) *listAllMembers {
	return &listAllMembers{
		findAllMembers,
	}
}

type ListAllMembersEnvelop struct {
	Members []*dtos.ListAllMembersResponse `json:"members"`
}

func (lt *listAllMembers) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "listAllMembers.go",
		"func": "listAllMembers.Handle",
		"line": 0,
	}
	teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
	if teamIdErr != nil {
		metadataErr["line"] = 33
		utils.BadRequestResponse(w, r, teamIdErr, metadataErr)
		return
	}
	members, findErr := lt.findAllMembers.Execute(teamId)
	if findErr != nil {
		metadataErr["line"] = 39
		utils.ServerErrorResponse(w, r, findErr, metadataErr)
		return
	}
	var response []*dtos.ListAllMembersResponse
	for _, member := range members {
		response = append(response, dtos.NewListAllMembersResponse(
			member.ID,
			member.TeamId,
			member.User.Name,
			member.User.Email,
			member.Role.ID,
			member.Role.Name,
			member.Status,
			member.CreatedAt,
			member.UpdatedAt,
		))
	}
	writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"members": response}, nil)
	if writeJsonErr != nil {
		metadataErr["line"] = 57
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
