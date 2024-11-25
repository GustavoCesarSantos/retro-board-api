package oauth2

import (
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/configs"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func SetProvider() {
    oauth2Configs := configs.LoadOAuth2Config()
    goth.UseProviders(
		google.New(
			oauth2Configs.Google.ClientID, 
			oauth2Configs.Google.ClientSecret, 
			oauth2Configs.Google.ClientCallbackUrl, 
			"email", 
			"profile",
		),
	)
}
