package configs


type JwtConfig struct {
    Audiences []string
    Issuer string
	Secret  string
}

func LoadJwtConfig() JwtConfig {
	return JwtConfig{
        Audiences: []string{GetEnv("JWT_AUDIENCES", "teste")},
        Issuer: GetEnv("JWT_ISSUER", "teste"),
		Secret:  GetEnv("JWT_SECRET", "teste"),
	}
}
