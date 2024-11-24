package configs

import "fmt"

type OAuth2Config struct {
	Google struct {
		ClientID          string
		ClientSecret      string
		ClientCallbackUrl string
	}
}

func LoadOAuth2Config() OAuth2Config {
    serverConfig := LoadServerConfig()
	return OAuth2Config{
		Google: struct {
			ClientID          string
			ClientSecret      string
			ClientCallbackUrl string
		}{
            ClientID: GetEnv("GOOGLE_CLIENT_ID", "teste"), 
            ClientSecret: GetEnv("GOOGLE_CLIENT_SECRET", "teste"),
            ClientCallbackUrl: GetEnv("GOOGLE_CLIENT_CALLBACK_URL", fmt.Sprintf("http://localhost:%d/v1/auth/signin/google/callback", serverConfig.Port)),
		},
	}
}
