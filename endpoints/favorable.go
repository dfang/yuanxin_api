package endpoints

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	null "gopkg.in/guregu/null.v3"

	"github.com/dfang/yuanxin_api/model"
	"github.com/dfang/yuanxin_api/util"
)

// FavorableEndpoint 收藏和取消收藏
func FavorableEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("NOT IMPLEMENTED"))

		CheckRequiredParameters(r, "favorable_type", "favorable_id")
		err := r.ParseForm()
		PanicIfNotNil(err)

		userID := GetUIDFromContext(r)

		var item model.Favorite
		if err := util.SchemaDecoder.Decode(&item, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}

		favorite, err := model.GetFavoriteBy(db, item.FavorableType.String, item.FavorableID.Int64, int64(userID))
		fmt.Println("lallallalal")
		fmt.Println(favorite)

		if favorite == nil || err != nil {
			userID := GetUIDFromContext(r)
			item.UserID = null.IntFrom(int64(userID))
			item.CreatedAt = null.TimeFrom(utcTimeWithNanos())
			err := item.Insert(db)
			PanicIfNotNil(err)
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int            `json:"status_code"`
				Message    string         `json:"msg"`
				Favorite   model.Favorite `json:"favorite"`
			}{
				StatusCode: 200,
				Message:    "收藏成功",
				Favorite:   item,
			})
			return
		}

		err = favorite.Delete(db)
		PanicIfNotNil(err)
		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "取消收藏成功"})
		return
	})
}

// 收藏, 未使用
func PubishFavoriteEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CheckRequiredParameters(r, "favorable_type", "favorable_id")

		err := r.ParseForm()
		PanicIfNotNil(err)

		var item model.Favorite
		if err := util.SchemaDecoder.Decode(&item, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}
		userID := GetUIDFromContext(r)
		item.UserID = null.IntFrom(int64(userID))
		item.CreatedAt = null.TimeFrom(utcTimeWithNanos())
		favorite, err := model.GetFavoriteBy(db, item.FavorableType.String, item.FavorableID.Int64, item.UserID.Int64)
		// PanicIfNotNil(err)
		if favorite == nil || err != nil {
			err := item.Insert(db)
			if err != nil {
				util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{220, err.Error()})
				return
			}
		}

		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "收藏成功"})
	})
}

// 取消收藏, 未使用
func DestroyFavoriteEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic("convertion error")
		}

		fmt.Println(id)

		favorite, err := model.FavoriteByID(db, id)
		PanicIfNotNil(err)

		if favorite != nil {
			err = favorite.Delete(db)
			PanicIfNotNil(err)

			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "操作成功"})
			return
		}

		// PanicIfNotNil(err)
		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "操作成功"})
	})
}

// 我的收藏 /my/favorites
func ListMyFavoritesEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		qs := r.URL.Query()
		// CheckRequiredQueryStrings(r, "favorable_type", "favorable_id")
		userID := GetUIDFromContext(r)
		count, _ := strconv.Atoi(qs.Get("count"))
		start, _ := strconv.Atoi(qs.Get("start"))
		// favorableID, _ := strconv.Atoi(qs.Get("favorable_id"))
		// favorableType := qs.Get("favorable_type")

		if count < 1 {
			count = 10
		}

		if start < 0 {
			start = 0
		}

		favorites, err := model.GetMyFavorites(db, start, count, userID)
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

// 收藏 /favorites?favorable_type=news
func ListMyFavoritesByTypeEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		qs := r.URL.Query()
		CheckRequiredQueryStrings(r, "favorable_type")
		userID := GetUIDFromContext(r)
		count, _ := strconv.Atoi(qs.Get("count"))
		start, _ := strconv.Atoi(qs.Get("start"))
		favorableType := qs.Get("favorable_type")

		if count < 1 {
			count = 10
		}

		if start < 0 {
			start = 0
		}

		switch favorableType {
		case "news_item":
			favorites, err := model.GetNewsFavorites(db, start, count, userID)
			if err != nil {
				util.RespondWithJSON(w, http.StatusInternalServerError, err.Error())
				return
			}

			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int              `json:"status_code"`
				Message    string           `json:"msg"`
				Data       []model.NewsItem `json:"data"`
			}{
				StatusCode: 200,
				Message:    "查询成功",
				Data:       favorites,
			})
		case "chip":
			favorites, err := model.GetChipsFavorites(db, start, count, userID)
			if err != nil {
				util.RespondWithJSON(w, http.StatusInternalServerError, err.Error())
				return
			}

			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int          `json:"status_code"`
				Message    string       `json:"msg"`
				Data       []model.Chip `json:"data"`
			}{
				StatusCode: 200,
				Message:    "查询成功",
				Data:       favorites,
			})
		case "buy_request":
			favorites, err := model.GetBuyRequestFavorites(db, start, count, userID)
			if err != nil {
				util.RespondWithJSON(w, http.StatusInternalServerError, err.Error())
				return
			}
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int                `json:"status_code"`
				Message    string             `json:"msg"`
				Data       []model.BuyRequest `json:"data"`
			}{
				StatusCode: 200,
				Message:    "查询成功",
				Data:       favorites,
			})
		case "help_request":
			favorites, err := model.GetHelpRequestFavorites(db, start, count, userID)
			if err != nil {
				util.RespondWithJSON(w, http.StatusInternalServerError, err.Error())
				return
			}
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int                 `json:"status_code"`
				Message    string              `json:"msg"`
				Data       []model.HelpRequest `json:"data"`
			}{
				StatusCode: 200,
				Message:    "查询成功",
				Data:       favorites,
			})
		}
	})
}
