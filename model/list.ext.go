package model

import (
	"database/sql"
	"fmt"
	"log"
)

func GetChips(db *sql.DB, start, count int) ([]Chip, error) {
	statement := fmt.Sprintf("SELECT id, user_id, serial_number, vendor, amount, manufacture_date, unit_price, is_verified FROM chips ORDER BY manufacture_date DESC LIMIT %d, %d", start, count)
	XOLog(statement)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	chips := []Chip{}

	for rows.Next() {
		var chip Chip
		if err := rows.Scan(&chip.ID, &chip.UserID, &chip.SerialNumber, &chip.Vendor, &chip.Amount, &chip.ManufactureDate, &chip.UnitPrice, &chip.IsVerified); err != nil {
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

func GetComments(db *sql.DB, start, count int, commentable_type string, commentable_id int) ([]Comment, error) {
	statement := fmt.Sprintf("SELECT * FROM news.comments where commentable_type = '%s' AND commentable_id = '%d' LIMIT %d, %d", commentable_type, commentable_id, start, count)
	log.Println(statement)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []Comment{}

	for rows.Next() {
		var item Comment
		if err := rows.Scan(&item.ID, &item.UserID, &item.CommentableType, &item.CommentableID, &item.Content, &item.IsPicked, &item.Likes, &item.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func GetFavorites(db *sql.DB, start, count int, favorable_type string, favorable_id int) ([]Favorite, error) {
	statement := fmt.Sprintf("SELECT * FROM news.favorites where favorable_type = '%s' AND favorable_id = '%d' LIMIT %d, %d", favorable_type, favorable_id, start, count)
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

func GetFavoriteBy(db *sql.DB, favorable_type string, favorable_id int64, user_id int64) (*Favorite, error) {

	var err error

	// sql query
	// const sqlstr = `SELECT ` +
	// 	`id, user_id, favorable_type, favorable_id, created_at ` +
	// 	`FROM news.favorites ` +
	// 	`WHERE id = ?`
	sqlstr := fmt.Sprintf("SELECT * FROM news.favorites where favorable_type = '%s' AND favorable_id = '%d' AND user_id = '%d' ", favorable_type, favorable_id, user_id)

	// run query
	XOLog(sqlstr, favorable_type, favorable_id, user_id)
	item := Favorite{
		_exists: true,
	}

	err = db.QueryRow(sqlstr).Scan(&item.ID, &item.UserID, &item.FavorableType, &item.FavorableID, &item.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
