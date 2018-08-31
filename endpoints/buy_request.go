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

// 发布求购
func PublishBuyRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CheckRequiredParameters(r, "title", "content", "amount")

		err := r.ParseForm()
		PanicIfNotNil(err)

		var br model.BuyRequest
		userID := GetUIDFromContext(r)

		if err := util.SchemaDecoder.Decode(&br, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}

		br.UserID = null.IntFrom(int64(userID))
		br.CreatedAt = null.TimeFrom(utcTimeWithNanos())
		err = br.Insert(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{220, err.Error()})
			return
		}
		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "发布成功"})
	})
}

// 求购列表
func ListBuyRequestEndpoint(db *sql.DB) http.HandlerFunc {
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

		buyRequests, err := model.GetBuyRequests(db, start, count)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{http.StatusInternalServerError, err.Error()})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int                `json:"status_code"`
			Message    string             `json:"msg"`
			Data       []model.BuyRequest `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       buyRequests,
		})
	})
}

// 我的求购列表
func ListMyBuyRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := GetUIDFromContext(r)
		qs := r.URL.Query()
		count, _ := strconv.Atoi(qs.Get("count"))
		start, _ := strconv.Atoi(qs.Get("start"))

		if count < 1 {
			count = 10
		}

		if start < 0 {
			start = 0
		}

		buyRequests, err := model.BuyRequestsByUserID(db, userID, start, count)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{http.StatusInternalServerError, err.Error()})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int                `json:"status_code"`
			Message    string             `json:"msg"`
			Data       []model.BuyRequest `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       buyRequests,
		})
	})
}

// 删除求购
func DeleteBuyRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		userID := GetUIDFromContext(r)

		br, err := model.BuyRequestByID(db, id)
		if br == nil || err != nil {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int    `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: 400,
				Message:    "找不到",
			})
			return
		}

		if br.UserID.Int64 != int64(userID) {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int    `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: 401,
				Message:    "无操作权限",
			})
			return
		}

		err = br.Delete(db)
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

type buyRequestDetailResult struct {
	ID        int         `json:"id"`         // id
	UserID    null.Int    `json:"user_id" `   // user_id
	Title     null.String `json:"title"`      // title
	Content   null.String `json:"content"`    // content
	Amount    null.Int    `json:"amount"`     // amount
	CreatedAt null.Time   `json:"created_at"` // created_at
	NickName  null.String `json:"nickname"`
	Avatar    null.String `json:"avatar"`
	IsLiked   null.Bool   `json:"is_liked"`
}

// GetBuyRequestEndpoint Get buy_request detail
func GetBuyRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sqlstr := "SELECT buy_requests.*, users.nickname, users.avatar FROM buy_requests JOIN users on users.id = buy_requests.user_id where buy_requests.id = ?;"
		// userID := GetUIDFromContext(r)
		vars := mux.Vars(r)
		var userID int
		userID, _ = strconv.Atoi(r.Header.Get("X-UserID"))

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic("convertion error")
		}
		var result buyRequestDetailResult

		err = db.QueryRow(sqlstr, id).Scan(&result.ID, &result.UserID, &result.Title, &result.Content, &result.Amount, &result.CreatedAt, &result.NickName, &result.Avatar, &result.IsLiked)
		PanicIfNotNil(err)

		// result.IsLiked = null.BoolFrom(false)
		flag := false
		if userID != 0 {
			flag = model.IsLikedByUser(db, "buy_request", id, userID)
		}
		result.IsLiked = null.BoolFrom(flag)

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int                    `json:"status_code"`
			Message    string                 `json:"msg"`
			Data       buyRequestDetailResult `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       result,
		})
	})
}
