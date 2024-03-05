package handlers

import (
	"encoding/json"
	"fmt"
	"gamestats-application/api"
	"gamestats-domain/entities"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type CreateMatchRequest struct {
	MatchType   entities.MatchType `json:"matchType"`
	DateStarted time.Time          `json:"dateStarted"`
	DateEnded   time.Time          `json:"dateEnded"`
}

// CreateMatch handles the creation of a new user.
// @Summary Create new match entry
// @Tags matches
// @Accept json
// @Produce json
// @Param request body CreateMatchRequest true "Match details"
// @Success 200 {object} api.MatchDto "Match created successfully"
// @Failure 400 {object} api.ErrorResponse "Invalid request format or parameters"
// @Failure 500 {object} api.ErrorResponse "Internal server error"
// @Security basicAuth
// @Router /matches [post]
func (h *Handler) CreateMatch(w http.ResponseWriter, r *http.Request) {
	if h.MatchRepo == nil {
		api.InternalErrorHandler(w, fmt.Errorf("match is not initialized"))
		return
	}

	var requestBody CreateMatchRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		api.RequestErrorHandler(w, err)
		return
	}

	match := entities.CreateMatch(requestBody.MatchType)

	err, match = h.MatchRepo.AddMatchDetails(match)
	if err != nil {
		api.InternalErrorHandler(w, err)
	}

	var response = api.ToMatchDto(match)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w, err)
		return
	}
}
