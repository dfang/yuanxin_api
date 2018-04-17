package model

import (
	"errors"
	"fmt"
)

// 注册
func (u *User) RegisterUser(db XODB) error {
	var err error

	u, err = UserByPhone(db, u.Phone.String)

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

// 登录
func SignInUser(db XODB, phone, pwd string) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, nickname, pwd, phone, email, avatar, gender, created_at, login_date ` +
		`FROM news.user ` +
		`WHERE phone = ? AND pwd = ?`

	// run query
	XOLog(sqlstr, phone, pwd)
	u := User{}

	err = db.QueryRow(sqlstr, phone, pwd).Scan(&u.ID, &u.Nickname, &u.Pwd, &u.Phone, &u.Email, &u.Avatar, &u.Gender, &u.CreatedAt, &u.LoginDate)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &u, nil
}

func UserByEmail(db XODB, email string) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, nickname, pwd, phone, email, avatar, gender, created_at, login_date ` +
		`FROM news.user ` +
		`WHERE email = ?`

	// run query
	XOLog(sqlstr, email)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, email).Scan(&u.ID, &u.Nickname, &u.Pwd, &u.Phone, &u.Email, &u.Avatar, &u.Gender, &u.CreatedAt, &u.LoginDate)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func UserByPhone(db XODB, phone string) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, nickname, pwd, phone, email, avatar, gender, created_at, login_date ` +
		`FROM news.user ` +
		`WHERE phone = ?`

	// run query
	XOLog(sqlstr, phone)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, phone).Scan(&u.ID, &u.Nickname, &u.Pwd, &u.Phone, &u.Email, &u.Avatar, &u.Gender, &u.CreatedAt, &u.LoginDate)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
