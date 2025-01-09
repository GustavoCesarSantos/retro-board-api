package middleware

import (
	"fmt"
	"net/http"

	"github.com/GustavoCesarSantos/retro-board-api/internal/shared/utils"
)

func RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metadataErr := utils.Envelope{
			"file": "recoverPanic.go",
			"func": "recoverPanic.RecoverPanic",
			"line": 0,
		}
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				metadataErr["line"] = 20
				utils.ServerErrorResponse(
					w, 
					r, 
					fmt.Errorf("%s", err), 
					metadataErr,
				)
			}
		}()
		next.ServeHTTP(w, r)
	})
}