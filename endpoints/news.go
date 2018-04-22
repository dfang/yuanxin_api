package endpoints

import (
	"database/sql"
	"net/http"
	"strconv"

	. "github.com/dfang/yuanxin/model"
	. "github.com/dfang/yuanxin/util"
	"github.com/gorilla/mux"
)

func ListNewsItemEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("not implemented"))
		vars := mux.Vars(r)

		count, err := strconv.Atoi(vars["count"])
		start, err := strconv.Atoi(vars["start"])
		t, err := strconv.Atoi(vars["type"])

		if count < 1 {
			count = 10
		}

		if start < 0 {
			start = 0
		}

		news, err := GetNews(db, start, count, NewsItemType(t))
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, news)

	})
}

func GetNewsItemEndpoint(db *sql.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("not implemented"))

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
			return
		}

		item := NewsItem{ID: id}
		if err := item.GetNewsItem(db); err != nil {
			switch err {
			case sql.ErrNoRows:
				RespondWithError(w, http.StatusNotFound, "User not found")
			default:
				RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		RespondWithJSON(w, http.StatusOK, item)
	}
	return http.HandlerFunc(fn)
}
