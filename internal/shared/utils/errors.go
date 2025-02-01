package utils

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

var (
	ErrMissingOrInvalidLimitQueryParam = errors.New("MISSING OR INVALID LIMIT QUERY PARAM")
	ErrInvalidLimitQueryParam = errors.New("INVALID LAST ID QUERY PARAM")
	ErrBoardNotInTeam = errors.New("BOARD DOES NOT BELONG TO THE SPECIFIED TEAM")
	ErrCardNotInColumn = errors.New("CARD DOES NOT BELONG TO THE SPECIFIED COLUMN")
	ErrColumnNotInBoard = errors.New("COLUMN DOES NOT BELONG TO THE SPECIFIED BOARD")
	ErrEditConflict = errors.New("EDIT CONFLICT")
	ErrOptionNotInPoll = errors.New("OPTION DOES NOT BELONG TO THE SPECIFIED POLL")
	ErrPollNotInTeam = errors.New("POLL DOES NOT BELONG TO THE SPECIFIED TEAM")
	ErrRecordNotFound = errors.New("RECORD NOT FOUND")
	ErrUserNoEditPermission = errors.New("USER DOES NOT HAVE EDIT PERMISSION")
	ErrUserNotInTeam = errors.New("USER DOES NOT BELONG TO THE SPECIFIED TEAM")
)

type ErrorEnvelope struct {
	Error string `json:"error" example:"error message"`
}

func logError(r *http.Request, err error, metadataErr Envelope) {
	var (
		method = r.Method
		url    = r.URL.RequestURI()
	)
    slog.Error(err.Error(), "method", method, "url", url, "meta", fmt.Sprintf("%s", metadataErr))
}

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message any, metadataErr Envelope) {
	data := Envelope{"error": message}
	err := WriteJSON(w, status, data, nil)
	if err != nil {
        logError(r, err, Envelope{"file": "errors.go", "func": "errorResponse", "line": 39})
		w.WriteHeader(500)
	}
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error, metadataErr Envelope) {
	errorResponse(w, r, http.StatusBadRequest, err.Error(), metadataErr)
}

func ForbiddenResponse(w http.ResponseWriter, r *http.Request, err error, metadataErr Envelope) {
	message := "Forbidden Access" 
	if err != nil {
		message = err.Error()
	}
	errorResponse(w, r, http.StatusForbidden, message, metadataErr)

}

func InvalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request, metadataErr Envelope) {
    w.Header().Set("WWW-Authenticate", "Bearer")
	message := "invalid or missing authentication token"
	errorResponse(w, r, http.StatusUnauthorized, message, metadataErr)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request, metadataErr Envelope) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	errorResponse(w, r, http.StatusMethodNotAllowed, message, metadataErr)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request, metadataErr Envelope) {
	message := "The requeted resource could not be found"
	errorResponse(w, r, http.StatusNotFound, message, metadataErr)
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error, metadataErr Envelope) {
	logError(r, err, metadataErr)
	message := "The server encountered a problem and could not process your request"
	errorResponse(w, r, http.StatusInternalServerError, message, metadataErr)
}
