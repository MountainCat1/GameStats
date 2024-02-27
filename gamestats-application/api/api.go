package api

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UserDetailsResponse struct {
	Username        string `json:"username"`
	FavouriteNumber int    `json:"favouriteNumber"`
}

type Error struct {
	Message      string `json:"message"`
	ResponseCode int    `json:"responseCode"`
}

func (e Error) Error() string {
	return e.Message
}

func writeError(w http.ResponseWriter, message string, code int) {
	response := Error{
		Message:      message,
		ResponseCode: code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}

var RequestErrorHandler = func(w http.ResponseWriter, err error) {
	writeError(w, err.Error(), http.StatusBadRequest)
}

var UnauthorizedErrorHandler = func(w http.ResponseWriter, err error) {
	writeError(w, err.Error(), http.StatusUnauthorized)
}

var InternalErrorHandler = func(w http.ResponseWriter, err error) {
	log.Error(err.Error())
	writeError(w, "Something went wrong...", http.StatusInternalServerError)
}

var NotFoundErrorHandler = func(w http.ResponseWriter, err error) {
	writeError(w, err.Error(), http.StatusNotFound)
}

var NotFoundError = errors.New("Not Found")
