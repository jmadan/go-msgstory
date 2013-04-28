package authenticate

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	User "user"
)

type Authenticate struct {
	email           string
	password        string
	user_id         string
	isAuthenticated bool
}

//private function to verify credentials with MySQL
func (a *Authenticate) authorize() (User.User, Authenticate) {
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

	// err = stmtOut.QueryRow(useremail, userpassword).Scan(&authorize.user_id, &authorize.email)

	err = stmtOut.QueryRow(a.email, a.password).Scan(&a.user_id, &a.email)

	if err != nil {
		log.Print(err.Error())
		log.Println("Not found")
	} else {
		a.isAuthenticated = true
		log.Println("Logged In")
		person = User.GetByEmail(a.email)
	}

	return person, *a
}

//Login function to check users credentials
func Login(email, password string) User.User {
	a := Authenticate{email, password, "", false}
	res, auth := a.authorize()
	log.Println(auth.isAuthenticated)
	return res
}
