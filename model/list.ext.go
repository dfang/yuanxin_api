package model

import (
	"database/sql"
	"fmt"
	"log"

	null "gopkg.in/guregu/null.v3"
)

type ChipDetailResult struct {
	ID              int                `json:"id"`               // id
	UserID          null.Int           `json:"user_id"`          // user_id
	SerialNumber    null.String        `json:"serial_number"`    // serial_number
	Vendor          null.String        `json:"vendor"`           // vendor
	Amount          null.Int           `json:"amount"`           // amount
	ManufactureDate null.Time          `json:"manufacture_date"` // manufacture_date
	UnitPrice       null.Float         `json:"unit_price"`       // unit_price
	Specification   null.String        `json:"specification"`    // specification
	IsVerified      null.Bool          `json:"is_verified"`      // is_verified
	Version         null.String        `json:"version"`
	Volume          null.String        `json:"volume"`
	NickName        null.String        `json:"nickname"`
	Avatar          null.String        `json:"avatar"`
	IsLiked         null.Bool          `json:"is_liked"`
	Chips           []ChipDetailResult `json:"chips"`
	BuyRequests     []BuyRequest       `json:"buy_requests"`
}

func GetChips(db *sql.DB, start, count int) ([]Chip, error) {
	statement := fmt.Sprintf("SELECT id, user_id, serial_number, vendor, amount, manufacture_date, unit_price, specification, is_verified FROM chips ORDER BY manufacture_date DESC LIMIT %d, %d", start, count)
	XOLog(statement)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	chips := []Chip{}

	for rows.Next() {
		var chip Chip
		if err := rows.Scan(&chip.ID, &chip.UserID, &chip.SerialNumber, &chip.Vendor, &chip.Amount, &chip.ManufactureDate, &chip.UnitPrice, &chip.Specification, &chip.IsVerified); err != nil {
			return nil, err
		}
		chips = append(chips, chip)
	}

	return chips, nil
}

func GetHelpRequests(db *sql.DB, start, count int) ([]HelpRequest, error) {
	statement := fmt.Sprintf("SELECT id, user_id, title, content, amount, created_at, is_liked FROM news.help_requests ORDER BY created_at DESC LIMIT %d, %d", start, count)
	XOLog(statement)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	hrs := []HelpRequest{}

	for rows.Next() {
		var hr HelpRequest
		if err := rows.Scan(&hr.ID, &hr.UserID, &hr.Title, &hr.Content, &hr.Amount, &hr.CreatedAt, &hr.IsLiked); err != nil {
			return nil, err
		}
		hrs = append(hrs, hr)
	}

	return hrs, nil
}

func GetBuyRequests(db *sql.DB, start, count int) ([]BuyRequest, error) {
	statement := fmt.Sprintf("SELECT id, user_id, title, content, amount, created_at FROM news.buy_requests ORDER BY created_at DESC LIMIT %d, %d", start, count)
	XOLog(statement)

	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	brs := []BuyRequest{}

	for rows.Next() {
		var br BuyRequest
		if err := rows.Scan(&br.ID, &br.UserID, &br.Title, &br.Content, &br.Amount, &br.CreatedAt); err != nil {
			return nil, err
		}
		brs = append(brs, br)
	}

	return brs, nil
}

type CommentResult struct {
	ID              int         `json:"id"`                                         // id
	UserID          null.Int    `json:"user_id" schema:"user_id"`                   // user_id
	CommentableType null.String `json:"commentable_type" schema:"commentable_type"` // commentable_type
	CommentableID   null.Int    `json:"commentable_id" schema:"commentable_id"`     // commentable_id
	Title           null.String `json:"title"`                                      // 评论的对象的title
	Content         null.String `json:"content"`                                    // content
	IsPicked        null.Bool   `json:"is_picked" schema:"is_picked"`               // is_picked
	Likes           null.Int    `json:"likes"`                                      // likes
	CreatedAt       null.Time   `json:"created_at"`                                 // created_at
	IsLiked         null.Bool   `json:"is_liked" schema:"is_liked"`
	Nickname        null.String `json:"nickname"`
	Avatar          null.String `json:"avatar"`
}

func GetComments(db *sql.DB, start, count int, commentableType string, commentableID int) ([]CommentResult, error) {
	statement := fmt.Sprintf("SELECT comments.*, users.nickname, users.avatar FROM news.comments LEFT JOIN users ON users.id = comments.user_id where commentable_type = '%s' AND commentable_id = '%d' LIMIT %d, %d", commentableType, commentableID, start, count)
	log.Println(statement)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []CommentResult{}

	for rows.Next() {
		var item CommentResult
		if err := rows.Scan(&item.ID, &item.UserID, &item.CommentableType, &item.CommentableID, &item.Content, &item.IsPicked, &item.Likes, &item.CreatedAt, &item.Nickname, &item.Avatar); err != nil {
			return nil, err
		}

		// 此处应该查询，而不是默认返回false
		item.IsLiked = null.BoolFrom(false)

		items = append(items, item)
	}

	return items, nil
}

// 我的帮助(针对help_request的评论)
func GetMyHelpRequestComments(db *sql.DB, start, count int, userID int) ([]CommentResult, error) {
	statement := fmt.Sprintf("SELECT comments.*, help_requests.title, users.nickname, users.avatar FROM news.comments LEFT JOIN users ON users.id = comments.user_id LEFT JOIN help_requests on comments.commentable_id = help_requests.id where comments.user_id = %d AND comments.commentable_type = 'help_request' LIMIT %d, %d", userID, start, count)
	log.Println(statement)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []CommentResult{}

	for rows.Next() {
		var item CommentResult
		if err := rows.Scan(&item.ID, &item.UserID, &item.CommentableType, &item.CommentableID, &item.Content, &item.IsPicked, &item.Likes, &item.CreatedAt, &item.Title, &item.Nickname, &item.Avatar); err != nil {
			return nil, err
		}

		// 此处应该查询，而不是默认返回false
		item.IsLiked = null.BoolFrom(false)

		items = append(items, item)
	}

	return items, nil
}

// GetMyFavorites Get My Favorites
func GetMyFavorites(db *sql.DB, start, count int, userID int) ([]Favorite, error) {
	statement := fmt.Sprintf("SELECT * FROM news.favorites where user_id = '%d' LIMIT %d, %d", userID, start, count)
	log.Println(statement)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []Favorite{}

	for rows.Next() {
		var item Favorite
		if err := rows.Scan(&item.ID, &item.UserID, &item.FavorableType, &item.FavorableID, &item.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func GetNewsFavorites(db *sql.DB, start, count int, userID int) ([]NewsItem, error) {
	statement := fmt.Sprintf("SELECT * FROM news_items where id in ( SELECT favorable_id FROM news.favorites where user_id = '%d' and favorable_type = 'news_item') LIMIT %d, %d", userID, start, count)
	log.Println(statement)

	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}

	news := []NewsItem{}

	for rows.Next() {
		var item NewsItem
		if err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.Body, &item.Type, &item.Link, &item.Image, &item.Source, &item.UpdatedAt, &item.IsLiked); err != nil {
			return nil, err
		}
		news = append(news, item)
	}

	return news, nil
}

func GetChipsFavorites(db *sql.DB, start, count int, userID int) ([]Chip, error) {
	statement := fmt.Sprintf("SELECT * FROM chips where id in ( SELECT favorable_id FROM news.favorites where user_id = '%d' and favorable_type = 'chip') LIMIT %d, %d", userID, start, count)
	log.Println(statement)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []Chip{}

	for rows.Next() {
		var item Chip
		if err := rows.Scan(&item.ID, &item.UserID, &item.SerialNumber, &item.Vendor, &item.Amount, &item.ManufactureDate, &item.UnitPrice, &item.Specification, &item.IsVerified, &item.Version, &item.Volume, &item.IsLiked); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func GetBuyRequestFavorites(db *sql.DB, start, count int, userID int) ([]BuyRequest, error) {
	statement := fmt.Sprintf("SELECT * FROM buy_requests where id in ( SELECT favorable_id FROM news.favorites where user_id = '%d' and favorable_type = 'buy_request') LIMIT %d, %d", userID, start, count)
	log.Println(statement)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []BuyRequest{}

	for rows.Next() {
		var item BuyRequest
		if err := rows.Scan(&item.ID, &item.UserID, &item.Title, &item.Content, &item.Amount, &item.CreatedAt, &item.IsLiked); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func GetHelpRequestFavorites(db *sql.DB, start, count int, userID int) ([]HelpRequest, error) {
	statement := fmt.Sprintf("SELECT * FROM help_requests where id in ( SELECT favorable_id FROM news.favorites where user_id = '%d' and favorable_type = 'help_request') LIMIT %d, %d", userID, start, count)
	log.Println(statement)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []HelpRequest{}

	for rows.Next() {
		var item HelpRequest
		if err := rows.Scan(&item.ID, &item.UserID, &item.Title, &item.Content, &item.Amount, &item.CreatedAt, &item.IsLiked); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
func GetFavoriteBy(db *sql.DB, favorableType string, favorableID int64, userID int64) (*Favorite, error) {

	var err error

	// sql query
	// const sqlstr = `SELECT ` +
	// 	`id, user_id, favorable_type, favorable_id, created_at ` +
	// 	`FROM news.favorites ` +
	// 	`WHERE id = ?`
	sqlstr := fmt.Sprintf("SELECT * FROM news.favorites where favorable_type = '%s' AND favorable_id = '%d' AND user_id = '%d' ", favorableType, favorableID, userID)

	// run query
	XOLog(sqlstr, favorableType, favorableID, userID)
	item := Favorite{
		_exists: true,
	}

	err = db.QueryRow(sqlstr).Scan(&item.ID, &item.UserID, &item.FavorableType, &item.FavorableID, &item.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func GetLikeBy(db *sql.DB, commentID int64, userID int64) (*Like, error) {

	var err error

	// sql query
	sqlstr := fmt.Sprintf("SELECT * FROM news.likes where  comment_id = '%d' AND user_id = '%d' ", commentID, userID)

	// run query
	XOLog(sqlstr, commentID, userID)
	item := Like{
		_exists: true,
	}

	err = db.QueryRow(sqlstr).Scan(&item.ID, &item.UserID, &item.CommentID, &item.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// SearchChips  芯片详情页面
func SearchChips(db *sql.DB, q string, start, count int) ([]ChipDetailResult, error) {
	statement := fmt.Sprintf("select chips.*, users.nickname, users.avatar  from chips JOIN users on users.id = chips.user_id where serial_number like '%%%s%%' LIMIT %d, %d", q, start, count)
	XOLog(statement)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	chips := []ChipDetailResult{}

	for rows.Next() {
		var result ChipDetailResult
		if err := rows.Scan(&result.ID, &result.UserID, &result.SerialNumber, &result.Vendor, &result.Amount, &result.ManufactureDate, &result.UnitPrice, &result.Specification, &result.IsVerified, &result.Version, &result.Volume, &result.IsLiked, &result.NickName, &result.Avatar); err != nil {
			return nil, err
		}
		chips = append(chips, result)
	}

	return chips, nil
}

// SearchChipsInBuyRequests 芯片详情页面
func SearchChipsInBuyRequests(db *sql.DB, q string, start, count int) ([]BuyRequest, error) {
	statement := fmt.Sprintf("select * from news.buy_requests where title like '%%%s%%' OR content like '%%%s%%' LIMIT %d, %d", q, q, start, count)
	XOLog(statement)

	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	brs := []BuyRequest{}

	for rows.Next() {
		var br BuyRequest
		if err := rows.Scan(&br.ID, &br.UserID, &br.Title, &br.Content, &br.Amount, &br.CreatedAt, &br.IsLiked); err != nil {
			return nil, err
		}
		brs = append(brs, br)
	}
	return brs, nil
}

func ChipsByUserID(db *sql.DB, userID, start, count int) ([]Chip, error) {
	statement := fmt.Sprintf("SELECT id, user_id, serial_number, vendor, amount, manufacture_date, unit_price, specification, is_verified, version, volume, is_liked FROM chips WHERE user_id = %d ORDER BY manufacture_date DESC LIMIT %d, %d", userID, start, count)
	XOLog(statement)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	chips := []Chip{}

	for rows.Next() {
		var chip Chip
		if err := rows.Scan(&chip.ID, &chip.UserID, &chip.SerialNumber, &chip.Vendor, &chip.Amount, &chip.ManufactureDate, &chip.UnitPrice, &chip.Specification, &chip.IsVerified, &chip.Version, &chip.Volume, &chip.IsLiked); err != nil {
			return nil, err
		}
		chips = append(chips, chip)
	}

	return chips, nil
}

func BuyRequestsByUserID(db *sql.DB, userID, start, count int) ([]BuyRequest, error) {
	statement := fmt.Sprintf("SELECT id, user_id, title, content, amount, created_at, is_liked FROM news.buy_requests WHERE user_id = %d ORDER BY created_at DESC LIMIT %d, %d ", userID, start, count)
	XOLog(statement)

	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	brs := []BuyRequest{}

	for rows.Next() {
		var br BuyRequest
		if err := rows.Scan(&br.ID, &br.UserID, &br.Title, &br.Content, &br.Amount, &br.CreatedAt, &br.IsLiked); err != nil {
			return nil, err
		}
		brs = append(brs, br)
	}

	return brs, nil
}

func HelpRequestsByUserID(db *sql.DB, userID, start, count int) ([]HelpRequest, error) {
	statement := fmt.Sprintf("SELECT id, user_id, title, content, amount, created_at, is_liked FROM news.help_requests WHERE user_id = %d ORDER BY created_at DESC LIMIT %d, %d", userID, start, count)
	XOLog(statement)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	hrs := []HelpRequest{}

	for rows.Next() {
		var hr HelpRequest
		if err := rows.Scan(&hr.ID, &hr.UserID, &hr.Title, &hr.Content, &hr.Amount, &hr.CreatedAt, &hr.IsLiked); err != nil {
			return nil, err
		}
		hrs = append(hrs, hr)
	}

	return hrs, nil
}
