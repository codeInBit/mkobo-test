package middlewares

import (
	"errors"
	"net/http"

	"github.com/codeInBit/mkobo-test/auth"
	"github.com/codeInBit/mkobo-test/utilities"
)

// CommonMiddleware - Set content-type
func CommonMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	}
}

//AuthenticationMiddleware - Prevents unauthorized access
func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			utilities.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"), "Invalid Token")
			return
		}
		next(w, r)
	}
}
