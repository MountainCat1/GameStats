package middleware

import (
	"GameStats.Api/api"
	"errors"
	"net/http"
)

var UnAuthorizedError = errors.New("invalid token")

func (h *Handler) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, r *http.Request) {
		var username, password, ok = r.BasicAuth()

		if !ok {
			api.UnauthorizedErrorHandler(responseWriter, UnAuthorizedError)
			return
		}

		if username == "" || password == "" {
			api.UnauthorizedErrorHandler(responseWriter, UnAuthorizedError)
			return
		}

		var loginDetails, err = h.UserRepo.GetUserLoginDetails(username)
		if err != nil {
			api.InternalErrorHandler(responseWriter, err)
			return
		}

		if loginDetails == nil || (password != (*loginDetails).AuthToken) {
			api.UnauthorizedErrorHandler(responseWriter, UnAuthorizedError)
			return
		}

		next.ServeHTTP(responseWriter, r)
	})
}
