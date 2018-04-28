package endpoints

import (
	"database/sql"
	"net/http"
)

// 收藏和取消收藏
func LikableEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("NOT IMPLEMENTED"))
	})
}
