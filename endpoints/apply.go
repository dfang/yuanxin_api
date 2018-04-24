package endpoints

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
	null "gopkg.in/guregu/null.v3"
)

// 申请成为卖家
func ApplySellerEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.PostFormValue("user_id") == "" {
			str := fmt.Sprintf("参数%s缺失", "user_id")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("real_name") == "" {
			str := fmt.Sprintf("参数%s缺失", "real_name")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("identity_card_num") == "" {
			str := fmt.Sprintf("参数%s缺失", "identity_card_num")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("identity_card_front") == "" {
			str := fmt.Sprintf("参数%s缺失", "identity_card_front")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("identity_card_back") == "" {
			str := fmt.Sprintf("参数%s缺失", "identity_card_back")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("license") == "" {
			str := fmt.Sprintf("参数%s缺失", "license")
			w.Write([]byte(str))
			return
		}

		// 检查必须的参数
		// Find

		user_id, err := strconv.Atoi(r.PostFormValue("user_id"))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		user, err := model.UserByID(db, user_id)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
			// panic(err)
		}

		if user == nil {
			w.Write([]byte("找不到此用户"))
			return
		}

		user.RealName = null.StringFrom(r.PostFormValue("real_name"))
		user.License = null.StringFrom(r.PostFormValue("license"))
		user.IdentityCardNum = null.StringFrom(r.PostFormValue("identity_card_num"))
		user.IdentityCardFront = null.StringFrom(r.PostFormValue("identity_card_front"))
		user.IdentityCardBack = null.StringFrom(r.PostFormValue("identity_card_back"))

		user.IsVerified = null.BoolFrom(true)
		user.Role = null.IntFrom(2)

		err = user.ApplySeller(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode string `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: "220",
				Message:    "申请失败",
			})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode string `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: "200",
			Message:    "申请成功，请等待审核",
		})
		return

	})
}

// 申请成为专家
func ApplyExpertEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("not implemented"))

		if r.PostFormValue("user_id") == "" {
			str := fmt.Sprintf("参数%s缺失", "user_id")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("real_name") == "" {
			str := fmt.Sprintf("参数%s缺失", "real_name")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("identity_card_num") == "" {
			str := fmt.Sprintf("参数%s缺失", "identity_card_num")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("identity_card_front") == "" {
			str := fmt.Sprintf("参数%s缺失", "identity_card_front")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("identity_card_back") == "" {
			str := fmt.Sprintf("参数%s缺失", "identity_card_back")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("expertise") == "" {
			str := fmt.Sprintf("参数%s缺失", "expertise")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("resume") == "" {
			str := fmt.Sprintf("参数%s缺失", "resume")
			w.Write([]byte(str))
			return
		}

		user_id, err := strconv.Atoi(r.PostFormValue("user_id"))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		user, err := model.UserByID(db, user_id)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
			// panic(err)
		}

		if user == nil {
			w.Write([]byte("找不到此用户"))
			return
		}

		user.RealName = null.StringFrom(r.PostFormValue("real_name"))
		user.IdentityCardNum = null.StringFrom(r.PostFormValue("identity_card_num"))
		user.IdentityCardFront = null.StringFrom(r.PostFormValue("identity_card_front"))
		user.IdentityCardBack = null.StringFrom(r.PostFormValue("identity_card_back"))
		user.Expertise = null.StringFrom(r.PostFormValue("expertise"))
		user.Resume = null.StringFrom(r.PostFormValue("resume"))

		user.IsVerified = null.BoolFrom(true)
		user.Role = null.IntFrom(1)

		err = user.ApplyExpert(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode string `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: "220",
				Message:    "申请失败",
			})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode string `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: "200",
			Message:    "申请成功，请等待审核",
		})
		return
	})
}
