// Package model contains the types for schema 'news'.
package model

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Authentication represents a row from 'news.authentication'.
type Authentication struct {
	ID     int            `json:"id"`      // id
	UserID sql.NullInt64  `json:"user_id"` // user_id
	UUID   sql.NullString `json:"uuid"`    // uuid
	Token  sql.NullString `json:"token"`   // token

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Authentication exists in the database.
func (a *Authentication) Exists() bool {
	return a._exists
}

// Deleted provides information if the Authentication has been deleted from the database.
func (a *Authentication) Deleted() bool {
	return a._deleted
}

// Insert inserts the Authentication to the database.
func (a *Authentication) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO news.authentication (` +
		`user_id, uuid, token` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, a.UserID, a.UUID, a.Token)
	res, err := db.Exec(sqlstr, a.UserID, a.UUID, a.Token)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	a.ID = int(id)
	a._exists = true

	return nil
}

// Update updates the Authentication in the database.
func (a *Authentication) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if a._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE news.authentication SET ` +
		`user_id = ?, uuid = ?, token = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, a.UserID, a.UUID, a.Token, a.ID)
	_, err = db.Exec(sqlstr, a.UserID, a.UUID, a.Token, a.ID)
	return err
}

// Save saves the Authentication to the database.
func (a *Authentication) Save(db XODB) error {
	if a.Exists() {
		return a.Update(db)
	}

	return a.Insert(db)
}

// Delete deletes the Authentication from the database.
func (a *Authentication) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !a._exists {
		return nil
	}

	// if deleted, bail
	if a._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM news.authentication WHERE id = ?`

	// run query
	XOLog(sqlstr, a.ID)
	_, err = db.Exec(sqlstr, a.ID)
	if err != nil {
		return err
	}

	// set deleted
	a._deleted = true

	return nil
}

// AuthenticationByID retrieves a row from 'news.authentication' as a Authentication.
//
// Generated from index 'authentication_id_pkey'.
func AuthenticationByID(db XODB, id int) (*Authentication, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, user_id, uuid, token ` +
		`FROM news.authentication ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	a := Authentication{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&a.ID, &a.UserID, &a.UUID, &a.Token)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
