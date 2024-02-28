package repositories

import "gamestats-domain/entities"

type IUserDetailsRepo interface {
	GetUserLoginDetails(username string) (*entities.LoginDetails, error)
	AddUser(user *entities.User) (error, *entities.User)
}
