package identity

import (
	"errors"
	"net/http"
	"time"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/application"
	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/presentation/dtos"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type signinUser struct {
    creatAuthToken application.ICreateAuthToken
    findUserByEmail application.IFindUserByEmail
	incrementVersion application.IIncrementVersion
}

func NewSigninUser(
	createAuthToken application.ICreateAuthToken,  
	findUserByEmail application.IFindUserByEmail,
	incrementVersion application.IIncrementVersion,
) *signinUser {
    return &signinUser{
        createAuthToken,
        findUserByEmail,
		incrementVersion,
    }
}

type SigninUserEnvelope struct {
	SigninTokens dtos.SigninUserResponse `json:"signin_tokens"`
}

// SignInUser handles user sign-in, creates access and refresh tokens.
// @Summary Sign in user and generate access and refresh tokens
// @Description This endpoint signs in a user based on their email, generates an access token and a refresh token.
// @Tags Identity
// @Accept json
// @Produce json
// @Param input body dtos.SigninUserRequest true "User email"
// @Success 200 {object} identity.SigninUserEnvelope "Tokens generated successfully"
// @Failure 400 {object} utils.ErrorEnvelope "Invalid request (e.g., missing parameters)"
// @Failure 404 {object} utils.ErrorEnvelope "User not found"
// @Failure 500 {object} utils.ErrorEnvelope "Internal server error"
// @Router /auth/signin [post]
func(su *signinUser) Handle(w http.ResponseWriter, r *http.Request) {
	var input dtos.SigninUserRequest
	readErr := utils.ReadJSON(w, r, &input)
	if readErr != nil {
		utils.BadRequestResponse(w, r, readErr)
		return
	}
	user, findUserErr := su.findUserByEmail.Execute(input.Email)
    if findUserErr != nil {
		switch {
		case errors.Is(findUserErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r)
		default:
            utils.ServerErrorResponse(w, r, findUserErr)
		}
		return
    }
    incrementErr:= su.incrementVersion.Execute(user)
    if incrementErr != nil {
		switch {
		case errors.Is(incrementErr, utils.ErrRecordNotFound):
            utils.NotFoundResponse(w, r)
		default:
            utils.ServerErrorResponse(w, r, incrementErr)
		}
		return
    }
    accessToken, accessTokenErr := su.creatAuthToken.Execute(*user, 15 * time.Minute)
	if accessTokenErr != nil {
		utils.ServerErrorResponse(w, r, accessTokenErr)
	}
    refreshToken, refreshTokenErr := su.creatAuthToken.Execute(*user, 24 * time.Hour * time.Duration(7))
    if refreshTokenErr != nil {
		utils.ServerErrorResponse(w, r, refreshTokenErr)
	}
	response := dtos.NewSigninUserResponse(accessToken, refreshToken)
	data := utils.Envelope{
		"signin_tokens": response,
	}
	writeErr := utils.WriteJSON(w, http.StatusOK, data, nil)
	if writeErr != nil {
		utils.ServerErrorResponse(w, r, writeErr)
	}
}
