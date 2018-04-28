package endpoints

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/dfang/yuanxin/util"
	"github.com/gorilla/mux"
	"github.com/guregu/null"
)

type ChipDetailResult struct {
	ID              int         `json:"id"`               // id
	UserID          null.Int    `json:"user_id"`          // user_id
	SerialNumber    null.String `json:"serial_number"`    // serial_number
	Vendor          null.String `json:"vendor"`           // vendor
	Amount          null.Int    `json:"amount"`           // amount
	ManufactureDate null.Time   `json:"manufacture_date"` // manufacture_date
	UnitPrice       null.Float  `json:"unit_price"`       // unit_price
	IsVerified      null.Bool   `json:"is_verified"`      // is_verified
	NickName        null.String `json:"nickname"`
	Avatar          null.String `json:"avatar"`
	IsLiked         null.Bool   `json:"is_liked"`
}

type HelpRequestDetailResult struct {
	ID        int         `json:"id"`         // id
	UserID    null.Int    `json:"user_id"`    // user_id
	Title     null.String `json:"title"`      // title
	Content   null.String `json:"content"`    // content
	Amount    null.Int    `json:"amount"`     // amount
	CreatedAt null.Time   `json:"created_at"` // created_at
	NickName  null.String `json:"nickname"`
	Avatar    null.String `json:"avatar"`
	IsLiked   null.Bool   `json:"is_liked"`
}

type BuyRequestDetailResult struct {
	ID        int         `json:"id"`         // id
	UserID    null.Int    `json:"user_id" `   // user_id
	Title     null.String `json:"title"`      // title
	Content   null.String `json:"content"`    // content
	Amount    null.Int    `json:"amount"`     // amount
	CreatedAt null.Time   `json:"created_at"` // created_at
	NickName  null.String `json:"nickname"`
	Avatar    null.String `json:"avatar"`
	IsLiked   null.Bool   `json:"is_liked"`
}

func GetChipEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sqlstr := "SELECT chips.*, users.nickname, users.avatar FROM chips JOIN users on users.id = chips.user_id where chips.id = ?;"

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic("convertion error")
		}
		var result ChipDetailResult
		err = db.QueryRow(sqlstr, id).Scan(&result.ID, &result.UserID, &result.SerialNumber, &result.Vendor, &result.Amount, &result.ManufactureDate, &result.UnitPrice, &result.IsVerified, &result.NickName, &result.Avatar)
		PanicIfNotNil(err)

		result.IsLiked = null.BoolFrom(false)

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int              `json:"status_code"`
			Message    string           `json:"msg"`
			Data       ChipDetailResult `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       result,
		})

	})
}

func GetHelpRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sqlstr := "SELECT help_requests.*, users.nickname, users.avatar FROM help_requests JOIN users on users.id = help_requests.user_id where help_requests.id = ?;"

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic("convertion error")
		}
		var result HelpRequestDetailResult

		err = db.QueryRow(sqlstr, id).Scan(&result.ID, &result.UserID, &result.Title, &result.Content, &result.Amount, &result.CreatedAt, &result.NickName, &result.Avatar)
		PanicIfNotNil(err)

		result.IsLiked = null.BoolFrom(false)

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int                     `json:"status_code"`
			Message    string                  `json:"msg"`
			Data       HelpRequestDetailResult `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       result,
		})
	})
}

func GetBuyRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sqlstr := "SELECT buy_requests.*, users.nickname, users.avatar FROM buy_requests JOIN users on users.id = buy_requests.user_id where buy_requests.id = ?;"

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic("convertion error")
		}
		var result BuyRequestDetailResult

		err = db.QueryRow(sqlstr, id).Scan(&result.ID, &result.UserID, &result.Title, &result.Content, &result.Amount, &result.CreatedAt, &result.NickName, &result.Avatar)
		PanicIfNotNil(err)

		result.IsLiked = null.BoolFrom(false)

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int                    `json:"status_code"`
			Message    string                 `json:"msg"`
			Data       BuyRequestDetailResult `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       result,
		})
	})
}
