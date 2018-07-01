package endpoints

import (
	"database/sql"

	"net/http"
	"strconv"

	"github.com/dfang/yuanxin_api/model"
	"github.com/dfang/yuanxin_api/util"
	"github.com/gorilla/mux"
	null "gopkg.in/guregu/null.v3"
)

// 发布求助
func PublishHelpRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CheckRequiredParameters(r, "title", "content", "amount")

		err := r.ParseForm()
		PanicIfNotNil(err)

		var hr model.HelpRequest
		userID := GetUIDFromContext(r)

		if err := util.SchemaDecoder.Decode(&hr, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}

		hr.UserID = null.IntFrom(int64(userID))
		hr.CreatedAt = null.TimeFrom(utcTimeWithNanos())

		err = hr.Insert(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{220, err.Error()})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "发布成功"})
		return
	})
}

// 求助列表
func ListHelpRequestEndpoint(db *sql.DB) http.HandlerFunc {
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

		helpRequests, err := model.GetHelpRequests(db, start, count)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{http.StatusInternalServerError, err.Error()})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int                 `json:"status_code"`
			Message    string              `json:"msg"`
			Data       []model.HelpRequest `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       helpRequests,
		})
	})
}

// 我的求助列表
func ListMyHelpRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		qs := r.URL.Query()
		userID := GetUIDFromContext(r)
		count, _ := strconv.Atoi(qs.Get("count"))
		start, _ := strconv.Atoi(qs.Get("start"))

		if count < 1 {
			count = 10
		}

		if start < 0 {
			start = 0
		}

		helpRequests, err := model.HelpRequestsByUserID(db, userID, start, count)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{http.StatusInternalServerError, err.Error()})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int                 `json:"status_code"`
			Message    string              `json:"msg"`
			Data       []model.HelpRequest `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       helpRequests,
		})
	})
}

// 删除求助
func DeleteHelpRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		userID := GetUIDFromContext(r)

		hr, err := model.HelpRequestByID(db, id)
		if hr == nil || err != nil {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int    `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: 400,
				Message:    "找不到",
			})
			return
		}

		if hr.UserID.Int64 != int64(userID) {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int    `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: 401,
				Message:    "无操作权限",
			})
			return
		}

		err = hr.Delete(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{http.StatusInternalServerError, err.Error()})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int    `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: 200,
			Message:    "操作成功",
		})

	})
}

type helpRequestDetailResult struct {
	ID        int         `json:"id"`         // id
	UserID    null.Int    `json:"user_id"`    // user_id
	Title     null.String `json:"title"`      // title
	Content   null.String `json:"content"`    // content
	Amount    null.Int    `json:"amount"`     // amount
	CreatedAt null.Time   `json:"created_at"` // created_at
	NickName  null.String `json:"nickname"`
	Avatar    null.String `json:"avatar"`
	IsLiked   null.Bool   `json:"is_liked"`
}

// GetHelpRequestEndpoint Get help_request detail
func GetHelpRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sqlstr := "SELECT help_requests.*, users.nickname, users.avatar FROM help_requests JOIN users on users.id = help_requests.user_id where help_requests.id = ?;"

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic("convertion error")
		}
		var result helpRequestDetailResult

		err = db.QueryRow(sqlstr, id).Scan(&result.ID, &result.UserID, &result.Title, &result.Content, &result.Amount, &result.CreatedAt, &result.NickName, &result.Avatar)
		PanicIfNotNil(err)

		result.IsLiked = null.BoolFrom(false)

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int                     `json:"status_code"`
			Message    string                  `json:"msg"`
			Data       helpRequestDetailResult `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       result,
		})
	})
}
