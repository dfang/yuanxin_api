package endpoints

import (
	"database/sql"

	"net/http"
	"strconv"

	"github.com/dfang/yuanxin_api/model"
	"github.com/dfang/yuanxin_api/util"
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

// 收藏
func ListFavoritesEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		qs := r.URL.Query()
		CheckRequiredQueryStrings(r, "favorable_type", "favorable_id")

		count, _ := strconv.Atoi(qs.Get("count"))
		start, _ := strconv.Atoi(qs.Get("start"))
		favorableID, _ := strconv.Atoi(qs.Get("favorable_id"))
		favorableType := qs.Get("favorable_type")

		if count < 1 {
			count = 10
		}

		if start < 0 {
			start = 0
		}

		favorites, err := model.GetFavorites(db, start, count, favorableType, favorableID)
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
