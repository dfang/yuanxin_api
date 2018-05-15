package endpoints

import (
	"database/sql"
	"net/http"

	"github.com/dfang/yuanxin_api/model"
	"github.com/dfang/yuanxin_api/util"
)

// phone, email, phone&captcha
func ExistsEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("not implemented"))
		phone := r.PostFormValue("phone")
		email := r.PostFormValue("email")
		captcha := r.PostFormValue("captcha")

		if captcha != "" && phone != "" {
			CaptchaExists(w, db, phone, captcha)
		}

		if phone != "" {
			PhoneExists(w, db, phone)
		}

		if email != "" {
			EmailExists(w, db, email)
		}
	})
}

func CaptchaExists(w http.ResponseWriter, db *sql.DB, phone, code string) {
	captcha, _ := model.CaptchaByPhoneAndCode(db, phone, code)
	if captcha != nil {
		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int    `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: 200,
			Message:    "验证码有效",
		})
	} else {
		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int    `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: 204,
			Message:    "验证码无效",
		})
	}
	return
}

func PhoneExists(w http.ResponseWriter, db *sql.DB, phone string) {
	user, _ := model.UserByPhone(db, phone)
	if user != nil {
		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode string `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: "202",
			Message:    "手机号码已经被注册",
		})
	} else {
		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int    `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: 200,
			Message:    "手机号码没有被注册",
		})
	}
	return
}

func EmailExists(w http.ResponseWriter, db *sql.DB, email string) {
	user, _ := model.UserByEmail(db, email)
	if user != nil {
		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode string `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: "203",
			Message:    "邮箱已经被注册",
		})
	} else {
		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int    `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: http.StatusOK,
			Message:    "邮箱没有被注册",
		})
	}
	return
}
