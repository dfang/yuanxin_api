package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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

func (u *User) UpdateAccInfo(db XODB) error {
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
		`acc_id = ?, acc_token = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, u.AccID, u.AccToken, u.ID)
	_, err = db.Exec(sqlstr, u.AccID, u.AccToken, u.ID)
	return err
}

func UserByEmail(db XODB, email string) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, nickname, pwd, phone, email, avatar, gender, biography, created_at, login_date, real_name, identity_card_num, identity_card_front, identity_card_back, from_code, license, expertise, resume, role, is_verified ` +
		`FROM news.users ` +
		`WHERE email = ?`

	// run query
	XOLog(sqlstr, email)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, email).Scan(&u.ID, &u.Nickname, &u.Pwd, &u.Phone, &u.Email, &u.Avatar, &u.Gender, &u.Biography, &u.CreatedAt, &u.LoginDate, &u.RealName, &u.IdentityCardNum, &u.IdentityCardFront, &u.IdentityCardBack, &u.FromCode, &u.License, &u.Expertise, &u.Resume, &u.Role, &u.IsVerified)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func UserByPhone(db XODB, phone string) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, nickname, pwd, phone, email, avatar, gender, biography, created_at, login_date, real_name, identity_card_num, identity_card_front, identity_card_back, from_code, license, expertise, resume, role, is_verified ` +
		`FROM news.users ` +
		`WHERE phone = ?`

	// run query
	XOLog(sqlstr, phone)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, phone).Scan(&u.ID, &u.Nickname, &u.Pwd, &u.Phone, &u.Email, &u.Avatar, &u.Gender, &u.Biography, &u.CreatedAt, &u.LoginDate, &u.RealName, &u.IdentityCardNum, &u.IdentityCardFront, &u.IdentityCardBack, &u.FromCode, &u.License, &u.Expertise, &u.Resume, &u.Role, &u.IsVerified)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func UserByPhoneOrEmail(db XODB, phone, email string) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, nickname, pwd, phone, email, avatar, gender, biography, created_at, login_date, real_name, identity_card_num, identity_card_front, identity_card_back, from_code, license, expertise, resume, role, is_verified ` +
		`FROM news.users ` +
		`WHERE phone = ? OR email = ?`

	// run query
	XOLog(sqlstr, phone, email)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, phone, email).Scan(&u.ID, &u.Nickname, &u.Pwd, &u.Phone, &u.Email, &u.Avatar, &u.Gender, &u.Biography, &u.CreatedAt, &u.LoginDate, &u.RealName, &u.IdentityCardNum, &u.IdentityCardFront, &u.IdentityCardBack, &u.FromCode, &u.License, &u.Expertise, &u.Resume, &u.Role, &u.IsVerified)
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
		`nickname = ?, phone = ?, avatar = ?, gender = ?, biography = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, u.Nickname, u.Phone, u.Avatar, u.Gender, u.Biography, u.ID)
	_, err = db.Exec(sqlstr, u.Nickname, u.Phone, u.Avatar, u.Gender, u.Biography, u.ID)
	return err
}

// UpdatePassword reset password
func (u *User) UpdatePassword(db XODB) error {
	var err error
	// sql query
	const sqlstr = `UPDATE news.users SET ` +
		` pwd= ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, u.Pwd, u.ID)
	_, err = db.Exec(sqlstr, u.Pwd, u.ID)
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
	XOLog(sqlstr, u.RealName.String, u.IdentityCardNum.String, u.IdentityCardFront.String, u.IdentityCardBack.String, u.Expertise.String, u.Resume.String, u.FromCode.String, u.ID)
	_, err = db.Exec(sqlstr, u.RealName.String, u.IdentityCardNum.String, u.IdentityCardFront.String, u.IdentityCardBack.String, u.Expertise.String, u.Resume.String, u.FromCode.String, u.ID)
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

// GetAllExpertUsers 认证专家列表
func GetAllExpertUsers(db *sql.DB, start, count int) ([]User, error) {
	sqlstr := fmt.Sprintf("SELECT id, nickname, pwd, phone, email, avatar, gender, biography, created_at, login_date, real_name, identity_card_num, identity_card_front, identity_card_back, from_code, license, expertise, resume, role, is_verified FROM news.users WHERE is_verified = TRUE AND role = 3 LIMIT %d, %d ", start, count)

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

// GetAllSellerUsers 认证专家列表
func GetAllSellerUsers(db *sql.DB, start, count int) ([]User, error) {
	sqlstr := fmt.Sprintf("SELECT id, nickname, pwd, phone, email, avatar, gender, biography, created_at, login_date, real_name, identity_card_num, identity_card_front, identity_card_back, from_code, license, expertise, resume, role, is_verified FROM news.users WHERE is_verified = TRUE AND role = 2 LIMIT %d, %d ", start, count)

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

// 查询news_item,help_request, buy_request, chips是否被当前用户liked
func IsLikedByUser(db *sql.DB, objType string, objID int, userID int) bool {
	var count int
	var sqlstr string
	switch objType {
	case "news_item":
		sqlstr = fmt.Sprintf("SELECT count(*) from news_items join favorites where favorable_type = 'news_item' and favorable_id = '%d' and user_id = '%d'", objID, userID)
	case "chip":
		sqlstr = fmt.Sprintf("SELECT count(*) from chips join favorites where chips.user_id = favorites.user_id and favorable_type = 'chip' and favorable_id = '%d' and chips.user_id = '%d'", objID, userID)
	case "buy_request":
		sqlstr = fmt.Sprintf("SELECT count(*) from buy_requests join favorites where buy_requests.user_id = favorites.user_id and favorable_type = 'buy_request' and favorable_id = '%d' and buy_requests.user_id = '%d'", objID, userID)
	case "help_request":
		sqlstr = fmt.Sprintf("SELECT count(*) from help_requests join favorites where help_requests.user_id = favorites.user_id and favorable_type = 'help_request' and favorable_id = '%d' and help_requests.user_id = '%d'", objID, userID)
	}

	fmt.Println(sqlstr)

	err := db.QueryRow(sqlstr).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
	}
	if count >= 1 {
		return true
	}
	return false
}
