package board

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/board/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type listAllBoards struct {
    findAllBoards application.IFindAllBoards
}

func NewListAllBoards(findAllBoards application.IFindAllBoards) *listAllBoards {
    return &listAllBoards {
        findAllBoards,
    }
}

func(lb *listAllBoards) Handle(w http.ResponseWriter, r *http.Request) {
	//TO-DO criar caso de uso para verificar se o usuário pertence ao time
    teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
	if teamIdErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	boards, findErr := lb.findAllBoards.Execute(teamId)
	if findErr != nil {
		utils.ServerErrorResponse(w, r, findErr)
		return
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"boards": boards}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
