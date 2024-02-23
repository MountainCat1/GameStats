package handlers

import (
	"GameStats/api"
	"GameStats/internal/tools"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	var database *tools.DatabaseInterface
	database, err := tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w, err)
		return
	}

	var username string = r.URL.Query().Get("username")

	var loginDetails *tools.LoginDetails
	loginDetails = (*database).GetUserLoginDetails(username)

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
