package middleware

import (
	"GameStats/api"
	"GameStats/internal/tools"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var UnAuthorizedError = errors.New("invalid token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(responseWriter, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			log.Error(err)
			api.InternalErrorHandler(responseWriter, err)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(responseWriter, UnAuthorizedError)
			return
		}

		next.ServeHTTP(responseWriter, r)
	})
}
