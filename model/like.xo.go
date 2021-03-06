// Package model contains the types for schema 'news'.
package model

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	null "gopkg.in/guregu/null.v3"
)

// Like represents a row from 'news.likes'.
type Like struct {
	ID        int       `json:"id"`                             // id
	UserID    null.Int  `json:"user_id" schema:"user_id"`       // user_id
	CommentID null.Int  `json:"comment_id" schema:"comment_id"` // comment_id
	CreatedAt null.Time `json:"created_at" schema:"created_at"` // created_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Like exists in the database.
func (l *Like) Exists() bool {
	return l._exists
}

// Deleted provides information if the Like has been deleted from the database.
func (l *Like) Deleted() bool {
	return l._deleted
}

// Insert inserts the Like to the database.
func (l *Like) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if l._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO news.likes (` +
		`user_id, comment_id` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, l.UserID, l.CommentID)
	res, err := db.Exec(sqlstr, l.UserID, l.CommentID)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	l.ID = int(id)
	l._exists = true

	return nil
}

// Update updates the Like in the database.
func (l *Like) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !l._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if l._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE news.likes SET ` +
		`user_id = ?, comment_id = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, l.UserID, l.CommentID, l.ID)
	_, err = db.Exec(sqlstr, l.UserID, l.CommentID, l.ID)
	return err
}

// Save saves the Like to the database.
func (l *Like) Save(db XODB) error {
	if l.Exists() {
		return l.Update(db)
	}

	return l.Insert(db)
}

// Delete deletes the Like from the database.
func (l *Like) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !l._exists {
		return nil
	}

	// if deleted, bail
	if l._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM news.likes WHERE id = ?`

	// run query
	XOLog(sqlstr, l.ID)
	_, err = db.Exec(sqlstr, l.ID)
	if err != nil {
		return err
	}

	// set deleted
	l._deleted = true

	return nil
}

// LikeByID retrieves a row from 'news.likes' as a Like.
//
// Generated from index 'likes_id_pkey'.
func LikeByID(db XODB, id int) (*Like, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, user_id, comment_id ` +
		`FROM news.likes ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	l := Like{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&l.ID, &l.UserID, &l.CommentID)
	if err != nil {
		return nil, err
	}

	return &l, nil
}
