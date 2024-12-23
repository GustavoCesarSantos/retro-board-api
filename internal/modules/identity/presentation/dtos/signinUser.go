package dtos

type SigninUserRequest struct {
	Email string `json:"email" example:"teste@teste.com"`
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