package endpoints

import (
	"database/sql"
	"net/http"
	"time"

	null "gopkg.in/guregu/null.v3"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
)

// 发布求助
func PublishHelpRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		CheckRequiredParameters(r, "title", "content", "amount")

		var hr model.HelpRequest
		userID := GetUIDFromContext(r)

		if err := util.SchemaDecoder.Decode(&hr, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}

		hr.UserID = null.IntFrom(int64(userID))
		hr.CreatedAt = null.TimeFrom(utcTimeWithNanos())

		err := hr.Insert(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{220, err.Error()})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "发布成功"})
		return
	})
}

// 发布求购
func PublishBuyRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		CheckRequiredParameters(r, "title", "content", "amount")
		var br model.BuyRequest
		userID := GetUIDFromContext(r)

		if err := util.SchemaDecoder.Decode(&br, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}

		br.UserID = null.IntFrom(int64(userID))
		br.CreatedAt = null.TimeFrom(utcTimeWithNanos())
		err := br.Insert(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{220, err.Error()})
			return
		}
		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "发布成功"})
	})
}

// 发布芯片
func PublishChipEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		CheckRequiredParameters(r, "serial_number", "vendor", "amount", "manufacture_date", "unit_price")
		var chip model.Chip

		manufactureDate := r.PostForm.Get("manufacture_date")
		r.PostForm.Del("manufacture_date")
		if err := util.SchemaDecoder.Decode(&chip, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}

		// 临时hack
		t, err := time.Parse("2006-01", manufactureDate)
		if err != nil {
			panic("manufacture_date 不合法")
		}
		userID := GetUIDFromContext(r)
		chip.UserID = null.IntFrom(int64(userID))
		chip.ManufactureDate = null.TimeFrom(t)
		chip.IsVerified = null.BoolFrom(true)

		err = chip.Insert(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{220, err.Error()})
			return
		}
		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "发布成功"})
	})
}

// 发布评论
func PublishCommentEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		CheckRequiredParameters(r, "commentable_type", "commentable_id", "content")
		var comment model.Comment

		if err := util.SchemaDecoder.Decode(&comment, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}
		userID := GetUIDFromContext(r)

		comment.UserID = null.IntFrom(int64(userID))
		comment.CreatedAt = null.TimeFrom(utcTimeWithNanos())
		comment.Likes = null.IntFrom(0)
		comment.IsPicked = null.BoolFrom(false)

		err := comment.Insert(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{220, err.Error()})
			return
		}
		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int           `json:"status_code"`
			Message    string        `json:"msg"`
			Data       model.Comment `json:"data"`
		}{
			StatusCode: 200,
			Message:    "评论成功",
			Data:       comment,
		})
	})
}

func utcTimeWithNanos() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC)
}
