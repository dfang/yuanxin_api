package endpoints

import (
	"database/sql"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

// 收藏和取消收藏
func LikableEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("NOT IMPLEMENTED"))

		user := r.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)
		// fmt.Println(user.Valid())
		fmt.Fprintf(w, "%v", user)
	})
}
