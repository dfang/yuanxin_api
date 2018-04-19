package endpoints

import (
	"database/sql"
	"fmt"
	"net/http"

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
