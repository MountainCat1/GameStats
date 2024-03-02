package middleware

import (
	"gamestats-domain"
)

type Handler struct {
	UserRepo  domain.UserRepo
	MatchRepo domain.MatchRepo
}
