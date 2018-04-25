// Package model contains the types for schema 'news'.
package model

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
)

// Comment represents a row from 'news.comments'.
type Comment struct {
	ID              int            `json:"id"`               // id
	CommentableType sql.NullString `json:"commentable_type"` // commentable_type
	CommentableID   sql.NullInt64  `json:"commentable_id"`   // commentable_id
	UserID          sql.NullInt64  `json:"user_id"`          // user_id
	CreatedAt       mysql.NullTime `json:"created_at"`       // created_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Comment exists in the database.
func (c *Comment) Exists() bool {
	return c._exists
}

// Deleted provides information if the Comment has been deleted from the database.
func (c *Comment) Deleted() bool {
	return c._deleted
}

// Insert inserts the Comment to the database.
func (c *Comment) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if c._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO news.comments (` +
		`commentable_type, commentable_id, user_id, created_at` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, c.CommentableType, c.CommentableID, c.UserID, c.CreatedAt)
	res, err := db.Exec(sqlstr, c.CommentableType, c.CommentableID, c.UserID, c.CreatedAt)
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

// Update updates the Comment in the database.
func (c *Comment) Update(db XODB) error {
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
	const sqlstr = `UPDATE news.comments SET ` +
		`commentable_type = ?, commentable_id = ?, user_id = ?, created_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, c.CommentableType, c.CommentableID, c.UserID, c.CreatedAt, c.ID)
	_, err = db.Exec(sqlstr, c.CommentableType, c.CommentableID, c.UserID, c.CreatedAt, c.ID)
	return err
}

// Save saves the Comment to the database.
func (c *Comment) Save(db XODB) error {
	if c.Exists() {
		return c.Update(db)
	}

	return c.Insert(db)
}

// Delete deletes the Comment from the database.
func (c *Comment) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM news.comments WHERE id = ?`

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

// CommentByID retrieves a row from 'news.comments' as a Comment.
//
// Generated from index 'comments_id_pkey'.
func CommentByID(db XODB, id int) (*Comment, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, commentable_type, commentable_id, user_id, created_at ` +
		`FROM news.comments ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	c := Comment{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&c.ID, &c.CommentableType, &c.CommentableID, &c.UserID, &c.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &c, nil
}