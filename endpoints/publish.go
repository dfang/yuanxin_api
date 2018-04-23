package endpoints

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
	null "gopkg.in/guregu/null.v3"
)

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

// 发布求助
func PublishHelpRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("not implemented"))

		if r.PostFormValue("user_id") == "" {
			str := fmt.Sprintf("参数%s缺失", "user_id")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("title") == "" {
			str := fmt.Sprintf("参数%s缺失", "title")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("content") == "" {
			str := fmt.Sprintf("参数%s缺失", "content")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("amount") == "" {
			str := fmt.Sprintf("参数%s缺失", "amount")
			w.Write([]byte(str))
			return
		}

		user_id, err := strconv.Atoi(r.PostFormValue("user_id"))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		amount, err := strconv.Atoi(r.PostFormValue("amount"))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		user, err := model.UserByID(db, user_id)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
			// panic(err)
		}

		hr := model.HelpRequest{
			UserID:  null.IntFrom(int64(user.ID)),
			Title:   null.StringFrom(r.PostFormValue("title")),
			Content: null.StringFrom(r.PostFormValue("content")),
			Amount:  null.IntFrom(int64(amount)),
		}

		err = hr.Insert(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode string `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: "220",
				Message:    "发布失败",
			})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode string `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: "200",
			Message:    "发布成功",
		})
		return
	})
}

// 发布求购
func PublishBuyRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("not implemented"))

		if r.PostFormValue("user_id") == "" {
			str := fmt.Sprintf("参数%s缺失", "user_id")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("title") == "" {
			str := fmt.Sprintf("参数%s缺失", "title")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("content") == "" {
			str := fmt.Sprintf("参数%s缺失", "content")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("amount") == "" {
			str := fmt.Sprintf("参数%s缺失", "amount")
			w.Write([]byte(str))
			return
		}

		user_id, err := strconv.Atoi(r.PostFormValue("user_id"))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		amount, err := strconv.Atoi(r.PostFormValue("amount"))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		user, err := model.UserByID(db, user_id)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
			// panic(err)
		}

		br := model.BuyRequest{
			UserID:  null.IntFrom(int64(user.ID)),
			Title:   null.StringFrom(r.PostFormValue("title")),
			Content: null.StringFrom(r.PostFormValue("content")),
			Amount:  null.IntFrom(int64(amount)),
		}

		br.Insert(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode string `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: "220",
				Message:    err.Error(),
			})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode string `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: "200",
			Message:    "发布成功",
		})
		return
	})
}
