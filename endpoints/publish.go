package endpoints

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	null "gopkg.in/guregu/null.v3"

	"github.com/dfang/yuanxin/model"
	"github.com/dfang/yuanxin/util"
)

func ListHelpRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if vars["user_id"] == "" {
			str := fmt.Sprintf("参数%s缺失", "user_id")
			w.Write([]byte(str))
			return
		}

	})
}

// 发布求助
func PublishHelpRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer RecoverEndpoint(w)

		CheckRequiredParameters(r, "user_id", "title", "content", "amount")
		// user_id := ParseParameterToInt(r, "user_id")
		// amount := ParseParameterToInt(r, "amount")

		// user, err := model.UserByID(db, user_id)
		// if err != nil {
		// 	w.Write([]byte(err.Error()))
		// 	return
		// 	// panic(err)
		// }

		var hr model.HelpRequest
		if err := util.SchemaDecoder.Decode(&hr, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}

		err := hr.Insert(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{220, err.Error()})
			return
		}

		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "发布成功"})
		return
	})
}

// 发布求购
func PublishBuyRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer RecoverEndpoint(w)

		CheckRequiredParameters(r, "user_id", "title", "content", "amount")
		var br model.BuyRequest

		if err := util.SchemaDecoder.Decode(&br, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}
		br.CreatedAt = null.TimeFrom(time.Now())

		err := br.Insert(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{220, err.Error()})
			return
		}
		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "发布成功"})

	})
}

// 发布芯片
func PublishChipEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer RecoverEndpoint(w)

		CheckRequiredParameters(r, "serial_number", "vendor", "amount", "manufacture_date", "unit_price")
		var chip model.Chip

		manufacture_date := r.PostForm.Get("manufacture_date")
		r.PostForm.Del("manufacture_date")
		if err := util.SchemaDecoder.Decode(&chip, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}

		// 临时hack
		t, err := time.Parse("2006-01", manufacture_date)
		if err != nil {
			panic("manufacture_date 不合法")
		}
		chip.ManufactureDate = null.TimeFrom(t)

		err = chip.Insert(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{220, err.Error()})
			return
		}
		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "发布成功"})
	})
}
