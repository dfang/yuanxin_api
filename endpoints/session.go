package endpoints

import (
	"database/sql"
	"net/http"

	. "github.com/dfang/yuanxin/model"
	. "github.com/dfang/yuanxin/util"
)

func SessionEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		CheckRequiredParameters(r, "phone", "password")

		phone := r.PostFormValue("phone")
		password := r.PostFormValue("password")

		user, err := SignInUser(db, phone, password)
		if err != nil || user == nil {
			RespondWithJSON(w, http.StatusOK, PayLoadFrom{StatusCode: 205, Message: "手机号码或密码错误"})
			return
		}

		if user != nil {
			RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int    `json:"status_code"`
				Message    string `json:"msg"`
				Data       *User  `json:"data"`
			}{
				StatusCode: 200,
				Message:    "查询成功",
				Data:       user,
			})
			return
		}
	})
}
