package identity

import (
	"net/http"
	"time"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type signinUser struct {
    creatAuthToken application.ICreateAuthToken
    findUserByEmail application.IFindUserByEmail
	updateVersion application.IUpdateVersion
}

func NewSigninUser(
	createAuthToken application.ICreateAuthToken,  
	findUserByEmail application.IFindUserByEmail,
	updateVersion application.IUpdateVersion,
) *signinUser {
    return &signinUser{
        createAuthToken,
        findUserByEmail,
		updateVersion,
    }
}

func(su *signinUser) Handle(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email   string       `json:"email"`
	}
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr)
		return
	}
	user := su.findUserByEmail.Execute(input.Email);
    if user == nil {
        utils.NotFoundResponse(w, r)
		return
    }
	user.Version++
	su.updateVersion.Execute(user.ID, user.Version)
    accessToken, accessTokenErr := su.creatAuthToken.Execute(*user, 15 * time.Minute)
	if accessTokenErr != nil {
		utils.ServerErrorResponse(w, r, accessTokenErr)
	}
    refreshToken, refreshTokenErr := su.creatAuthToken.Execute(*user, 24 * time.Hour * time.Duration(7))
    if refreshTokenErr != nil {
		utils.ServerErrorResponse(w, r, refreshTokenErr)
	}
	data := utils.Envelope{
		"accessToken": accessToken,
		"refreshToken": refreshToken,
	}
	writeErr := utils.WriteJSON(w, http.StatusOK, data, nil)
	if writeErr != nil {
		utils.ServerErrorResponse(w, r, writeErr)
	}
}
