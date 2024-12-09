package team

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
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

func(lt *listAllTeams) Handle(w http.ResponseWriter, r *http.Request) {
    user := utils.ContextGetUser(r)
	teams := lt.findAllTeams.Execute(user.ID)
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"teams": teams}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
