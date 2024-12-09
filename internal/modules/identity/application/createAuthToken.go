package application

import (
	"time"

	"github.com/GustavoCesarSantos/retro-board-api/internal/modules/identity/domain"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/configs"
	"github.com/golang-jwt/jwt/v5"
)

type ICreateAuthToken interface {
    Execute(user domain.User, expTime time.Duration) (string, error)
}

type createAuthToken struct {}

func NewCreateAuthToken() ICreateAuthToken {
    return &createAuthToken{}
}

func (su *createAuthToken) Execute(user domain.User, expTime time.Duration) (string, error) {
    jwtConfigs := configs.LoadJwtConfig()
    claims := jwt.MapClaims{
        "issuer": jwtConfigs.Issuer,
        "name": user.Name,
        "email": user.Email,
        "version": user.Version,
        "exp": time.Now().Add(expTime).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenSigned, err := token.SignedString([]byte(jwtConfigs.Secret))
    if err != nil {
        return "", err
    }
    return tokenSigned, nil
}
