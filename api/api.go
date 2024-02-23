package api

import (
	"encoding/json"
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

func writeError(w http.ResponseWriter, message string, code int) {
	response := Error{
		Message:      message,
		ResponseCode: code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, "Something went wrong...", http.StatusInternalServerError)
	}
)
