package repositories

import "gamestats-domain/entities"

type IUserRepo interface {
	GetUserLoginDetails(username string) (*entities.LoginDetails, error)
	AddUser(user *entities.User) (error, *entities.User)
}
