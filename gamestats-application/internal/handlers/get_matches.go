package handlers

import (
	"encoding/json"
	"fmt"
	"gamestats-application/api"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// GetMatches godoc
// @Summary Get match details
// @Description get details of a matches
// @Tags matches
// @Accept  json
// @Produce  json
// @Success 200 {array} api.MatchDto
// @Success 401 {object} api.ErrorResponse "Unauthorized access"
// @Security basicAuth
// @Router /matches [get]
func (h *Handler) GetMatches(w http.ResponseWriter, r *http.Request) {
	if h.MatchRepo == nil {
		api.InternalErrorHandler(w, fmt.Errorf("match is not initialized"))
		return
	}

	// Call GetUserLoginDetails directly on the pointer
	err, matches := h.MatchRepo.GetAll()
	if err != nil {
		api.RequestErrorHandler(w, err)
		return
	}

	var response []*api.MatchDto
	for _, match := range matches {
		matchDto := api.ToMatchDto(match)
		response = append(response, matchDto)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w, err)
		return
	}
}
