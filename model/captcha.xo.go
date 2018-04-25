// Package model contains the types for schema 'news'.
package model

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"log"

	null "gopkg.in/guregu/null.v3"
)

// Captcha represents a row from 'news.captchas'.
type Captcha struct {
	ID    int         `json:"id"`    // id
	Phone null.String `json:"phone"` // phone
	Code  null.String `json:"code"`  // code

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Captcha exists in the database.
func (c *Captcha) Exists() bool {
	return c._exists
}

// Deleted provides information if the Captcha has been deleted from the database.
func (c *Captcha) Deleted() bool {
	return c._deleted
}

// Insert inserts the Captcha to the database.
func (c *Captcha) Insert(db XODB) error {
	var err error

	log.Println(c)

	// c, err = CaptchaByPhone(db, c.Phone.String)

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO news.captchas (` +
		`phone, code` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, c.Phone, c.Code)
	res, err := db.Exec(sqlstr, c.Phone, c.Code)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	c.ID = int(id)
	c._exists = true

	return nil
}

// Update updates the Captcha in the database.
func (c *Captcha) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if c._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE news.captchas SET ` +
		`phone = ?, code = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, c.Phone, c.Code, c.ID)
	_, err = db.Exec(sqlstr, c.Phone, c.Code, c.ID)
	return err
}

// Save saves the Captcha to the database.
func (c *Captcha) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Delete deletes the Captcha from the database.
func (c *Captcha) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !c._exists {
		return nil
	}

	// if deleted, bail
	if c._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM news.captchas WHERE id = ?`

	// run query
	XOLog(sqlstr, c.ID)
	_, err = db.Exec(sqlstr, c.ID)
	if err != nil {
		return err
	}

	// set deleted
	c._deleted = true

	return nil
}

// CaptchaByID retrieves a row from 'news.captchas' as a Captcha.
//
// Generated from index 'captcha_id_pkey'.
func CaptchaByID(db XODB, id int) (*Captcha, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, phone, code ` +
		`FROM news.captchas ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	c := Captcha{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&c.ID, &c.Phone, &c.Code)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
