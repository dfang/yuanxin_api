// Package model contains the types for schema 'news'.
package model

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	"gopkg.in/guregu/null.v3"
)

// User represents a row from 'news.user'.
type User struct {
	ID        int         `json:"id"`         // id
	Nickname  string      `json:"nickname"`   // nickname
	Pwd       string      `json:"pwd"`        // pwd
	Phone     string      `json:"phone"`      // phone
	Email     string      `json:"email"`      // email
	Avatar    null.String `json:"avatar"`     // avatar
	Gender    null.String `json:"gender"`     // gender
	CreatedAt null.String `json:"created_at"` // created_at
	LoginDate null.String `json:"login_date"` // login_date

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted provides information if the User has been deleted from the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO news.user (` +
		`nickname, pwd, phone, email, avatar, gender, created_at, login_date` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, u.Nickname, u.Pwd, u.Phone, u.Email, u.Avatar, u.Gender, u.CreatedAt, u.LoginDate)
	res, err := db.Exec(sqlstr, u.Nickname, u.Pwd, u.Phone, u.Email, u.Avatar, u.Gender, u.CreatedAt, u.LoginDate)
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

// Update updates the User in the database.
func (u *User) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if u._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE news.user SET ` +
		`nickname = ?, pwd = ?, phone = ?, email = ?, avatar = ?, gender = ?, created_at = ?, login_date = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, u.Nickname, u.Pwd, u.Phone, u.Email, u.Avatar, u.Gender, u.CreatedAt, u.LoginDate, u.ID)
	_, err = db.Exec(sqlstr, u.Nickname, u.Pwd, u.Phone, u.Email, u.Avatar, u.Gender, u.CreatedAt, u.LoginDate, u.ID)
	return err
}

// Save saves the User to the database.
func (u *User) Save(db XODB) error {
	if u.Exists() {
		return u.Update(db)
	}

	return u.Insert(db)
}

// Delete deletes the User from the database.
func (u *User) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return nil
	}

	// if deleted, bail
	if u._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM news.user WHERE id = ?`

	// run query
	XOLog(sqlstr, u.ID)
	_, err = db.Exec(sqlstr, u.ID)
	if err != nil {
		return err
	}

	// set deleted
	u._deleted = true

	return nil
}

// UserByID retrieves a row from 'news.user' as a User.
//
// Generated from index 'user_id_pkey'.
func UserByID(db XODB, id int) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, nickname, pwd, phone, email, avatar, gender, created_at, login_date ` +
		`FROM news.user ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&u.ID, &u.Nickname, &u.Pwd, &u.Phone, &u.Email, &u.Avatar, &u.Gender, &u.CreatedAt, &u.LoginDate)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
