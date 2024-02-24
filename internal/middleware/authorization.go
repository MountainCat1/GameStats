package middleware

import (
	"GameStats/api"
	"errors"
	"net/http"
)

var UnAuthorizedError = errors.New("invalid token")

func (handler *Handler) Authorization(next http.Handler) http.Handler {
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

		var userRepo = handler.UserRepo
		var loginDetails, err = userRepo.GetUserLoginDetails(username)
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
