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

type UserAuthenticator struct {
    provider interfaces.IUserIdentityApi
}

func NewUserAuthenticator(provider interfaces.IUserIdentityApi) *UserAuthenticator {
    return &UserAuthenticator{
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

func (ua *UserAuthenticator) Authenticate(next http.HandlerFunc) http.HandlerFunc {
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
			metadataErr["line"] = 50
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
			metadataErr["line"] = 62
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
        mappedClaims, ok := claims.Claims.(jwt.MapClaims)
		if !ok ||  !claims.Valid {
			metadataErr["line"] = 68
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
		issuer := mappedClaims["issuer"].(string) 
		if issuer != jwtConfigs.Issuer {
			metadataErr["line"] = 74
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
        authorizedIssuer := contains(jwtConfigs.Audiences, issuer)
		if !authorizedIssuer {
			metadataErr["line"] = 80
			utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
        email := mappedClaims["email"].(string)
        user, userErr := ua.provider.FindByEmail(email)
		if userErr != nil {
			metadataErr["line"] = 87
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
        version := int(mappedClaims["version"].(float64))
        if  version != user.Version {
			metadataErr["line"] = 93
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
        }
		r = utils.ContextSetUser(r, user)
		next.ServeHTTP(w, r)
	})
}

func (ua *UserAuthenticator) AuthenticateWebSocket(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metadataErr := utils.Envelope{
			"file": "userAuthenticator.go",
			"func": "userAuthenticator.AuthenticateWebSocket",
			"line": 0,
		}
		authorization := r.URL.Query().Get("authorization")
		if authorization == "" {
			metadataErr["line"] = 111
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
		authorizationParts := strings.Split(authorization, " ")
		if len(authorizationParts) != 2 || authorizationParts[0] != "Bearer" {
			metadataErr["line"] = 117
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
        jwtConfigs := configs.LoadJwtConfig()
		accessToken := authorizationParts[1]
        claims, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
                return []byte(jwtConfigs.Secret), nil
            }
            return nil, errors.New("INVALID TOKEN VALUE") 
        })
		if err != nil {
			metadataErr["line"] = 130
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
        mappedClaims, ok := claims.Claims.(jwt.MapClaims)
		if !ok ||  !claims.Valid {
			metadataErr["line"] = 136
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
		issuer := mappedClaims["issuer"].(string) 
		if issuer != jwtConfigs.Issuer {
			metadataErr["line"] = 137
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
        authorizedIssuer := contains(jwtConfigs.Audiences, issuer)
		if !authorizedIssuer {
			metadataErr["line"] = 148
			utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
        email := mappedClaims["email"].(string)
        user, userErr := ua.provider.FindByEmail(email)
		if userErr != nil {
			metadataErr["line"] = 155
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
		}
        version := int(mappedClaims["version"].(float64))
        if  version != user.Version {
			metadataErr["line"] = 161
            utils.InvalidAuthenticationTokenResponse(w, r, metadataErr)
			return
        }
		r = utils.ContextSetUser(r, user)
		next.ServeHTTP(w, r)
	})
}