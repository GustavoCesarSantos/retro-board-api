package utils

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

var (
	ErrRecordNotFound = errors.New("RECORD NOT FOUND")
	ErrEditConflict   = errors.New("EDIT CONFLICT")
)

type ErrorEnvelope struct {
	Error string `json:"error" example:"error message"`
}

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

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func InvalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("WWW-Authenticate", "Bearer")
	message := "invalid or missing authentication token"
	errorResponse(w, r, http.StatusUnauthorized, message)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "The requeted resource could not be found"
	errorResponse(w, r, http.StatusNotFound, message)
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, err)
	message := "The server encountered a problem and could not process your request"
	errorResponse(w, r, http.StatusInternalServerError, message)
}
