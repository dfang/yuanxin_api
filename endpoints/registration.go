package endpoints

import (
	"database/sql"
	"fmt"
	"net/http"

	null "gopkg.in/guregu/null.v3"

	. "github.com/dfang/yuanxin/model"
	. "github.com/dfang/yuanxin/util"
)

func RegistrationEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("not implemented"))

		for k, v := range r.PostForm {
			fmt.Printf("%s:%s\n", k, v)
		}

		user := User{
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

		// insert into db
		err := user.RegisterUser(db)
		if err != nil {
			// RespondWithError(w, http.StatusServiceUnavailable, err.Error())
			RespondWithJSON(w, http.StatusOK, struct {
				StatusCode string `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: "200",
				Message:    "注册失败",
			})
		} else {
			// RespondWithJSON(w, http.StatusOK, user)
			RespondWithJSON(w, http.StatusOK, struct {
				StatusCode string `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: "200",
				Message:    "注册成功",
			})
		}
	})
}
