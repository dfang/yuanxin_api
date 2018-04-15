// Package model contains the types for schema 'news'.
package model

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// UserRole represents a row from 'news.user_role'.
type UserRole struct {
	ID     int           `json:"id"`      // id
	RoleID sql.NullInt64 `json:"role_id"` // role_id
	UserID sql.NullInt64 `json:"user_id"` // user_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the UserRole exists in the database.
func (ur *UserRole) Exists() bool {
	return ur._exists
}

// Deleted provides information if the UserRole has been deleted from the database.
func (ur *UserRole) Deleted() bool {
	return ur._deleted
}

// Insert inserts the UserRole to the database.
func (ur *UserRole) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ur._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO news.user_role (` +
		`role_id, user_id` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, ur.RoleID, ur.UserID)
	res, err := db.Exec(sqlstr, ur.RoleID, ur.UserID)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	ur.ID = int(id)
	ur._exists = true

	return nil
}

// Update updates the UserRole in the database.
func (ur *UserRole) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ur._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ur._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE news.user_role SET ` +
		`role_id = ?, user_id = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, ur.RoleID, ur.UserID, ur.ID)
	_, err = db.Exec(sqlstr, ur.RoleID, ur.UserID, ur.ID)
	return err
}

// Save saves the UserRole to the database.
func (ur *UserRole) Save(db XODB) error {
	if ur.Exists() {
		return ur.Update(db)
	}

	return ur.Insert(db)
}

// Delete deletes the UserRole from the database.
func (ur *UserRole) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ur._exists {
		return nil
	}

	// if deleted, bail
	if ur._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM news.user_role WHERE id = ?`

	// run query
	XOLog(sqlstr, ur.ID)
	_, err = db.Exec(sqlstr, ur.ID)
	if err != nil {
		return err
	}

	// set deleted
	ur._deleted = true

	return nil
}

// UserRoleByID retrieves a row from 'news.user_role' as a UserRole.
//
// Generated from index 'user_role_id_pkey'.
func UserRoleByID(db XODB, id int) (*UserRole, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, role_id, user_id ` +
		`FROM news.user_role ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	ur := UserRole{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&ur.ID, &ur.RoleID, &ur.UserID)
	if err != nil {
		return nil, err
	}

	return &ur, nil
}
