package model

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// 注册
func (u *User) RegisterUser(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO news.users (` +
		`nickname, pwd, phone, email, created_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, u.Nickname, u.Pwd, u.Phone, u.Email, time.Now())
	res, err := db.Exec(sqlstr, u.Nickname, u.Pwd, u.Phone, u.Email, time.Now())
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
		`FROM news.users ` +
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
		`id, nickname, pwd, phone, email, biography, avatar, gender, created_at, login_date ` +
		`FROM news.users ` +
		`WHERE email = ?`

	// run query
	XOLog(sqlstr, email)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, email).Scan(&u.ID, &u.Nickname, &u.Pwd, &u.Phone, &u.Email, &u.Biography, &u.Avatar, &u.Gender, &u.CreatedAt, &u.LoginDate)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func UserByPhone(db XODB, phone string) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, nickname, pwd, phone, email, biography, avatar, gender, created_at, login_date ` +
		`FROM news.users ` +
		`WHERE phone = ?`

	// run query
	XOLog(sqlstr, phone)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, phone).Scan(&u.ID, &u.Nickname, &u.Pwd, &u.Phone, &u.Email, &u.Biography, &u.Avatar, &u.Gender, &u.CreatedAt, &u.LoginDate)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// 更改个人信息
func (u *User) UpdateRegistrationInfo(db XODB) error {
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
	const sqlstr = `UPDATE news.users SET ` +
		`nickname = ?, phone = ?, avatar = ?, gender = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, u.Nickname, u.Phone, u.Avatar, u.Gender, u.ID)
	_, err = db.Exec(sqlstr, u.Nickname, u.Phone, u.Avatar, u.Gender, u.ID)
	return err
}

// 申请成为卖家
func (u *User) ApplySeller(db XODB) error {
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
	const sqlstr = `UPDATE news.users SET ` +
		`real_name = ?, identity_card_num = ?, identity_card_front = ?, identity_card_back = ?, license = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, u.RealName.String, u.IdentityCardNum.String, u.IdentityCardFront.String, u.IdentityCardBack.String, u.License.String, u.ID)
	_, err = db.Exec(sqlstr, u.RealName.String, u.IdentityCardNum.String, u.IdentityCardFront.String, u.IdentityCardBack.String, u.License.String, u.ID)
	return err
}

// 申请成为专家
func (u *User) ApplyExpert(db XODB) error {
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
	const sqlstr = `UPDATE news.users SET ` +
		`real_name = ?, identity_card_num = ?, identity_card_front = ?, identity_card_back = ?, expertise = ?, resume = ?, from_code = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, u.RealName.String, u.IdentityCardNum.String, u.IdentityCardFront.String, u.IdentityCardBack.String, u.Expertise.String, u.Resume.String, x.FromCode.String u.ID)
	_, err = db.Exec(sqlstr, u.RealName.String, u.IdentityCardNum.String, u.IdentityCardFront.String, u.IdentityCardBack.String, u.Expertise.String, u.Resume.String, x.FromCode.String, u.ID)
	return err
}

// 用户列表
func GetAllUsers(db *sql.DB, start, count int) ([]User, error) {
	sqlstr := fmt.Sprintf("SELECT id, nickname, pwd, phone, email, avatar, gender, biography, created_at, login_date, real_name, identity_card_num, identity_card_front, identity_card_back, from_code, license, expertise, resume, role, is_verified FROM news.users LIMIT %d, %d", start, count)

	rows, err := db.Query(sqlstr)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Nickname, &user.Pwd, &user.Phone, &user.Email, &user.Avatar, &user.Gender, &user.Biography, &user.CreatedAt, &user.LoginDate, &user.RealName, &user.IdentityCardNum, &user.IdentityCardFront, &user.IdentityCardBack, &user.FromCode, &user.License, &user.Expertise, &user.Resume, &user.Role, &user.IsVerified); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
