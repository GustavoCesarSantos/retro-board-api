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

// SignoutUser handles the user sign-out process.
// @Summary Sign out the current user
// @Description This endpoint invalidates the user's current session by incrementing their token version. 
// @Tags Identity
// @Security BearerAuth
// @Produce json
// @Success 204 "User signed out successfully"
// @Failure 404 {object} utils.ErrorEnvelope "User not found"
// @Failure 500 {object} utils.ErrorEnvelope "Internal server error"
// @Router /auth/signout [post]
func(su *signoutUser) Handle(w http.ResponseWriter, r *http.Request) {
    metadataErr := utils.Envelope{
		"file": "signoutUser.go",
		"func": "signoutUser.Handle",
		"line": 0,
	}
	user := utils.ContextGetUser(r)
    incrementErr := su.incrementVersion.Execute(user)
    if incrementErr != nil {
		switch {
		case errors.Is(incrementErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r, metadataErr)
		default:
            utils.ServerErrorResponse(w, r, incrementErr, metadataErr)
		}
		return
    }
    writeJsonErr := utils.WriteJSON(w, http.StatusNoContent, nil, nil)
	if writeJsonErr != nil {
		utils.ServerErrorResponse(w, r, writeJsonErr, metadataErr)
	}
}
