package identity

import (
	"errors"
	"net/http"
	"time"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type refreshAuthToken struct {
    createAuthToken application.ICreateAuthToken
    decodeAuthToken application.IDecodeAuthToken
    findUserByEmail application.IFindUserByEmail
}

func NewRefreshAuthToken(
	createAuthToken application.ICreateAuthToken,
	decodeAuthToken application.IDecodeAuthToken,
	findUserByEmail application.IFindUserByEmail,
) *refreshAuthToken {
    return &refreshAuthToken{
        createAuthToken,
		decodeAuthToken,
        findUserByEmail,
    }
}

func(rt *refreshAuthToken) Handle(w http.ResponseWriter, r *http.Request) {
	var input struct {
		RefreshToken   string       `json:"refreshToken"`
	}
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr)
		return
	}
	decodedToken, decodedErr := rt.decodeAuthToken.Execute(input.RefreshToken)
	if decodedErr != nil {
		utils.BadRequestResponse(w, r, decodedErr)
		return
	}
	user := rt.findUserByEmail.Execute(decodedToken.Email)
    if user == nil {
        utils.NotFoundResponse(w, r)
		return
    }
	if user.Version != decodedToken.Version {
		utils.BadRequestResponse(w, r, errors.New("INVALID CREDENTIALS"))
		return
	}
	accessToken, accessTokenErr := rt.createAuthToken.Execute(*user, 15 * time.Minute)
	if accessTokenErr != nil {
		utils.ServerErrorResponse(w, r, accessTokenErr)
	}
	data := utils.Envelope{
		"accessToken": accessToken,
	}
	writeErr := utils.WriteJSON(w, http.StatusOK, data, nil)
	if writeErr != nil {
		utils.ServerErrorResponse(w, r, writeErr)
	}
}
