package middleware

import (
	"GameStats.Domain/repositories"
)

type Handler struct {
	UserRepo repositories.IUserDetailsRepo
}
