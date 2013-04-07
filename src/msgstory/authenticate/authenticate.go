package authenticate

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Authenticate struct {
	email   string
	user_id string
	auth    bool
}

func Authorize(useremail, userpassword string) string {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/msgstory")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtOut, err := db.Prepare("SELECT USER_ID, USEREMAIL FROM USERS WHERE USEREMAIL = ? AND PASSWORD = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	authorize := new(Authenticate)
	err = stmtOut.QueryRow(useremail, userpassword).Scan(&authorize.user_id, &authorize.email)

	if err != nil {
		panic(err.Error())
	} else {
		authorize.auth = true
	}
	msg := "Not found"
	if authorize.auth {
		msg = "Logged In"
	}

	return msg
}

func Login() string {
	return Authorize("jasdeepm@gmail.com", "98036054")
}
