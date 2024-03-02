package domain

import "gamestats-domain/entities"

type MatchRepo interface {
	GetMatchDetails(matchID string) (*entities.MatchDetails, error)
	AddMatchDetails(matchDetails *entities.MatchDetails) (error, *entities.MatchDetails)
}

type UserRepo interface {
	GetUserLoginDetails(username string) (*entities.LoginDetails, error)
	GetUserByUsername(username string) (*entities.User, error)
	AddUser(user *entities.User) (error, *entities.User)
}
