package endpoints

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dfang/yuanxin_api/model"
	"github.com/dfang/yuanxin_api/util"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	null "gopkg.in/guregu/null.v3"

	"github.com/gorilla/mux"
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
		user.Role = null.IntFrom(1)

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

// 修改密码
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
			StatusCode: 207,
			Message:    "手机不存在",
		})
		return
	})
}

func SessionEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		CheckRequiredParameters(r, "password")
		// TODO:  Validate phone format

		phone := r.PostFormValue("phone")
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		// user, err := SignInUser(db, phone, hashedPassword)
		user, err := model.UserByPhoneOrEmail(db, phone, email)
		if err != nil || user == nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{StatusCode: 205, Message: "手机号或密码错误"})
			return
		}

		if user != nil {
			err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(password))
			if err != nil {
				util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{StatusCode: 205, Message: "手机号或密码错误"})
				return
			}

			// TODO: Touch login_date after successful login
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"uid":        user.ID,
				"last_login": time.Now(),
			})

			// Sign and get the complete encoded token as a string using the secret
			tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
			if err != nil {
				panic("jwt sign token failed")
			}

			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int         `json:"status_code"`
				Message    string      `json:"msg"`
				Data       *model.User `json:"data"`
				JwtToken   string      `json:"token"`
			}{
				StatusCode: 200,
				Message:    "查询成功",
				Data:       user,
				JwtToken:   tokenString,
			})
			return
		}
	})
}

// GetUserEndpoint 获取用户
func GetUserEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		user, err := model.UserByID(db, id)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int    `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: 400,
				Message:    err.Error(),
			})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int         `json:"status_code"`
			Message    string      `json:"msg"`
			User       *model.User `json:"user"`
		}{
			StatusCode: 200,
			Message:    "用户详情",
			User:       user,
		})
	})
}

type UserProfileResult struct {
	User         *model.User         `json:"user"`
	Chips        []model.Chip        `json:"chips"`
	BuyRequests  []model.BuyRequest  `json:"buy_requests"`
	HelpRequests []model.HelpRequest `json:"help_requests"`
}

// GetUserProfileEndpoint 用户资料页
// 用户资料页  查看其发布的芯片发布数、求购、求助
func GetUserProfileEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		user, err := model.UserByID(db, id)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int    `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: 400,
				Message:    err.Error(),
			})
			return
		}

		chips, err := model.ChipsByUserID(db, user.ID, 0, 10)
		buyRequests, err := model.BuyRequestsByUserID(db, user.ID, 0, 10)
		helpRequests, err := model.HelpRequestsByUserID(db, user.ID, 0, 10)

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int               `json:"status_code"`
			Message    string            `json:"msg"`
			Data       UserProfileResult `json:"data"`
		}{
			StatusCode: 200,
			Message:    "用户资料页",
			Data: UserProfileResult{
				User:         user,
				Chips:        chips,
				BuyRequests:  buyRequests,
				HelpRequests: helpRequests,
			},
		})
	})
}

// ApplySellerEndpoint 申请成为卖家
func ApplySellerEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CheckRequiredParameters(r, "real_name", "identity_card_num", "identity_card_front", "identity_card_back")
		err := r.ParseForm()
		PanicIfNotNil(err)

		userID := GetUIDFromContext(r)
		user, err := model.UserByID(db, userID)
		PanicIfNotNil(err)
		if user == nil {
			panic(&RecordNotFound{
				Error: errors.New("找不到用户"),
			})
		}

		if err = util.SchemaDecoder.Decode(user, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}

		// user.IsVerified = null.BoolFrom(true)
		user.Role = null.IntFrom(2)

		err = user.ApplySeller(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{StatusCode: 220, Message: err.Error()})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{StatusCode: 200, Message: "申请成功，请等待审核"})
	})
}

// ApplyExpertEndpoint 申请成为专家
func ApplyExpertEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CheckRequiredParameters(r, "real_name", "identity_card_num", "identity_card_front", "identity_card_back", "expertise", "resume", "from_code")

		err := r.ParseForm()
		PanicIfNotNil(err)

		userID := GetUIDFromContext(r)
		user, err := model.UserByID(db, userID)
		PanicIfNotNil(err)
		if user == nil {
			panic(&RecordNotFound{
				Error: errors.New("找不到用户"),
			})
		}

		if err = util.SchemaDecoder.Decode(user, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}

		// TODO
		// user.IsVerified = null.BoolFrom(true)
		// TODO
		user.Role = null.IntFrom(3)

		err = user.ApplyExpert(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{
				StatusCode: 220,
				Message:    err.Error(),
			})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{
			StatusCode: 200,
			Message:    "申请成功，请等待审核",
		})
		return
	})
}

// 用户列表
func ListUsersEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		qs := r.URL.Query()
		count, _ := strconv.Atoi(qs.Get("count"))
		start, _ := strconv.Atoi(qs.Get("start"))

		if count < 1 {
			count = 10
		}

		if start < 0 {
			start = 0
		}

		users, err := model.GetAllUsers(db, start, count)
		if err != nil {
			util.RespondWithJSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int          `json:"status_code"`
			Message    string       `json:"msg"`
			Data       []model.User `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       users,
		})
	})
}
