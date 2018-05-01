package model

import (
	"database/sql"
	"fmt"
	"log"

	null "gopkg.in/guregu/null.v3"
)

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
	statement := fmt.Sprintf("SELECT id, user_id, title, content, amount, created_at FROM news.help_requests ORDER BY created_at DESC LIMIT %d, %d", start, count)
	XOLog(statement)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	hrs := []HelpRequest{}

	for rows.Next() {
		var hr HelpRequest
		if err := rows.Scan(&hr.ID, &hr.UserID, &hr.Title, &hr.Content, &hr.Amount, &hr.CreatedAt); err != nil {
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

func GetFavorites(db *sql.DB, start, count int, favorableType string, favorableID int) ([]Favorite, error) {
	statement := fmt.Sprintf("SELECT * FROM news.favorites where favorable_type = '%s' AND favorable_id = '%d' LIMIT %d, %d", favorableType, favorableID, start, count)
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
func SearchChips(db *sql.DB, q string, start, count int) ([]Chip, error) {
	statement := fmt.Sprintf("select * from chips where serial_number like '%%%s%%' LIMIT %d, %d", q, start, count)
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
		if err := rows.Scan(&br.ID, &br.UserID, &br.Title, &br.Content, &br.Amount, &br.CreatedAt); err != nil {
			return nil, err
		}
		brs = append(brs, br)
	}
	return brs, nil
}

func ChipsByUserID(db *sql.DB, userID, start, count int) ([]Chip, error) {
	statement := fmt.Sprintf("SELECT id, user_id, serial_number, vendor, amount, manufacture_date, unit_price, specification, is_verified FROM chips WHERE user_id = %d ORDER BY manufacture_date DESC LIMIT %d, %d", userID, start, count)
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

func BuyRequestsByUserID(db *sql.DB, userID, start, count int) ([]BuyRequest, error) {
	statement := fmt.Sprintf("SELECT id, user_id, title, content, amount, created_at FROM news.buy_requests WHERE user_id = %d ORDER BY created_at DESC LIMIT %d, %d ", userID, start, count)
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

func HelpRequestsByUserID(db *sql.DB, userID, start, count int) ([]HelpRequest, error) {
	statement := fmt.Sprintf("SELECT id, user_id, title, content, amount, created_at FROM news.help_requests WHERE user_id = %d ORDER BY created_at DESC LIMIT %d, %d", userID, start, count)
	XOLog(statement)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	hrs := []HelpRequest{}

	for rows.Next() {
		var hr HelpRequest
		if err := rows.Scan(&hr.ID, &hr.UserID, &hr.Title, &hr.Content, &hr.Amount, &hr.CreatedAt); err != nil {
			return nil, err
		}
		hrs = append(hrs, hr)
	}

	return hrs, nil
}
