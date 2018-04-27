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

// 收藏
func PubishFavoriteEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer RecoverEndpoint(w)

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

// 取消收藏
func DestroyFavoriteEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer RecoverEndpoint(w)

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
