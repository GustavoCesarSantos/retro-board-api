package identity

import (
	"net/http"

	"github.com/markbates/goth/gothic"
)

type signinWithGoogle struct {}

func NewSigninWithGoogle() *signinWithGoogle {
    return &signinWithGoogle{}
}

func(sg *signinWithGoogle) Handle(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query()
    q.Add("provider", "google")
    r.URL.RawQuery = q.Encode()
    gothic.BeginAuthHandler(w, r)
}
