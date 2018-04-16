package model

import "errors"

// Insert inserts the User to the database.
func (u *User) RegisterUser(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO news.user (` +
		`nickname, pwd, phone, email` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, u.Nickname, u.Pwd, u.Phone, u.Email)
	res, err := db.Exec(sqlstr, u.Nickname, u.Pwd, u.Phone, u.Email)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	u.ID = int(id)
	u._exists = true

	return nil
}
