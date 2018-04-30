package endpoints

import (
	"database/sql"
	"net/http"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
)

// CheckInvitationCodeEndpoint Check invitation code valid when apply expert
func CheckInvitationCodeEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CheckRequiredParameter(r, "invitation_code")
		invitationCode := r.PostFormValue("invitation_code")

		invitation, err := model.InvitationByCode(db, invitationCode)
		if err != nil {
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
		}

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
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode string `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: "200",
			Message:    "验证码有效",
		})
		return
	})
}
