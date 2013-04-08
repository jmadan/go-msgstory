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

func (a *Authenticate) Authorize(useremail, userpassword string) string {
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

	// authorize := Authenticate{nil, nil, false}
	// err = stmtOut.QueryRow(useremail, userpassword).Scan(&authorize.user_id, &authorize.email)
	err = stmtOut.QueryRow(useremail, userpassword).Scan(&a.email, &a.user_id)

	if err != nil {
		panic(err.Error())
	} else {
		a.auth = true
	}
	msg := "Not found"
	if a.auth {
		msg = "Logged In"
	}

	return msg
}

func Login() string {
	a := Authenticate{"", "", false}
	return a.Authorize("jasdeepm@gmail.com", "98036054")
}
