package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/configs"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/integration/interfaces"
	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

type userAuthenticator struct {
    provider interfaces.IUserIdentityApi
}

func NewUserAuthenticator(provider interfaces.IUserIdentityApi) *userAuthenticator {
    return &userAuthenticator{
        provider,
    }
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func (ua *userAuthenticator) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
            utils.InvalidAuthenticationTokenResponse(w, r)
			return
		}
		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
            utils.InvalidAuthenticationTokenResponse(w, r)
			return
		}
        jwtConfigs := configs.LoadJwtConfig()
        claims, err := jwt.Parse(headerParts[1], func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
                return []byte(jwtConfigs.Secret), nil
            }
            return nil, errors.New("INVALID TOKEN VALUE") 
        })
		if err != nil {
            utils.InvalidAuthenticationTokenResponse(w, r)
			return
		}
        mappedClaims, ok := claims.Claims.(jwt.MapClaims)
		if !ok ||  !claims.Valid {
            utils.InvalidAuthenticationTokenResponse(w, r)
			return
		}
		issuer := mappedClaims["issuer"].(string) 
		if issuer != jwtConfigs.Issuer {
            utils.InvalidAuthenticationTokenResponse(w, r)
			return
		}
        authorizedIssuer := contains(jwtConfigs.Audiences, issuer)
		if !authorizedIssuer {
			utils.InvalidAuthenticationTokenResponse(w, r)
			return
		}
        email := mappedClaims["email"].(string)
        user, userErr := ua.provider.FindByEmail(email)
		if userErr != nil {
            utils.InvalidAuthenticationTokenResponse(w, r)
			return
		}
        version := int(mappedClaims["version"].(float64))
        if  version != user.Version {
            utils.InvalidAuthenticationTokenResponse(w, r)
			return
        }
		r = utils.ContextSetUser(r, user)
		next.ServeHTTP(w, r)
	})
}
