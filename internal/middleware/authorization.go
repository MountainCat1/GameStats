package middleware

//
//import (
//	"GameStats/api"
//	"errors"
//	log "github.com/sirupsen/logrus"
//	"net/http"
//)
//
//var UnAuthorizedError = errors.New("invalid token")
//
//func (handler MiddlewareHandler) Authorization(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(responseWriter http.ResponseWriter, r *http.Request) {
//		var username string = r.URL.Query().Get("username")
//		var token string = r.Header.Get("Authorization")
//
//		if username == "" || token == "" {
//			log.Error(UnAuthorizedError)
//			api.RequestErrorHandler(responseWriter, UnAuthorizedError)
//			return
//		}
//
//		var loginDetails = handler.DB.
//		loginDetails = (*database).GetUserLoginDetails(username)
//
//		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
//			log.Error(UnAuthorizedError)
//			api.RequestErrorHandler(responseWriter, UnAuthorizedError)
//			return
//		}
//
//		next.ServeHTTP(responseWriter, r)
//	})
//}
