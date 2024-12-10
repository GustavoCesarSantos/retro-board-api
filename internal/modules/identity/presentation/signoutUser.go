package identity

import (
	"errors"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type signoutUser struct {
	incrementVersion application.IIncrementVersion
}

func NewSignoutUser(
	incrementVersion application.IIncrementVersion,
) *signoutUser {
    return &signoutUser{
		incrementVersion,
    }
}

func(su *signoutUser) Handle(w http.ResponseWriter, r *http.Request) {
    user := utils.ContextGetUser(r)
    incrementErr := su.incrementVersion.Execute(user)
    if incrementErr != nil {
		switch {
		case errors.Is(incrementErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r)
		default:
            utils.ServerErrorResponse(w, r, incrementErr)
		}
		return
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr)
	}
}
