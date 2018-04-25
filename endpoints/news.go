package endpoints

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
	"github.com/gorilla/mux"
)

func ListNewsItemEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("not implemented"))
		vars := mux.Vars(r)

		count, _ := strconv.Atoi(vars["count"])
		start, _ := strconv.Atoi(vars["start"])
		t, _ := strconv.Atoi(vars["type"])

		if count < 1 {
			count = 10
		}

		if start < 0 {
			start = 0
		}

		news, err := model.GetNews(db, start, count, model.NewsItemType(t))
		if err != nil {
			util.RespondWithJSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int              `json:"status_code"`
			Message    string           `json:"msg"`
			Data       []model.NewsItem `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       news,
		})
	})
}

func GetNewsItemEndpoint(db *sql.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("not implemented"))

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			util.RespondWithJSON(w, http.StatusBadRequest, "Invalid ID")
			return
		}

		item := model.NewsItem{ID: id}
		if err := item.GetNewsItem(db); err != nil {
			switch err {
			case sql.ErrNoRows:
				// util.RespondWithJSON(w, http.StatusNotFound, "NewsItem not found")
				util.RespondWithJSON(w, http.StatusOK, struct {
					StatusCode int         `json:"status_code"`
					Message    string      `json:"msg"`
					Data       interface{} `json:"data"`
				}{
					StatusCode: 200,
					Message:    "查询成功",
					Data:       struct{}{},
				})
			default:
				util.RespondWithJSON(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int            `json:"status_code"`
			Message    string         `json:"msg"`
			Data       model.NewsItem `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       item,
		})
	}
	return http.HandlerFunc(fn)
}
