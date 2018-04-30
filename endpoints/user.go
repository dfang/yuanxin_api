package endpoints

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"

	"github.com/gorilla/mux"
)

// GetUserEndpoint 获取用户
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

type UserProfileResult struct {
	User         *model.User         `json:"user"`
	Chips        []model.Chip        `json:"chips"`
	BuyRequests  []model.BuyRequest  `json:"buy_requests"`
	HelpRequests []model.HelpRequest `json:"help_requests"`
}

// GetUserProfileEndpoint 用户资料页
// 用户资料页  查看其发布的芯片发布数、求购、求助
func GetUserProfileEndpoint(db *sql.DB) http.HandlerFunc {
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

		chips, err := model.ChipsByUserID(db, user.ID, 0, 10)
		buyRequests, err := model.BuyRequestsByUserID(db, user.ID, 0, 10)
		helpRequests, err := model.HelpRequestsByUserID(db, user.ID, 0, 10)

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int               `json:"status_code"`
			Message    string            `json:"msg"`
			Data       UserProfileResult `json:"data"`
		}{
			StatusCode: 200,
			Message:    "用户资料页",
			Data: UserProfileResult{
				User:         user,
				Chips:        chips,
				BuyRequests:  buyRequests,
				HelpRequests: helpRequests,
			},
		})
	})
}
