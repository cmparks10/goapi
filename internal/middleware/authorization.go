package middleware

import (
	"errors"
	"goapi/api"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid Username or Token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)){

		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var error error

		//If username or token is emptym throw UnAuthError
		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		//Instantiate db
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		
		//Error Check
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDeatils *tools.loginDeatils
		loginDeatils = (*database).GetUserLoginDeatils(username)

		if(loginDeatils == nil || (token != (*loginDeatils).AuthToken)){
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w,UnAuthorizedError)
			return
		}

		next.ServeHTTP(w,r)

	}
}