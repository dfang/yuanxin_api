package endpoints

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	null "gopkg.in/guregu/null.v3"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
)

// 收藏和取消收藏
func FavorableEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("NOT IMPLEMENTED"))

		CheckRequiredParameters(r, "user_id", "favorable_type", "favorable_id")
		var item model.Favorite
		if err := util.SchemaDecoder.Decode(&item, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}

		favorite, err := model.GetFavoriteBy(db, item.FavorableType.String, item.FavorableID.Int64, item.UserID.Int64)
		if favorite == nil || err != nil {
			item.CreatedAt = null.TimeFrom(time.Now())
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
		} else {
			err = favorite.Delete(db)
			PanicIfNotNil(err)
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "取消收藏成功"})
			return
		}

	})
}

// 收藏, 未使用
func PubishFavoriteEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CheckRequiredParameters(r, "user_id", "favorable_type", "favorable_id")
		var item model.Favorite
		if err := util.SchemaDecoder.Decode(&item, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}
		item.CreatedAt = null.TimeFrom(time.Now())

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
