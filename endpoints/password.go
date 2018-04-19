package endpoints

import (
	"database/sql"
	"net/http"
)

func PasswordEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not implemented"))
	})
}
