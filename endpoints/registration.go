package endpoints

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	null "gopkg.in/guregu/null.v3"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
)

func RegistrationEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.PostForm {
			fmt.Printf("%s:%s\n", k, v)
		}

		user := model.User{
			Nickname: null.StringFrom(r.PostFormValue("nickname")),
			Phone:    null.StringFrom(r.PostFormValue("phone")),
			Email:    null.StringFrom(r.PostFormValue("email")),
			Pwd:      null.StringFrom(r.PostFormValue("password")),
		}

		// use gorilla scheme to decode form values to user, that's called data binding in rails/asp.net mvc
		// user := new(User)
		// decoder := SchemaDecoder
		// decoder :=
		// schema.NewDecoder().Decode(&user, r.PostForm)
		// decoder.Decode(user, r.PostForm)

		fmt.Printf("%+v\n", user)

		// EmailExists(w, db, user.Email.String)
		u, _ := model.UserByEmail(db, user.Email.String)
		if u != nil {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode string `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: "203",
				Message:    "邮箱已经被注册",
			})
			return
		}

		u2, _ := model.UserByPhone(db, user.Phone.String)
		if u2 != nil {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode string `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: "202",
				Message:    "手机号码已经被注册",
			})
			return
		}

		// PhoneExists(w, db, user.Phone.String)

		// insert into db
		err := user.RegisterUser(db)
		if err != nil {
			// RespondWithError(w, http.StatusServiceUnavailable, err.Error())
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode string `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: "200",
				Message:    "注册失败",
			})
			return
		} else {
			// RespondWithJSON(w, http.StatusOK, user)
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode string `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: "200",
				Message:    "注册成功",
			})
			return
		}
	})
}

func UpdateRegistrationInfo(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("not implemented"))
		user_id, err := strconv.Atoi(r.PostFormValue("user_id"))
		if err != nil {
			log.Fatal(err)
			w.Write([]byte(err.Error()))
			return
		}

		user, err := model.UserByID(db, user_id)
		if err != nil {
			log.Fatal(err)
			w.Write([]byte("数据库查询出错"))
			return
		}

		if user == nil {
			w.Write([]byte("找不到此用户"))
		}

		// gender, _ := strconv.Atoi(r.PostFormValue("gender"))

		user.Gender = null.IntFrom(1)
		user.Phone = null.StringFrom(r.PostFormValue("phone"))
		user.Avatar = null.StringFrom(r.PostFormValue("avatar"))
		user.Nickname = null.StringFrom(r.PostFormValue("nickname"))
		user.Biography = null.StringFrom(r.PostFormValue("biography"))

		err = user.UpdateRegistrationInfo(db)
		if err != nil {
			w.Write([]byte("内部出错"))
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode string      `json:"status_code"`
			Message    string      `json:"msg"`
			User       *model.User `json:"user"`
		}{
			StatusCode: "200",
			Message:    "更新成功",
			User:       user,
		})
		return
	})
}
