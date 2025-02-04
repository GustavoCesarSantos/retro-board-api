package team

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type showTeam struct {
    findTeam application.IFindTeam
}

func NewShowTeam(findTeam application.IFindTeam) *showTeam {
    return &showTeam {
        findTeam,
    }
}

type ShowTeamEnvelop struct {
	Team *dtos.ShowTeamResponse `json:"team"`
}

// ShowTeam retrieves details of a specific team.
// @Summary Get details of a specific team
// @Description This endpoint retrieves detailed information about a specific team. The user must be a member of the team to access this information.
// @Tags Team
// @Security BearerAuth
// @Param teamId path int true "Team ID"
// @Produce json
// @Success 200 {object} team.ShowTeamEnvelop "Details of the team"
// @Failure 404 {object} utils.ErrorEnvelope "Not Found - Team not found or user not a member"
// @Failure 500 {object} utils.ErrorEnvelope "Internal server error"
// @Router /teams/:teamId [get]
func(st *showTeam) Handle(w http.ResponseWriter, r *http.Request) {
    metadataErr := utils.Envelope{
		"file": "showTeam.go",
		"func": "showTeam.Handle",
		"line": 0,
	}
	user := utils.ContextGetUser(r)
	id, readErr := utils.ReadIDParam(r, "teamId")
	if readErr != nil {
		metadataErr["line"] = 46
		utils.BadRequestResponse(w, r, readErr, metadataErr)
		return
	}
    team, findErr := st.findTeam.Execute(id, user.ID)
	if findErr != nil {
		switch {
		case errors.Is(findErr, utils.ErrRecordNotFound):
			metadataErr["line"] = 54
            utils.NotFoundResponse(w, r, metadataErr)
		default:
			metadataErr["line"] = 57
            utils.ServerErrorResponse(w, r, findErr, metadataErr)
		}
		return
    }
	response := dtos.NewShowTeamResponse(
		team.ID, 
		team.Name, 
		team.CreatedAt, 
		team.UpdatedAt,
	)
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"team": response}, nil)
	if writeJsonErr != nil {
		metadataErr["line"] = 70
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
