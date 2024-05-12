package middleware

import (
	"errors"
	"net/http"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid Username or Token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)){
		
	}
}