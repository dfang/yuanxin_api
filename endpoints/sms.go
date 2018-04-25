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

		CheckRequiredParameter(r, "phone")
		// TODO: validate phone format

		captcha := model.Captcha{
			Phone: null.StringFrom(r.PostFormValue("phone")),
			Code:  null.StringFrom(util.GenCaptcha()),
		}

		fmt.Println(captcha)

		_, err := captcha.InsertOrUpdate(db, captcha.Phone.String, captcha.Code.String)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{500, "发送失败"})
			return
		}

		_, err = util.NewSMSAccount().Send(captcha.Phone.String, captcha.Code.String)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{500, "发送失败"})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "发送成功"})
		return
	})
}

func ValidateSMSEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("not implemented"))
		CheckRequiredParameters(r, "phone", "captcha")
		// TODO: validate phone format

		phone := r.PostFormValue("phone")
		code := r.PostFormValue("captcha")

		captcha, _ := model.CaptchaByPhoneAndCode(db, phone, code)

		if captcha != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "验证码有效"})
		} else {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{204, "验证码无效"})
		}
		return
	})
}
