package endpoints

import (
	"database/sql"
	"net/http"
	"time"

	null "gopkg.in/guregu/null.v3"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
)

type LikeResult struct {
	ID         int       `json:"id"`          // id
	UserID     null.Int  `json:"user_id"`     // user_id
	CommentID  null.Int  `json:"comment_id"`  // comment_id
	CreatedAt  null.Time `json:"created_at"`  // created_at
	LikesCount int       `json:"likes_count"` // likes_count
}

// 赞和取消赞（目前只有评论能赞)
func LikableEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("NOT IMPLEMENTED"))

		// user := r.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)
		// fmt.Println(user.Valid())
		// fmt.Fprintf(w, "%v", user)
		CheckRequiredParameters(r, "user_id", "comment_id")
		var item model.Like
		if err := util.SchemaDecoder.Decode(&item, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}

		like, err := model.GetLikeBy(db, item.CommentID.Int64, item.UserID.Int64)
		if like == nil || err != nil {
			item.CreatedAt = null.TimeFrom(time.Now())
			err := item.Insert(db)
			PanicIfNotNil(err)
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int        `json:"status_code"`
				Message    string     `json:"msg"`
				Like       LikeResult `json:"like"`
			}{
				StatusCode: 200,
				Message:    "点赞成功",
				Like: LikeResult{
					CommentID:  item.CommentID,
					UserID:     item.UserID,
					CreatedAt:  item.CreatedAt,
					LikesCount: 0,
				},
			})
			return
		} else {
			err = like.Delete(db)
			PanicIfNotNil(err)
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int        `json:"status_code"`
				Message    string     `json:"msg"`
				Like       LikeResult `json:"like"`
			}{
				StatusCode: 200,
				Message:    "取消赞成功",
				Like: LikeResult{
					CommentID:  item.CommentID,
					UserID:     item.UserID,
					CreatedAt:  item.CreatedAt,
					LikesCount: 0,
				},
			})
			return
		}
	})
}
