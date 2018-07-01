package endpoints

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/dfang/yuanxin_api/model"
	"github.com/dfang/yuanxin_api/util"
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

// 邀请码列表
func ListInvitationsEndpoint(db *sql.DB) http.HandlerFunc {
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

		invitations, err := model.GetAllInvitations(db, start, count)
		if err != nil {
			util.RespondWithJSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int                `json:"status_code"`
			Message    string             `json:"msg"`
			Data       []model.Invitation `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       invitations,
		})

	})
}
