package endpoints

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	null "gopkg.in/guregu/null.v3"

	"github.com/dfang/yuanxin_api/model"
	"github.com/dfang/yuanxin_api/util"
)

// 发布评论
func PublishCommentEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CheckRequiredParameters(r, "commentable_type", "commentable_id", "content")
		err := r.ParseForm()
		PanicIfNotNil(err)

		var comment model.Comment

		if err := util.SchemaDecoder.Decode(&comment, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}
		userID := GetUIDFromContext(r)

		comment.UserID = null.IntFrom(int64(userID))
		comment.CreatedAt = null.TimeFrom(utcTimeWithNanos())
		comment.Likes = null.IntFrom(0)
		comment.IsPicked = null.BoolFrom(false)

		err = comment.Insert(db)
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

// 新闻 求助 求购的评论
func ListCommentsEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		qs := r.URL.Query()
		CheckRequiredQueryStrings(r, "commentable_type", "commentable_id")

		count, _ := strconv.Atoi(qs.Get("count"))
		start, _ := strconv.Atoi(qs.Get("start"))
		commentableID, _ := strconv.Atoi(qs.Get("commentable_id"))
		commentableType := qs.Get("commentable_type")

		// CheckRequiredParameters(r, "commentable_type", "commentable_id")

		if count < 1 {
			count = 10
		}

		if start < 0 {
			start = 0
		}

		comments, err := model.GetComments(db, start, count, commentableType, commentableID)
		if err != nil {
			util.RespondWithJSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int                   `json:"status_code"`
			Message    string                `json:"msg"`
			Data       []model.CommentResult `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       comments,
		})
	})
}

func ListNewsCommentsEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func ListBuyRequestCommentEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func ListHelpRequestCommentEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}
