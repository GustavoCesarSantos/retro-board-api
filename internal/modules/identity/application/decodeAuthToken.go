package application

import (
	"errors"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/configs"
	"github.com/golang-jwt/jwt/v5"
)

type DecodedToken struct {
    Email string
    Version int
}


type IDecodeAuthToken interface {
    Execute(refreshToken string) (*DecodedToken, error)
}

type decodeAuthToken struct {}

func NewDecodeAuthToken() IDecodeAuthToken {
    return &decodeAuthToken{}
}

func (da *decodeAuthToken) Execute(refreshToken string) (*DecodedToken, error) {
    jwtConfigs := configs.LoadJwtConfig()
    claims, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
            return []byte(jwtConfigs.Secret), nil
        }
        return nil, errors.New("INVALID TOKEN VALUE") 
    })
    if err != nil {
        return nil, err
    }
    mappedClaims, ok := claims.Claims.(jwt.MapClaims)
    if !ok ||  !claims.Valid {
        return nil, errors.New("INVALID TOKEN VALUE")
    }
    return &DecodedToken{
        Email: mappedClaims["email"].(string),
        Version: int(mappedClaims["version"].(float64)),
    }, nil
}
