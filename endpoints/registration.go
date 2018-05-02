package endpoints

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
)

// RegistrationEndpoint 注册
func RegistrationEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CheckRequiredParameters(r, "nickname", "phone", "email", "password")
		// TODO Validate Email
		// TODO Validate Phone
		err := r.ParseForm()
		PanicIfNotNil(err)

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

		user.Pwd = hashAndSalt([]byte(user.Pwd))
		err = user.RegisterUser(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, err.Error()})
			return
		}
		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int        `json:"status_code"`
			Message    string     `json:"msg"`
			Data       model.User `json:"data"`
		}{
			StatusCode: 200,
			Message:    "注册成功",
			Data:       user,
		})
		return
	})
}

// UpdateRegistrationInfo 更改个人信息
func UpdateRegistrationInfo(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CheckRequiredParameters(r, "nickname", "phone", "gender", "avatar")
		// TODO: validate phone

		err := r.ParseForm()
		PanicIfNotNil(err)

		userID := GetUIDFromContext(r)
		user, err := model.UserByID(db, int(userID))
		PanicIfNotNil(err)

		if user == nil {
			panic(&RecordNotFound{
				Error: errors.New("找不到用户"),
			})
		}

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
			Message:    "更新成功",
			Data:       user,
		})
	})
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
