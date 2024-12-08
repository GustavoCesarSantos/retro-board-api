package poll

import (
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/pool/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type listAllPolls struct {
    findAllPolls application.IFindAllPolls
}

func NewListAllPolls(findAllPolls application.IFindAllPolls) *listAllPolls {
    return &listAllPolls {
        findAllPolls,
    }
}

func(lp *listAllPolls) Handle(w http.ResponseWriter, r *http.Request) {
    teamId, teamIdErr := utils.ReadIDParam(r, "teamId")
	if teamIdErr != nil {
		utils.NotFoundResponse(w, r)
		return
	}
	polls := lp.findAllPolls.Execute(teamId)
    writeJsonErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"polls": polls}, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
