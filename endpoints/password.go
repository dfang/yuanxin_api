package endpoints

import (
	"database/sql"
	"net/http"

	"github.com/dfang/yuanxin_api/model"
	"github.com/dfang/yuanxin_api/util"
)

func PasswordEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CheckRequiredParameters(r, "phone", "code", "password")
		r.ParseForm()

		// userID := GetUIDFromContext(r)
		_, err := model.CaptchaByPhoneAndCode(db, r.PostFormValue("phone"), r.PostFormValue("code"))
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int    `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: 208,
				Message:    "手机号或者验证码不对",
			})
			return
		}

		u2, _ := model.UserByPhone(db, r.PostFormValue("phone"))
		if u2 != nil {
			// update password
			u2.Pwd = hashAndSalt([]byte(r.PostFormValue("password")))
			err := u2.UpdatePassword(db)
			if err != nil {
				util.RespondWithJSON(w, http.StatusOK, struct {
					StatusCode int    `json:"status_code"`
					Message    string `json:"msg"`
				}{
					StatusCode: 208,
					Message:    "密码修改失败",
				})
				return
			}

			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int    `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: 200,
				Message:    "密码修改成功",
			})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int    `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: 208,
			Message:    "密码修改失败",
		})
		return
	})
}
