package endpoints

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/dfang/yuanxin_api/model"
	"github.com/dfang/yuanxin_api/util"
	null "gopkg.in/guregu/null.v3"
)

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
		user.IsVerified = null.BoolFrom(true)
		// TODO
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
