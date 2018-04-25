package endpoints

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
)

// 注册
func RegistrationEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer RecoverEndpoint(w)

		CheckRequiredParameters(r, "nickname", "phone", "email", "password")
		// TODO Validate Email
		// TODO Validate Phone

		var user model.User
		// use gorilla scheme to decode form values to user, that's called data binding in rails/asp.net mvc
		if err := util.SchemaDecoder.Decode(&user, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}

		u, _ := model.UserByEmail(db, user.Email.String)
		if u != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{StatusCode: 203, Message: "邮箱已经被注册"})
			return
		}

		u2, _ := model.UserByPhone(db, user.Phone.String)
		if u2 != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{StatusCode: 202, Message: "手机号码已经被注册"})
			return
		}

		err := user.RegisterUser(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "注册失败"})
			return
		} else {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int        `json:"status_code"`
				Message    string     `json:"msg"`
				Data       model.User `json:"data"`
			}{
				StatusCode: 200,
				Message:    "查询成功",
				Data:       user,
			})
			return
		}
	})
}

func RegistrationHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before")
		defer log.Println("After")
	})
}

// 更改个人信息
func UpdateRegistrationInfo(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer RecoverEndpoint(w)

		CheckRequiredParameter(r, "user_id")
		user_id := ParseParameterToInt(r, "user_id")

		user, err := model.UserByID(db, user_id)
		PanicIfNotNil(err)

		if user == nil {
			panic(&RecordNotFound{
				Error: errors.New("找不到用户"),
			})
		}

		r.PostForm.Del("user_id")
		if err = util.SchemaDecoder.Decode(user, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}

		err = user.UpdateRegistrationInfo(db)
		PanicIfNotNil(err)

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int         `json:"status_code"`
			Message    string      `json:"msg"`
			Data       *model.User `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       user,
		})
	})
}
