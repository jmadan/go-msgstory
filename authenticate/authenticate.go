package authenticate

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	User "msgstory/user"
)

type Authenticate struct {
	email    string
	password string
	user_id  string
	auth     bool
}

func (a *Authenticate) Authorize() user {
	var json_user, response string
	var person User.User

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/msgstory")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	stmtOut, err := db.Prepare("SELECT USER_ID, USEREMAIL FROM USERS WHERE USEREMAIL = ? AND PASSWORD = ?")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stmtOut.Close()

	// authorize := Authenticate{nil, nil, false}
	// err = stmtOut.QueryRow(useremail, userpassword).Scan(&authorize.user_id, &authorize.email)
	err = stmtOut.QueryRow(a.email, a.password).Scan(&a.email, &a.user_id)

	if err != nil {
		log.Print(err.Error())
		response = "Not found"
		log.Print(response)
	} else {
		a.auth = true
		response = "Logged In"
		person.Email = a.email
		json_user = person.GetByEmail()
	}

	return response
}

func Login(email, password string) s {
	a := Authenticate{email, password, "", false}
	res := a.Authorize()
	return res
}
