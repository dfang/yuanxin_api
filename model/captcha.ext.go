// Package model contains the types for schema 'news'.
package model

// Code generated by xo. DO NOT EDIT.

func CaptchaBy(db XODB, phone, code string) (*Captcha, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, phone, code ` +
		`FROM news.captcha ` +
		`WHERE phone = ? AND code = ?`

	// run query
	XOLog(sqlstr, phone, code)
	c := Captcha{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, phone, code).Scan(&c.ID, &c.Phone, &c.Code)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
