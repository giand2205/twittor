package middlew

import (
	"net/http"
	"github.com/giand2205/twittor/db"
)

func ValidateDB(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if db.ValidateConnection() == 0 {
			http.Error(writer, "Lose connection to Database", 500)
			return
		}
		next.ServeHTTP(writer, request)
	}
}