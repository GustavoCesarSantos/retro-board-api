package dtos

type SigninUserRequest struct {
	SigninToken string `json:"signin_token" example:"asdfasdfasd"`
}

type SigninUserResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewSigninUserResponse(accessToken string, refreshToken string) *SigninUserResponse {
	return &SigninUserResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}
}
