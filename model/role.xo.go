// Package model contains the types for schema 'news'.
package model

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Role represents a row from 'news.role'.
type Role struct {
	ID                int            `json:"id"`                  // id
	RealName          sql.NullString `json:"real_name"`           // real_name
	IdentityCardNum   sql.NullString `json:"identity_card_num"`   // identity_card_num
	IdentityCardFront sql.NullString `json:"identity_card_front"` // identity_card_front
	IdentityCardEnd   sql.NullString `json:"identity_card_end"`   // identity_card_end
	License           sql.NullString `json:"license"`             // license
	Expertise         sql.NullString `json:"expertise"`           // expertise

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Role exists in the database.
func (r *Role) Exists() bool {
	return r._exists
}

// Deleted provides information if the Role has been deleted from the database.
func (r *Role) Deleted() bool {
	return r._deleted
}

// Insert inserts the Role to the database.
func (r *Role) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if r._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO news.role (` +
		`real_name, identity_card_num, identity_card_front, identity_card_end, license, expertise` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, r.RealName, r.IdentityCardNum, r.IdentityCardFront, r.IdentityCardEnd, r.License, r.Expertise)
	res, err := db.Exec(sqlstr, r.RealName, r.IdentityCardNum, r.IdentityCardFront, r.IdentityCardEnd, r.License, r.Expertise)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	r.ID = int(id)
	r._exists = true

	return nil
}

// Update updates the Role in the database.
func (r *Role) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !r._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if r._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE news.role SET ` +
		`real_name = ?, identity_card_num = ?, identity_card_front = ?, identity_card_end = ?, license = ?, expertise = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, r.RealName, r.IdentityCardNum, r.IdentityCardFront, r.IdentityCardEnd, r.License, r.Expertise, r.ID)
	_, err = db.Exec(sqlstr, r.RealName, r.IdentityCardNum, r.IdentityCardFront, r.IdentityCardEnd, r.License, r.Expertise, r.ID)
	return err
}

// Save saves the Role to the database.
func (r *Role) Save(db XODB) error {
	if r.Exists() {
		return r.Update(db)
	}

	return r.Insert(db)
}

// Delete deletes the Role from the database.
func (r *Role) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !r._exists {
		return nil
	}

	// if deleted, bail
	if r._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM news.role WHERE id = ?`

	// run query
	XOLog(sqlstr, r.ID)
	_, err = db.Exec(sqlstr, r.ID)
	if err != nil {
		return err
	}

	// set deleted
	r._deleted = true

	return nil
}

// RoleByID retrieves a row from 'news.role' as a Role.
//
// Generated from index 'role_id_pkey'.
func RoleByID(db XODB, id int) (*Role, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, real_name, identity_card_num, identity_card_front, identity_card_end, license, expertise ` +
		`FROM news.role ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	r := Role{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&r.ID, &r.RealName, &r.IdentityCardNum, &r.IdentityCardFront, &r.IdentityCardEnd, &r.License, &r.Expertise)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
