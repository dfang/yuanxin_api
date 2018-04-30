package endpoints

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	. "github.com/dfang/yuanxin/model"
	. "github.com/dfang/yuanxin/util"
)

func SessionEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		CheckRequiredParameters(r, "password")
		// TODO:  Validate phone format

		phone := r.PostFormValue("phone")
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")

		// user, err := SignInUser(db, phone, hashedPassword)
		user, err := UserByPhoneOrEmail(db, phone, email)
		if err != nil || user == nil {
			RespondWithJSON(w, http.StatusOK, PayLoadFrom{StatusCode: 205, Message: "手机号或密码错误"})
			return
		}

		if user != nil {
			err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(password))
			if err != nil {
				RespondWithJSON(w, http.StatusOK, PayLoadFrom{StatusCode: 205, Message: "手机号或密码错误"})
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

			RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int    `json:"status_code"`
				Message    string `json:"msg"`
				Data       *User  `json:"data"`
				JwtToken   string `json:"token"`
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
