package team

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/team/application"
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

func(st *showTeam) Handle(w http.ResponseWriter, r *http.Request) {
    user := utils.ContextGetUser(r)
	id, err := utils.ReadIDParam(r, "teamId")
	if err != nil {
		utils.NotFoundResponse(w, r)
		return
	}
    team, findErr := st.findTeam.Execute(id, user.ID)
	if findErr != nil {
		switch {
		case errors.Is(findErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r)
		default:
            utils.ServerErrorResponse(w, r, findErr)
		}
		return
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"teams": team}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
