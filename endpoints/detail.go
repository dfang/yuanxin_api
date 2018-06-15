package endpoints

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/dfang/yuanxin_api/model"
	"github.com/dfang/yuanxin_api/util"
	"github.com/gorilla/mux"
	"github.com/guregu/null"
)

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

type helpRequestDetailResult struct {
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

type buyRequestDetailResult struct {
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

// GetChipEndpoint Get chip detail
func GetChipEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sqlstr := "SELECT chips.*, users.nickname, users.avatar FROM chips JOIN users on users.id = chips.user_id where chips.id = ?;"
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic("convertion error")
		}
		var result chipDetailResult
		err = db.QueryRow(sqlstr, id).Scan(&result.ID, &result.UserID, &result.SerialNumber, &result.Vendor, &result.Amount, &result.ManufactureDate, &result.UnitPrice, &result.Specification, &result.IsVerified, &result.Vendor, &result.Volume, &result.NickName, &result.Avatar)
		PanicIfNotNil(err)

		// TODO need to query db
		result.IsLiked = null.BoolFrom(false)

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

// func GetChipRelatedEndpoint(db *sql.DB) http.HandlerFunc {
// 	// 芯片详情页面
// 	// select * from chips where serial_number like "%11%";
// 	// select * from buy_requests where title like "%11%" OR content like "%11%";
// }

// GetHelpRequestEndpoint Get help_request detail
func GetHelpRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sqlstr := "SELECT help_requests.*, users.nickname, users.avatar FROM help_requests JOIN users on users.id = help_requests.user_id where help_requests.id = ?;"

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic("convertion error")
		}
		var result helpRequestDetailResult

		err = db.QueryRow(sqlstr, id).Scan(&result.ID, &result.UserID, &result.Title, &result.Content, &result.Amount, &result.CreatedAt, &result.NickName, &result.Avatar)
		PanicIfNotNil(err)

		result.IsLiked = null.BoolFrom(false)

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int                     `json:"status_code"`
			Message    string                  `json:"msg"`
			Data       helpRequestDetailResult `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       result,
		})
	})
}

// GetBuyRequestEndpoint Get buy_request detail
func GetBuyRequestEndpoint(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sqlstr := "SELECT buy_requests.*, users.nickname, users.avatar FROM buy_requests JOIN users on users.id = buy_requests.user_id where buy_requests.id = ?;"

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic("convertion error")
		}
		var result buyRequestDetailResult

		err = db.QueryRow(sqlstr, id).Scan(&result.ID, &result.UserID, &result.Title, &result.Content, &result.Amount, &result.CreatedAt, &result.NickName, &result.Avatar)
		PanicIfNotNil(err)

		result.IsLiked = null.BoolFrom(false)

		util.RespondWithJSON(w, http.StatusOK, struct {
			StatusCode int                    `json:"status_code"`
			Message    string                 `json:"msg"`
			Data       buyRequestDetailResult `json:"data"`
		}{
			StatusCode: 200,
			Message:    "查询成功",
			Data:       result,
		})
	})
}
