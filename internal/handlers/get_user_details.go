package handlers

import (
	"GameStats/api"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (handler Handler) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	userRepo := *handler.UserRepo

	var username string = r.URL.Query().Get("username")

	loginDetails, err := userRepo.GetUserLoginDetails(username)
	if err != nil {
		api.RequestErrorHandler(w, err)
		return
	}

	if loginDetails == nil {
		api.NotFoundErrorHandler(w, api.NotFoundError)
		return
	}

	var response = api.UserDetailsResponse{
		Username:        (*loginDetails).Username,
		FavouriteNumber: (*loginDetails).FavouriteNumber,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w, err)
		return
	}
}
