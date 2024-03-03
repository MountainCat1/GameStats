package handlers

import (
	"encoding/json"
	"fmt"
	"gamestats-application/api"
	"gamestats-application/internal/tools"
	"gamestats-domain/entities"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateUser handles the creation of a new user.
// @Summary Create a new user
// @Description This API endpoint creates a new user with the provided username and password.
// @Tags users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "User details"
// @Success 200 {object} api.UserDto "User created successfully"
// @Failure 400 {object} api.ErrorResponse "Invalid request format or parameters"
// @Failure 500 {object} api.ErrorResponse "Internal server error"
// @Router /users [post]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if h.UserRepo == nil {
		api.InternalErrorHandler(w, fmt.Errorf("UserRepo is not initialized"))
		return
	}

	var requestBody CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		api.RequestErrorHandler(w, err)
		return
	}

	hashedPassword, err := tools.HashPassword(requestBody.Password)

	user := entities.CreateUser(requestBody.Username, hashedPassword)

	err, user = h.UserRepo.AddUser(user)
	if err != nil {
		api.InternalErrorHandler(w, err)
	}

	var response = api.ToUserDto(user)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w, err)
		return
	}
}
