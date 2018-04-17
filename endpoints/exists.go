package endpoints

import (
	"database/sql"
	"net/http"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
)

// phone, email, phone&captcha
func ExistsEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("not implemented"))
		phone := r.PostFormValue("phone")
		email := r.PostFormValue("email")
		captcha := r.PostFormValue("captcha")

		if captcha != "" && phone != "" {
			captchaExists(w, db, phone, captcha)
		}

		if phone != "" {
			phoneExists(w, db, phone)
		}

		if email != "" {
			emailExists(w, db, email)
		}
	})
}

func captchaExists(w http.ResponseWriter, db *sql.DB, phone, code string) {
	captcha, _ := model.CaptchaBy(db, phone, code)

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

func phoneExists(w http.ResponseWriter, db *sql.DB, phone string) {
	user, err := model.UserByPhone(db, phone)
<<<<<<< HEAD
=======

>>>>>>> develop
	if err != nil {
		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int    `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: 200,
			Message:    "手机号码没有被注册",
		})
	}
<<<<<<< HEAD
=======

>>>>>>> develop
	if user != nil {
		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode string `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: "202",
			Message:    "手机号码已经被注册",
		})
	}
	return
}

func emailExists(w http.ResponseWriter, db *sql.DB, email string) {
	user, err := model.UserByEmail(db, email)
	if err != nil {
		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int    `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: http.StatusOK,
			Message:    "邮箱没有被注册",
		})
	}

	if user != nil {
		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode string `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: "203",
			Message:    "邮箱已经被注册",
		})
	}
	return
}
