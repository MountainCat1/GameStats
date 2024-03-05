package api

import (
	"gamestats-domain/entities"
	"gamestats-domain/value_objects"
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
	PlayerStats []*PlayerStatsDto  `json:"playerStats"`
}

func ToMatchDto(matchDetails *entities.MatchDetails) *MatchDto {
	playerStatsDto := make([]*PlayerStatsDto, 0)
	for _, playerStats := range matchDetails.PlayerStats {
		playerStatsDto = append(playerStatsDto, ToPlayerStatsDto(playerStats))
	}

	return &MatchDto{
		ID:          matchDetails.ID,
		DateStarted: matchDetails.DateStarted,
		DateEnded:   matchDetails.DateEnded,
		MatchType:   matchDetails.MatchType,
		PlayerStats: playerStatsDto,
	}
}

type PlayerStatsDto struct {
	PlayerId string `json:"playerId"`
	Kills    int    `json:"kills"`
	Deaths   int    `json:"deaths"`
	Score    int    `json:"score"`
}

func ToPlayerStatsDto(playerStats *value_objects.PlayerStats) *PlayerStatsDto {
	return &PlayerStatsDto{
		PlayerId: playerStats.PlayerId,
		Kills:    playerStats.Kills,
		Deaths:   playerStats.Deaths,
	}
}
