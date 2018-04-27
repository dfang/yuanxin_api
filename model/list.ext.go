package model

import (
	"database/sql"
	"fmt"
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
