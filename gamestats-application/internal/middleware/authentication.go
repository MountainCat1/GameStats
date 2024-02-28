package middleware

import (
	"gamestats-application/api"
	"gamestats-application/internal/tools"
	"net/http"
)

func (h *Handler) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, r *http.Request) {
		var username, password, ok = r.BasicAuth()
		if !ok {
			api.UnauthorizedErrorHandler(responseWriter, api.UnAuthorizedError)
			return
		}

		if username == "" || password == "" {
			api.UnauthorizedErrorHandler(responseWriter, api.UnAuthorizedError)
			return
		}

		loginDetails, err := h.UserRepo.GetUserLoginDetails(username)
		if err != nil {
			api.InternalErrorHandler(responseWriter, err)
			return
		}

		if tools.ComparePassword(password, (*loginDetails).HashedPassword) != nil {
			api.UnauthorizedErrorHandler(responseWriter, api.UnAuthorizedError)
			return
		}

		next.ServeHTTP(responseWriter, r)
	})
}
