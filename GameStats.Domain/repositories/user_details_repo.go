package repositories

import "GameStats.Domain/entities"

type IUserDetailsRepo interface {
	GetUserLoginDetails(username string) (*entities.LoginDetails, error)
}
