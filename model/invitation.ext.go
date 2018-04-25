// Package model contains the types for schema 'news'.
package model

// Code generated by xo. DO NOT EDIT.

// InvitationByID retrieves a row from 'news.invitation' as a Invitation.
//
// Generated from index 'invitation_id_pkey'.
func InvitationByCode(db XODB, code string) (*Invitation, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, invitation_code, has_activated ` +
		`FROM news.invitations ` +
		`WHERE invitation_code = ?`

	// run query
	XOLog(sqlstr, code)
	i := Invitation{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, code).Scan(&i.ID, &i.InvitationCode, &i.HasActivated)
	if err != nil {
		return nil, err
	}

	return &i, nil
}
