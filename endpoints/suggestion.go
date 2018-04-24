package endpoints

import (
	"database/sql"
	"net/http"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
	null "gopkg.in/guregu/null.v3"
)

func SuggestionEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer RecoverEndpoint(w)

		CheckRequiredParameters(r, "user_id", "content")

		user_id := ParseParameterToInt(r, "user_id")

		suggestion := model.Suggestion{
			UserID:  null.IntFrom(int64(user_id)),
			Content: null.StringFrom(r.PostFormValue("content")),
		}

		err := suggestion.Insert(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{220, "提交失败"})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "提交成功"})
	})
}
