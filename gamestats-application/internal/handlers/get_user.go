package handlers

import (
	"encoding/json"
	"fmt"
	"gamestats-application/api"
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
// @Success 200 {object} api.UserDto
// @Security basicAuth
// @Router /users [get]
func (h *Handler) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	// Ensure that h.UserRepo is a pointer and check it is not nil
	if h.UserRepo == nil {
		api.InternalErrorHandler(w, fmt.Errorf("UserRepo is not initialized"))
		return
	}

	// No need to dereference h.UserRepo
	var username string = r.URL.Query().Get("username")

	// Call GetUserLoginDetails directly on the pointer
	user, err := h.UserRepo.GetUserByUsername(username)
	if err != nil {
		api.RequestErrorHandler(w, err)
		return
	}

	if user == nil {
		api.NotFoundErrorHandler(w, api.NotFoundError)
		return
	}

	// No need to use the dereference operator as user is already a pointer
	var response = api.ToUserDto(user)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w, err)
		return
	}
}
