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
