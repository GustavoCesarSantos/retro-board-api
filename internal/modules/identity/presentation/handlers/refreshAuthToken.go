package identity

import (
	"errors"
	"net/http"
	"time"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/presentation/dtos"
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

type RefreshAuthTokenEnvelope struct {
	RefreshedToken string `json:"refreshed_token"`
}

// RefreshAuthToken handles the refresh of access tokens using the refresh token.
// @Summary Refresh access token using the provided refresh token
// @Description This endpoint accepts a refresh token and returns a new access token if the refresh token is valid.
// @Tags Identity
// @Accept json
// @Produce json
// @Param input body dtos.RefreshAuthTokenRequest true "Refresh token"
// @Success 200 {object} identity.RefreshAuthTokenEnvelope "Access token refreshed successfully"
// @Failure 400 {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters or invalid refresh token or invalid credentials)"
// @Failure 404 {object} utils.ErrorEnvelope "User not found"
// @Failure 500 {object} utils.ErrorEnvelope "Internal server error"
// @Router /auth/refresh-token [post]
func(rt *refreshAuthToken) Handle(w http.ResponseWriter, r *http.Request) {
	metadataErr := utils.Envelope{
		"file": "refreshAuthToken.go",
		"func": "refreshAuthToken.Handle",
		"line": 0,
	}
	var input dtos.RefreshAuthTokenRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr, metadataErr)
		return
	}
	decodedToken, decodedErr := rt.decodeAuthToken.Execute(input.RefreshToken)
	if decodedErr != nil {
		utils.BadRequestResponse(w, r, decodedErr, metadataErr)
		return
	}
	user, findUserErr := rt.findUserByEmail.Execute(decodedToken.Email)
    if findUserErr != nil {
		switch {
		case errors.Is(findUserErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r, metadataErr)
		default:
            utils.ServerErrorResponse(w, r, findUserErr, metadataErr)
		}
		return
    }
	if user.Version != decodedToken.Version {
		utils.BadRequestResponse(w, r, errors.New("INVALID CREDENTIALS"), metadataErr)
		return
	}
	accessToken, accessTokenErr := rt.createAuthToken.Execute(*user, 15 * time.Minute)
	if accessTokenErr != nil {
		utils.ServerErrorResponse(w, r, accessTokenErr, metadataErr)
	}
	response := dtos.NewRefreshAuthTokenResponse(accessToken)
	data := utils.Envelope{
		"refreshed_token": response.AccessToken,
	}
	writeErr := utils.WriteJSON(w, http.StatusOK, data, nil)
	if writeErr != nil {
		utils.ServerErrorResponse(w, r, writeErr, metadataErr)
	}
}
