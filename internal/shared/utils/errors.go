package utils

import (
	"log/slog"
	"net/http"
)

func logError(r *http.Request, err error) {
	var (
		method = r.Method
		url    = r.URL.RequestURI()
	)
	slog.Error(err.Error(), "method", method, "url", url)
}

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	data := Envelope{"error": message}
	err := WriteJSON(w, status, data, nil)
	if err != nil {
		logError(r, err)
		w.WriteHeader(500)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, err)
	message := "The server encountered a problem and could not process your request"
	errorResponse(w, r, http.StatusInternalServerError, message)
}