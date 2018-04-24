package endpoints

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"

	"github.com/gorilla/mux"
)

// 获取用户
func GetUserEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		user, err := model.UserByID(db, id)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int    `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: 400,
				Message:    err.Error(),
			})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int         `json:"status_code"`
			Message    string      `json:"msg"`
			User       *model.User `json:"user"`
		}{
			StatusCode: 200,
			Message:    "用户详情",
			User:       user,
		})
	})
}
func GetUsersEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		count, _ := strconv.Atoi(vars["count"])
		start, _ := strconv.Atoi(vars["start"])

		if count < 1 {
			count = 10
		}

		if start < 0 {
			start = 0
		}

		users, err := model.GetAllUsers(db, start, count)
		if err != nil {
			util.RespondWithJSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		util.RespondWithJSON(w, http.StatusOK, users)

	})
}
