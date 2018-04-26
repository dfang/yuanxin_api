package endpoints

import (
	"database/sql"
	"fmt"

	"net/http"
	"strconv"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
	"github.com/gorilla/mux"
)

// 求助列表
func ListHelpRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars["user_id"] == "" {
			str := fmt.Sprintf("参数%s缺失", "user_id")
			w.Write([]byte(str))
			return
		}

	})
}

// 芯片列表
func ListChipsEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		qs := r.URL.Query()

		count, _ := strconv.Atoi(qs.Get("count"))
		start, _ := strconv.Atoi(qs.Get("start"))

		fmt.Println(qs)
		fmt.Println(count)
		fmt.Println(start)

		// var log = logrus.New()
		// log.Out = os.Stdout

		// log.Println(count)
		// log.Println(start)
		// log.WithFields(logrus.Fields{
		// 	"animal": "walrus",
		// 	"size":   10,
		// }).Info("A group of walrus emerges from the ocean")

		if count < 1 {
			count = 10
		}

		if start < 0 {
			start = 0
		}

		chips, err := model.GetChips(db, start, count)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{http.StatusInternalServerError, err.Error()})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int          `json:"status_code"`
			Message    string       `json:"msg"`
			Data       []model.Chip `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       chips,
		})
	})
}

// 用户列表
func ListUsersEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		qs := r.URL.Query()
		count, _ := strconv.Atoi(qs.Get("count"))
		start, _ := strconv.Atoi(qs.Get("start"))

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

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int          `json:"status_code"`
			Message    string       `json:"msg"`
			Data       []model.User `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       users,
		})
	})
}

// 邀请码列表
func ListInvitationsEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		qs := r.URL.Query()
		count, _ := strconv.Atoi(qs.Get("count"))
		start, _ := strconv.Atoi(qs.Get("start"))

		if count < 1 {
			count = 10
		}

		if start < 0 {
			start = 0
		}

		invitations, err := model.GetAllInvitations(db, start, count)
		if err != nil {
			util.RespondWithJSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int                `json:"status_code"`
			Message    string             `json:"msg"`
			Data       []model.Invitation `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       invitations,
		})

	})
}
