package api

import (
	"gamestats-domain/entities"
	"time"
)

type UserDto struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func ToUserDto(user *entities.User) *UserDto {
	return &UserDto{
		Id:       user.Id,
		Username: user.Username,
	}
}

type MatchDto struct {
	ID          string             `json:"id"`
	DateStarted time.Time          `json:"dateStarted"`
	DateEnded   time.Time          `json:"dateEnded"`
	MatchType   entities.MatchType `json:"matchType"`
}

func ToMatchDto(matchDetails *entities.MatchDetails) *MatchDto {
	return &MatchDto{
		ID:          matchDetails.ID,
		DateStarted: matchDetails.DateStarted,
		DateEnded:   matchDetails.DateEnded,
		MatchType:   matchDetails.MatchType,
	}
}
