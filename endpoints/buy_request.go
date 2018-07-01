package endpoints

import (
	"database/sql"

	"net/http"
	"strconv"

	"github.com/dfang/yuanxin_api/model"
	"github.com/dfang/yuanxin_api/util"
	"github.com/gorilla/mux"
)

// 删除求购
func DeleteBuyRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		userID := GetUIDFromContext(r)

		br, err := model.BuyRequestByID(db, id)
		if br == nil || err != nil {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int    `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: 400,
				Message:    "找不到",
			})
			return
		}

		if br.UserID.Int64 != int64(userID) {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int    `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: 401,
				Message:    "无操作权限",
			})
			return
		}

		err = br.Delete(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{http.StatusInternalServerError, err.Error()})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int    `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: 200,
			Message:    "操作成功",
		})

	})
}
