package identity

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
	"github.com/google/uuid"
	"github.com/markbates/goth/gothic"
)

type signinWithGoogleCallback struct {
	findUserByEmail   application.IFindUserByEmail
	saveUser          application.ISaveUser
	updateSigninToken application.IUpdateSigninToken
}

func NewSigninWithGoogleCallback(
	findUserByEmail application.IFindUserByEmail,
	saveUser application.ISaveUser,
	updateSigninToken application.IUpdateSigninToken,
) *signinWithGoogleCallback {
	return &signinWithGoogleCallback{
		findUserByEmail,
		saveUser,
		updateSigninToken,
	}
}

func (sgc *signinWithGoogleCallback) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "signinWithGoogleCallback.go",
		"func": "signinWithGoogleCallback.Handle",
		"line": 0,
	}
	q := r.URL.Query()
	q.Add("provider", "google")
	r.URL.RawQuery = q.Encode()
	userFromGoogle, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		utils.ServerErrorResponse(w, r, err, metadataErr)
		return
	}
	user, userErr := sgc.findUserByEmail.Execute(userFromGoogle.Email)
	if userErr != nil {
		switch {
		case errors.Is(userErr, utils.ErrRecordNotFound):
			user, userErr = sgc.saveUser.Execute(userFromGoogle.Name, userFromGoogle.Email)
			if userErr != nil {
				utils.ServerErrorResponse(w, r, userErr, metadataErr)
				return
			}
		default:
			utils.ServerErrorResponse(w, r, userErr, metadataErr)
			return
		}
	}
    signinToken := uuid.NewString()
	updateSigninTokenErr := sgc.updateSigninToken.Execute(user.ID, signinToken)
	if updateSigninTokenErr != nil {
		utils.ServerErrorResponse(w, r, updateSigninTokenErr, metadataErr)
	}
	frontendURL := fmt.Sprintf("http://localhost:5173/identity/auth?code=%s", signinToken)
	http.Redirect(w, r, frontendURL, http.StatusPermanentRedirect)
}
