package middleware

import (
	"gamestats-domain/repositories"
)

type Handler struct {
	UserRepo repositories.IUserRepo
}
