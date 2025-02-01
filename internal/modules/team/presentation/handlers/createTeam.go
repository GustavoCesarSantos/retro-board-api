package team

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type createTeam struct {
	removeTeam application.IRemoveTeam
    saveMember application.ISaveMember
    saveTeam application.ISaveTeam
}

func NewCreateTeam(
	removeTeam application.IRemoveTeam, 
	saveMember application.ISaveMember, 
	saveTeam application.ISaveTeam,
) *createTeam {
    return &createTeam{
		removeTeam,
		saveMember,
        saveTeam,
    }
}

type CreateTeamEnvelop struct {
	Team dtos.CreateTeamResponse `json:"team"`
}

// CreateTeam creates a new team.
// @Summary Create a new team
// @Description This endpoint allows a user to create a new team and automatically assigns the user as the admin of the team.
// @Tags Team
// @Security BearerAuth
// @Param input body dtos.CreateTeamRequest true "Name of the team to be created"
// @Produce json
// @Success 201 {object} team.CreateTeamEnvelop "Team created successfully"
// @Failure 400 {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters)"
// @Failure 500 {object} utils.ErrorEnvelope "Internal server error"
// @Router /teams [post]
func(ct *createTeam) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "createTeam.go",
		"func": "createTeam.Handle",
		"line": 0,
	}
	user := utils.ContextGetUser(r)
	var input dtos.CreateTeamRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
        metadataErr["line"] = 55
		utils.BadRequestResponse(w, r, readErr, metadataErr)
		return
	}
    team, saveTeamErr := ct.saveTeam.Execute(input.Name, user.ID)
	if team == nil {
		metadataErr["line"] = 61
		utils.ServerErrorResponse(w, r, errors.New("NIL TEAM"), metadataErr)
		return
	}
	if saveTeamErr != nil {
        metadataErr["line"] = 66
		utils.ServerErrorResponse(w, r, saveTeamErr, metadataErr)
		return
	}
	adminRoleId := int64(1)
	saveAdminErr := ct.saveMember.Execute(team.ID, team.AdminId, adminRoleId)
	if saveAdminErr != nil {
		removeErr := ct.removeTeam.Execute(team.ID, team.AdminId)
		if removeErr != nil {
			switch {
			case errors.Is(removeErr, utils.ErrRecordNotFound):
                metadataErr["line"] = 77
				utils.NotFoundResponse(w, r, metadataErr)
			default:
                metadataErr["line"] = 80
				utils.ServerErrorResponse(w, r, removeErr, metadataErr)
			}
			return
		}
        metadataErr["line"] = 85
		utils.ServerErrorResponse(w, r, saveTeamErr, metadataErr)
		return
	}
	response := dtos.NewCreateTeamResponse(team.ID, team.Name, team.CreatedAt)
	writeJsonErr := utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"team": response}, nil)
	if writeJsonErr != nil {
        metadataErr["line"] = 92
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
