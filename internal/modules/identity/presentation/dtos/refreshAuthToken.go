package dtos

type RefreshAuthTokenRequest struct {
	RefreshToken   string       `json:"refreshToken"`
}

type RefreshAuthTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func NewRefreshAuthTokenResponse(accessToken string) *RefreshAuthTokenResponse {
	return &RefreshAuthTokenResponse{
		AccessToken: accessToken,
	}
}