package team

import (
	"math"
	"net/http"
	"strconv"

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
	Teams dtos.ListAllTeamsResponsePaginated `json:"teams"`
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
	metadataErr := utils.Envelope{
		"file": "listAllTeams.go",
		"func": "listAllTeams.Handle",
		"line": 0,
	}
    user := utils.ContextGetUser(r)
	limitStr := r.URL.Query().Get("limit")
	lastIDStr := r.URL.Query().Get("lastId")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		metadataErr["line"] = 47
		utils.BadRequestResponse(w, r, utils.ErrMissingOrInvalidLimitQueryParam, metadataErr)
		return
	}
	lastID := math.MaxInt64
	if lastIDStr != "" {
		lastID, err = strconv.Atoi(lastIDStr)
		if err != nil {
			metadataErr["line"] = 55
			utils.BadRequestResponse(w, r, utils.ErrInvalidLimitQueryParam, metadataErr)
			return
		}
	}
	teams, findErr := lt.findAllTeams.Execute(user.ID, limit, lastID)
	if findErr != nil {
		metadataErr["line"] = 61
		utils.ServerErrorResponse(w, r, findErr, metadataErr)
		return
	}
	var response dtos.ListAllTeamsResponsePaginated
    for _, team := range teams.Items {
        response.Items = append(response.Items, dtos.NewListAllTeamsResponse(
			team.ID,
			team.Name,
			team.CreatedAt,
			team.UpdatedAt,
		))
    }
	response.NextCursor = teams.NextCursor
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"teams": response}, nil)
	if writeJsonErr != nil {
		metadataErr["line"] = 77
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
