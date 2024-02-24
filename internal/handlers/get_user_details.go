package handlers

import (
	"GameStats/api"
	"GameStats/internal/tools"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (handler Handler) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	userRepo := *handler.UserRepo

	var username string = r.URL.Query().Get("username")

	loginDetails, err := userRepo.GetUserLoginDetails(username)
	if err != nil {
		if errors.Is(err, tools.NotFoundError) {
			api.NotFoundErrorHandler(w, err)
			return
		}

		api.RequestErrorHandler(w, err)
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
