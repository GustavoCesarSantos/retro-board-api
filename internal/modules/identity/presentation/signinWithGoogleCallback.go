package identity

import (
	"fmt"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
	"github.com/markbates/goth/gothic"
)

type signinWithGoogleCallback struct {}

func NewSigninWithGoogleCallback() *signinWithGoogleCallback {
    return &signinWithGoogleCallback{}
}

func(sgc *signinWithGoogleCallback) Handle(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query()
    q.Add("provider", "google")
    r.URL.RawQuery = q.Encode()
    user, err := gothic.CompleteUserAuth(w, r)
    if err != nil {
        fmt.Println("ERRO")
        fmt.Println(err.Error())
        utils.ServerErrorResponse(w, r, err)
        return
    }
    fmt.Println(user)
    writeErr := utils.WriteJSON(w, http.StatusOK, utils.Envelope{"status": "success"}, nil)
	if writeErr != nil {
		utils.ServerErrorResponse(w, r, writeErr)
	}
}
