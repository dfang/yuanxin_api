package endpoints

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	. "github.com/dfang/yuanxin/model"
	. "github.com/dfang/yuanxin/util"
)

func SessionEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		CheckRequiredParameters(r, "phone", "password")
		// TODO:  Validate phone format

		phone := r.PostFormValue("phone")
		password := r.PostFormValue("password")

		user, err := SignInUser(db, phone, password)
		if err != nil || user == nil {
			RespondWithJSON(w, http.StatusOK, PayLoadFrom{StatusCode: 205, Message: "手机号码或密码错误"})
			return
		}

		if user != nil {
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
