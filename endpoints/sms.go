package endpoints

import (
	"database/sql"
	"fmt"
	"net/http"

	"gopkg.in/guregu/null.v3"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
)

func SendSMSEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("not implemented"))
		phone := r.PostFormValue("phone")
		// w.Write([]byte(phone))

		captcha := model.Captcha{
			Phone: null.StringFrom(phone),
			Code:  null.StringFrom(util.GenCaptcha()),
		}

		fmt.Println(captcha)

		// err := captcha.Insert(db)
		_, err := captcha.InsertOrUpdate(db, captcha.Phone.String, captcha.Code.String)
		// if c != nil {
		// 	c.Update(db)
		// }

		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode string `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: "500",
				Message:    "发送失败",
			})
			return

		}

		_, err = util.NewSMSAccount().Send(captcha.Phone.String, captcha.Code.String)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode string `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: "500",
				Message:    "发送失败",
			})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode string `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: "200",
			Message:    "发送成功",
		})
		return

	})
}

func ValidateSMSEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("not implemented"))
		phone := r.PostFormValue("phone")
		code := r.PostFormValue("captcha")

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
	})
}
