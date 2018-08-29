package endpoints

import (
	"database/sql"

	"net/http"
	"strconv"

	"github.com/dfang/yuanxin_api/model"
	"github.com/dfang/yuanxin_api/util"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	null "gopkg.in/guregu/null.v3"
)

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

// 发布芯片
func PublishChipEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CheckRequiredParameters(r, "serial_number", "version", "volume", "amount", "unit_price")

		err := r.ParseForm()
		PanicIfNotNil(err)

		var chip model.Chip

		// manufactureDate := r.PostForm.Get("manufacture_date")
		// r.PostForm.Del("manufacture_date")
		if err := util.SchemaDecoder.Decode(&chip, r.PostForm); err != nil {
			PanicIfNotNil(err)
		}

		// 临时hack
		// t, err := time.Parse("2006-01", manufactureDate)
		// if err != nil {
		// 	panic("manufacture_date 不合法")
		// }
		userID := GetUIDFromContext(r)
		chip.UserID = null.IntFrom(int64(userID))
		// chip.ManufactureDate = null.TimeFrom(t)
		chip.IsVerified = null.BoolFrom(true)

		err = chip.Insert(db)
		if err != nil {
			util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{220, err.Error()})
			return
		}
		util.RespondWithJSON(w, http.StatusOK, PayLoadFrom{200, "发布成功"})
	})
}

// 我的芯片列表
func ListMyChipsEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		qs := r.URL.Query()
		userID := GetUIDFromContext(r)
		count, _ := strconv.Atoi(qs.Get("count"))
		start, _ := strconv.Atoi(qs.Get("start"))

		if count < 1 {
			count = 10
		}

		if start < 0 {
			start = 0
		}

		chips, err := model.ChipsByUserID(db, userID, start, count)
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

// 删除芯片
func DeleteChipEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		userID := GetUIDFromContext(r)

		chip, err := model.ChipByID(db, id)
		if chip == nil || err != nil {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int    `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: 400,
				Message:    "找不到",
			})
			return
		}

		if chip.UserID.Int64 != int64(userID) {
			util.RespondWithJSON(w, http.StatusOK, struct {
				StatusCode int    `json:"status_code"`
				Message    string `json:"msg"`
			}{
				StatusCode: 401,
				Message:    "无操作权限",
			})
			return
		}

		err = chip.Delete(db)
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

type chipDetailResult struct {
	ID              int                `json:"id"`               // id
	UserID          null.Int           `json:"user_id"`          // user_id
	SerialNumber    null.String        `json:"serial_number"`    // serial_number
	Vendor          null.String        `json:"vendor"`           // vendor
	Amount          null.Int           `json:"amount"`           // amount
	ManufactureDate null.Time          `json:"manufacture_date"` // manufacture_date
	UnitPrice       null.Float         `json:"unit_price"`       // unit_price
	Specification   null.String        `json:"specification"`    // specification
	IsVerified      null.Bool          `json:"is_verified"`      // is_verified
	Vesion          null.String        `json:"version"`
	Volume          null.Int           `json:"volume"`
	NickName        null.String        `json:"nickname"`
	Avatar          null.String        `json:"avatar"`
	IsLiked         null.Bool          `json:"is_liked"`
	Chips           []model.Chip       `json:"chips"`
	BuyRequests     []model.BuyRequest `json:"buy_requests"`
}

// GetChipEndpoint Get chip detail
func GetChipEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sqlstr := "SELECT chips.*, users.nickname, users.avatar FROM chips JOIN users on users.id = chips.user_id where chips.id = ?;"
		vars := mux.Vars(r)

		var userID int
		currentUser := r.Context().Value("user")
		if currentUser != nil {
			claims := currentUser.(*jwt.Token).Claims.(jwt.MapClaims)
			userID = int(claims["uid"].(float64))
		}

		// userID := GetUIDFromContext(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic("convertion error")
		}
		var result chipDetailResult
		err = db.QueryRow(sqlstr, id).Scan(&result.ID, &result.UserID, &result.SerialNumber, &result.Vendor, &result.Amount, &result.ManufactureDate, &result.UnitPrice, &result.Specification, &result.IsVerified, &result.Vendor, &result.Volume, &result.IsLiked, &result.NickName, &result.Avatar)
		PanicIfNotNil(err)

		// TODO need to query db
		// result.IsLiked = null.BoolFrom(false)
		flag := false
		if userID != 0 {
			flag = model.IsLikedByUser(db, "chip", id, userID)
		}
		result.IsLiked = null.BoolFrom(flag)

		chips, err := model.SearchChips(db, result.SerialNumber.String, 0, 10)
		buyRequests, err := model.SearchChipsInBuyRequests(db, result.SerialNumber.String, 0, 10)

		result.Chips = chips
		result.BuyRequests = buyRequests

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int              `json:"status_code"`
			Message    string           `json:"msg"`
			Data       chipDetailResult `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       result,
		})

	})
}
