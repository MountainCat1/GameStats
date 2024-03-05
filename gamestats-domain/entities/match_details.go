package entities

import (
	"github.com/google/uuid"
	"time"
)

type MatchType int

const (
	Normal MatchType = iota // iota is 0
)

type MatchDetails struct {
	ID          string    `json:"id"`
	DateStarted time.Time `json:"dateStarted"`
	DateEnded   time.Time `json:"dateEnded"`
	MatchType   MatchType `json:"matchType"`
}

func CreateMatch(matchType MatchType, started time.Time, ended time.Time) *MatchDetails {
	return &MatchDetails{
		ID:          uuid.NewString(),
		MatchType:   matchType,
		DateStarted: started,
		DateEnded:   ended,
	}
}
