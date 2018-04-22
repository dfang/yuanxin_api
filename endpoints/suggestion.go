package endpoints

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
	null "gopkg.in/guregu/null.v3"
)

func SuggestionEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.PostFormValue("user_id") == "" {
			str := fmt.Sprintf("参数%s缺失", "user_id")
			w.Write([]byte(str))
			return
		}

		if r.PostFormValue("content") == "" {
			str := fmt.Sprintf("参数%s缺失", "content")
			w.Write([]byte(str))
			return
		}

		user_id, err := strconv.Atoi(r.PostFormValue("user_id"))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		suggestion := model.Suggestion{
			UserID:  null.IntFrom(int64(user_id)),
			Content: null.StringFrom(r.PostFormValue("content")),
		}

		err = suggestion.Insert(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode string `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: "220",
				Message:    "提交失败",
			})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode string `json:"status_code"`
			Message    string `json:"msg"`
		}{
			StatusCode: "200",
			Message:    "提交成功",
		})
		return

	})
}
