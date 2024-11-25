package configs


type JwtConfig struct {
	Secret  string
}

func LoadJwtConfig() JwtConfig {
	return JwtConfig{
		Secret:  GetEnv("JWT_SECRET", "teste"),
	}
}
