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
