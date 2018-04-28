package endpoints

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	. "github.com/dfang/yuanxin/model"
	. "github.com/dfang/yuanxin/util"
)

func SessionEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		CheckRequiredParameters(r, "phone", "password")

		phone := r.PostFormValue("phone")
		password := r.PostFormValue("password")

		user, err := SignInUser(db, phone, password)
		if err != nil || user == nil {
			RespondWithJSON(w, http.StatusOK, PayLoadFrom{StatusCode: 205, Message: "手机号码或密码错误"})
			return
		}

		if user != nil {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"user":      user.ID,
				"timestamp": time.Now(),
			})

			// Sign and get the complete encoded token as a string using the secret
			tokenString, err := token.SignedString([]byte("My Secret"))
			if err != nil {
				fmt.Println(err)
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
