package handlers

import (
	"GameStats.Api/api"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// GetUserDetails godoc
// @Summary Get user details
// @Description get details of a user by username
// @Tags users
// @Accept  json
// @Produce  json
// @Param username query string true "Username"
// @Success 200 {object} api.UserDetailsResponse
// @Router /test/super [get]
func (h Handler) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	// Ensure that h.UserRepo is a pointer and check it is not nil
	if h.UserRepo == nil {
		api.InternalErrorHandler(w, fmt.Errorf("UserRepo is not initialized"))
		return
	}

	// No need to dereference h.UserRepo
	var username string = r.URL.Query().Get("username")

	// Call GetUserLoginDetails directly on the pointer
	loginDetails, err := h.UserRepo.GetUserLoginDetails(username)
	if err != nil {
		api.RequestErrorHandler(w, err)
		return
	}

	if loginDetails == nil {
		api.NotFoundErrorHandler(w, api.NotFoundError)
		return
	}

	// No need to use the dereference operator as loginDetails is already a pointer
	var response = api.UserDetailsResponse{
		Username:        loginDetails.Username,
		FavouriteNumber: loginDetails.FavouriteNumber,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w, err)
		return
	}
}
