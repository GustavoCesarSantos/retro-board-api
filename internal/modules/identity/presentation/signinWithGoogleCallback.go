package identity

import (
	"fmt"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
	"github.com/markbates/goth/gothic"
)

type signinWithGoogleCallback struct {
    findUserByEmail application.IFindUserByEmail
    saveUser application.ISaveUser
}

func NewSigninWithGoogleCallback(findUserByEmail application.IFindUserByEmail, saveUser application.ISaveUser) *signinWithGoogleCallback {
    return &signinWithGoogleCallback{
        findUserByEmail,
        saveUser,
    }
}

func(sgc *signinWithGoogleCallback) Handle(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query()
    q.Add("provider", "google")
    r.URL.RawQuery = q.Encode()
    userFromGoogle, err := gothic.CompleteUserAuth(w, r)
    if err != nil {
        utils.ServerErrorResponse(w, r, err)
        return
    }
    fmt.Println(userFromGoogle)
    user := sgc.findUserByEmail.Execute(userFromGoogle.Email);
    if user == nil {
        sgc.saveUser.Execute(userFromGoogle.Name, userFromGoogle.Email)
        user = sgc.findUserByEmail.Execute(userFromGoogle.Email);
        if user == nil {
            utils.ServerErrorResponse(w, r, fmt.Errorf("USER REGISTRATION FAILED"))
            return
        }
    }
    frontendURL := fmt.Sprintf("http://localhost:8000/users/authentica?email=%s", user.Email)
	http.Redirect(w, r, frontendURL, http.StatusPermanentRedirect)
}
