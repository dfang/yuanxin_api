package model

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	return func(t *testing.T) {
		t.Log("teardown test case")
	}
}

func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("setup sub test")
	return func(t *testing.T) {
		t.Log("teardown sub test")
	}
}
func TestUser_InsertUser(t *testing.T) {

}

func TestUser_SignInUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	// rows := sqlmock.NewRows([]string{"id", "title"}).
	// 	AddRow(1, "one").
	// 	AddRow(2, "two")

	// mock.ExpectQuery("SELECT").WillReturnRows(rows)
	// rs, _ := db.Query("SELECT")
	// defer rs.Close()

	// for rs.Next() {
	// 	var id int
	// 	var title string
	// 	rs.Scan(&id, &title)
	// 	fmt.Println("scanned id:", id, "and title:", title)
	// }

	// if rs.Err() != nil {
	// 	fmt.Println("got rows error:", rs.Err())
	// }

	// rows := sqlmock.NewRows([]string{"phone", "pwd"}).AddRow("15618903080", "qaz123./")

	mock.ExpectQuery("SELECT phone, pwd FROM news.users WHERE phone = ? AND pwd = ?").WithArgs("15618903080", "qaz123./").WillReturnRows()
	// mock.ExpectQuery("SELECT (.+) FROM news.users WHERE phone = ? AND pwd = ?").WithArgs("15618903080", "qaz123./").WillReturnError(sqlmock.NewErrorResult())
	// mock.ExpectQuery(sqlstr).WithArgs("15618903080", "qaz123./").WillReturnRows(sqlmock.NewRows([]string{"id", "nickname"}))
	// ExpectExec("Sign In User").WithArgs(db, "15618903080", "qaz123./").WillReturnResult(sqlmock.NewResult(1, 1))

	// user, err := SignInUser(db, "15618903080", "qaz123./")

	// log.Println(user)
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// }

}
