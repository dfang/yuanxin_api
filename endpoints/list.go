package endpoints

import (
	"database/sql"

	"net/http"
	"strconv"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
)

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

		help_requests, err := model.GetHelpRequests(db, start, count)
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
			Data:       help_requests,
		})
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

		buy_requests, err := model.GetBuyRequests(db, start, count)
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
			Data:       buy_requests,
		})
	})
}

// 芯片列表
func ListChipsEndpoint(db *sql.DB) http.HandlerFunc {
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

		chips, err := model.GetChips(db, start, count)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{http.StatusInternalServerError, err.Error()})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int          `json:"status_code"`
			Message    string       `json:"msg"`
			Data       []model.Chip `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       chips,
		})
	})
}

// 用户列表
func ListUsersEndpoint(db *sql.DB) http.HandlerFunc {
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

		users, err := model.GetAllUsers(db, start, count)
		if err != nil {
			util.RespondWithJSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int          `json:"status_code"`
			Message    string       `json:"msg"`
			Data       []model.User `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       users,
		})
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

// 新闻 求助 求购的评论
func ListCommentsEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer RecoverEndpoint(w)

		qs := r.URL.Query()
		CheckRequiredQueryStrings(r, "commentable_type", "commentable_id")

		count, _ := strconv.Atoi(qs.Get("count"))
		start, _ := strconv.Atoi(qs.Get("start"))
		commentable_id, _ := strconv.Atoi(qs.Get("commentable_id"))
		commentable_type := qs.Get("commentable_type")

		// CheckRequiredParameters(r, "commentable_type", "commentable_id")

		if count < 1 {
			count = 10
		}

		if start < 0 {
			start = 0
		}

		comments, err := model.GetComments(db, start, count, commentable_type, commentable_id)
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

// 收藏
func ListFavoritesEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer RecoverEndpoint(w)

		qs := r.URL.Query()
		CheckRequiredQueryStrings(r, "favorable_type", "favorable_id")

		count, _ := strconv.Atoi(qs.Get("count"))
		start, _ := strconv.Atoi(qs.Get("start"))
		favorable_id, _ := strconv.Atoi(qs.Get("favorable_id"))
		favorable_type := qs.Get("favorable_type")

		if count < 1 {
			count = 10
		}

		if start < 0 {
			start = 0
		}

		favorites, err := model.GetFavorites(db, start, count, favorable_type, favorable_id)
		if err != nil {
			util.RespondWithJSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int              `json:"status_code"`
			Message    string           `json:"msg"`
			Data       []model.Favorite `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       favorites,
		})
	})
}
