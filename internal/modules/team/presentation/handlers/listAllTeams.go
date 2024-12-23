package team

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type listAllTeams struct {
    findAllTeams application.IFindAllTeams
}

func NewListAllTeams(findAllTeams application.IFindAllTeams) *listAllTeams {
    return &listAllTeams {
        findAllTeams,
    }
}

type ListAllTeamsEnvelop struct {
	Teams []*dtos.ListAllTeamsResponse `json:"teams"`
}

// ListAllTeams retrieves all teams associated with the authenticated user.
// @Summary List all teams
// @Description This endpoint fetches all teams associated with the currently authenticated user.
// @Tags Team
// @Security BearerAuth
// @Produce json
// @Success 200 {object} team.ListAllTeamsEnvelop "List of teams successfully retrieved"
// @Failure 500 {object} utils.ErrorEnvelope "Internal server error"
// @Router /teams [get]
func(lt *listAllTeams) Handle(w http.ResponseWriter, r *http.Request) {
    user := utils.ContextGetUser(r)
	teams, findErr := lt.findAllTeams.Execute(user.ID)
	if findErr != nil {
		utils.ServerErrorResponse(w, r, findErr)
		return
	}
	var response []*dtos.ListAllTeamsResponse
    for _, team := range teams {
        response = append(response, dtos.NewListAllTeamsResponse(
			team.ID,
			team.Name,
			team.CreatedAt,
			team.UpdatedAt,
		))
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"teams": response}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
