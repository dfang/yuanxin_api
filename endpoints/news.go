package endpoints

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dfang/yuanxin_api/model"
	"github.com/dfang/yuanxin_api/util"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gorilla/mux"
	null "gopkg.in/guregu/null.v3"
)

func ListNewsItemEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		qs := r.URL.Query()
		count, _ := strconv.Atoi(qs.Get("count"))
		start, _ := strconv.Atoi(qs.Get("start"))
		t, _ := strconv.Atoi(qs.Get("type"))

		if count < 1 {
			count = 10
		}

		if start < 0 {
			start = 0
		}

		news, err := model.GetNews(db, start, count, model.NewsItemType(t))
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
			Data:       news,
		})
	})
}

func GetNewsItemEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var userID int
		currentUser := r.Context().Value("user")
		if currentUser != nil {
			claims := currentUser.(*jwt.Token).Claims.(jwt.MapClaims)
			userID = int(claims["uid"].(float64))
		}

		fmt.Println(userID)
		// fmt.Println(currentUser.(*jwt.Token).Claims)
		// fmt.Println(claims["uid"])
		// if currentUser != nil {
		// 	currentUser.(*jwt.Token).Claims["uid"]
		// }

		// currentUser := context.Get(r, "user")
		// fmt.Println("kkkkkkkkkkkkkkkk")
		// fmt.Println(currentUser)

		// userID := GetUIDFromContext(r)
		// fmt.Println(userID)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			util.RespondWithJSON(w, http.StatusBadRequest, "Invalid ID")
			return
		}

		// item := model.NewsItem{ID: id}
		// if err := item.GetNewsItem(db); err != nil {
		if ni, err := model.NewsItemByID(db, id); err != nil {
			log.Println(ni)
			switch err {
			case sql.ErrNoRows:
				// util.RespondWithJSON(w, http.StatusNotFound, "NewsItem not found")
				util.RespondWithJSON(w, http.StatusOK, struct {
					StatusCode int         `json:"status_code"`
					Message    string      `json:"msg"`
					Data       interface{} `json:"data"`
				}{
					StatusCode: 200,
					Message:    "查询成功",
					Data:       struct{}{},
				})
			default:
				util.RespondWithJSON(w, http.StatusInternalServerError, err.Error())
			}
			return
		} else {
			flag := false
			if userID != 0 {
				flag = model.IsLikedByUser(db, "news_item", id, userID)
			}
			ni.IsLiked = null.BoolFrom(flag)
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int             `json:"status_code"`
				Message    string          `json:"msg"`
				Data       *model.NewsItem `json:"data"`
			}{
				StatusCode: 200,
				Message:    "查询成功",
				Data:       ni,
			})
		}

	})
}
