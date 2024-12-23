package identity

import (
	"net/http"

	"github.com/markbates/goth/gothic"
)

type signinWithGoogle struct {}

func NewSigninWithGoogle() *signinWithGoogle {
    return &signinWithGoogle{}
}

// Handle starts the OAuth2.0 login flow with Google.
// @Summary Initiates OAuth2.0 sign-in with Google
// @Description This endpoint starts the OAuth2.0 sign-in flow with Google. 
//              The user will be redirected to Google for authentication.
// @Tags Identity
// @Produce json
// @Success 200 {string} string "Redirecting to Google for authentication"
// @Failure 500 {object} utils.ErrorEnvelope "Internal server error"
// @Router /identity/signin/google [get]
func(sg *signinWithGoogle) Handle(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query()
    q.Add("provider", "google")
    r.URL.RawQuery = q.Encode()
    gothic.BeginAuthHandler(w, r)
}
