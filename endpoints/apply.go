package endpoints

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
	null "gopkg.in/guregu/null.v3"
)

// 申请成为卖家
func ApplySellerEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CheckRequiredParameters(r, "user_id", "real_name", "identity_card_num", "identity_card_front", "identity_card_back", "from_code")
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

		user.IsVerified = null.BoolFrom(true)
		user.Role = null.IntFrom(2)

		err = user.ApplySeller(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{StatusCode: 220, Message: err.Error()})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{StatusCode: 200, Message: "申请成功，请等待审核"})
	})
}

// 申请成为专家
func ApplyExpertEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		CheckRequiredParameters(r, "user_id", "real_name", "identity_card_num", "identity_card_front", "identity_card_back", "expertise", "resume")

		user_id := ParseParameterToInt(r, "user_id")

		user, err := model.UserByID(db, user_id)
		PanicIfNotNil(err)
		if user == nil {
			panic(&RecordNotFound{
				Error: errors.New("找不到用户"),
			})
		}

		// user.RealName = null.StringFrom(r.PostFormValue("real_name"))
		// user.IdentityCardNum = null.StringFrom(r.PostFormValue("identity_card_num"))
		// user.IdentityCardFront = null.StringFrom(r.PostFormValue("identity_card_front"))
		// user.IdentityCardBack = null.StringFrom(r.PostFormValue("identity_card_back"))
		// user.Expertise = null.StringFrom(r.PostFormValue("expertise"))
		// user.Resume = null.StringFrom(r.PostFormValue("resume"))
		r.PostForm.Del("user_id")
		if err = util.SchemaDecoder.Decode(user, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}

		user.IsVerified = null.BoolFrom(true)
		user.Role = null.IntFrom(1)

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
