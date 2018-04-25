// Package model contains the types for schema 'news'.
package model

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	null "gopkg.in/guregu/null.v3"
)

// NewsItem represents a row from 'news.news_items'.
type NewsItem struct {
	ID          int         `json:"id"`          // id
	Title       null.String `json:"title"`       // title
	Description null.String `json:"description"` // description
	Body        null.String `json:"body"`        // body
	Type        null.String `json:"type"`        // type
	Link        null.String `json:"link"`        // link
	Image       null.String `json:"image"`       // image
	Source      null.String `json:"source"`      // source
	UpdatedAt   null.Time   `json:"updated_at"`  // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the NewsItem exists in the database.
func (ni *NewsItem) Exists() bool {
	return ni._exists
}

// Deleted provides information if the NewsItem has been deleted from the database.
func (ni *NewsItem) Deleted() bool {
	return ni._deleted
}

// Insert inserts the NewsItem to the database.
func (ni *NewsItem) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ni._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO news.news_items (` +
		`title, description, body, type, link, image, source, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, ni.Title, ni.Description, ni.Body, ni.Type, ni.Link, ni.Image, ni.Source, ni.UpdatedAt)
	res, err := db.Exec(sqlstr, ni.Title, ni.Description, ni.Body, ni.Type, ni.Link, ni.Image, ni.Source, ni.UpdatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	ni.ID = int(id)
	ni._exists = true

	return nil
}

// Update updates the NewsItem in the database.
func (ni *NewsItem) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ni._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ni._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE news.news_items SET ` +
		`title = ?, description = ?, body = ?, type = ?, link = ?, image = ?, source = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, ni.Title, ni.Description, ni.Body, ni.Type, ni.Link, ni.Image, ni.Source, ni.UpdatedAt, ni.ID)
	_, err = db.Exec(sqlstr, ni.Title, ni.Description, ni.Body, ni.Type, ni.Link, ni.Image, ni.Source, ni.UpdatedAt, ni.ID)
	return err
}

// Save saves the NewsItem to the database.
func (ni *NewsItem) Save(db XODB) error {
	if ni.Exists() {
		return ni.Update(db)
	}

	return ni.Insert(db)
}

// Delete deletes the NewsItem from the database.
func (ni *NewsItem) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ni._exists {
		return nil
	}

	// if deleted, bail
	if ni._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM news.news_items WHERE id = ?`

	// run query
	XOLog(sqlstr, ni.ID)
	_, err = db.Exec(sqlstr, ni.ID)
	if err != nil {
		return err
	}

	// set deleted
	ni._deleted = true

	return nil
}

// NewsItemByID retrieves a row from 'news.news_items' as a NewsItem.
//
// Generated from index 'news_item_id_pkey'.
func NewsItemByID(db XODB, id int) (*NewsItem, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, title, description, body, type, link, image, source, updated_at ` +
		`FROM news.news_items ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	ni := NewsItem{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&ni.ID, &ni.Title, &ni.Description, &ni.Body, &ni.Type, &ni.Link, &ni.Image, &ni.Source, &ni.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &ni, nil
}
