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
		metadataErr := utils.Envelope{
			"file": "userAuthenticator.go",
			"func": "userAuthenticator.Authenticate",
			"line": 0,
		}
		w.Header().Add("Vary", "Authorization")
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			metadataErr["line"] = 44
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			metadataErr["line"] = 49
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
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
			metadataErr["line"] = 60
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
        mappedClaims, ok := claims.Claims.(jwt.MapClaims)
		if !ok ||  !claims.Valid {
			metadataErr["line"] = 65
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
		issuer := mappedClaims["issuer"].(string) 
		if issuer != jwtConfigs.Issuer {
			metadataErr["line"] = 70
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
        authorizedIssuer := contains(jwtConfigs.Audiences, issuer)
		if !authorizedIssuer {
			metadataErr["line"] = 75
			utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
        email := mappedClaims["email"].(string)
        user, userErr := ua.provider.FindByEmail(email)
		if userErr != nil {
			metadataErr["line"] = 81
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
        version := int(mappedClaims["version"].(float64))
        if  version != user.Version {
			metadataErr["line"] = 86
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
        }
		r = utils.ContextSetUser(r, user)
		next.ServeHTTP(w, r)
	})
}
