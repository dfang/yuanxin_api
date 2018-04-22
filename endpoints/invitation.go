package endpoints

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
)

func CheckInvitationCodeEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.PostFormValue("invitation_code") == "" {
			str := fmt.Sprintf("参数%s缺失", "invitation_code")
			w.Write([]byte(str))
			return
		}

		invitation_code := r.PostFormValue("invitation_code")

		invitation, err := model.InvitationByCode(db, invitation_code)
		if err != nil {
			log.Println(err)
			// w.Write([]byte("无效的邀请码"))
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode string `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: "221",
				Message:    "无效的邀请码",
			})
			return
		}

		if invitation == nil {
			// w.Write([]byte("无效的邀请码"))
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode string `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: "221",
				Message:    "无效的邀请码",
			})
			return
		} else {
			if invitation.HasActivated.Bool {
				// w.Write([]byte("邀请码已使用"))
				util.RespondWithJSON(w, http.StatusOK, struct {
					StatusCode string `json:"status_code"`
					Message    string `json:"msg"`
				}{
					StatusCode: "222",
					Message:    "邀请码已使用",
				})
				return
			} else {
				util.RespondWithJSON(w, http.StatusOK, struct {
					StatusCode string `json:"status_code"`
					Message    string `json:"msg"`
				}{
					StatusCode: "200",
					Message:    "验证码有效",
				})
				return
			}
		}
	})
}
