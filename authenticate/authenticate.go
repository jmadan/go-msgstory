package authenticate

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// User "github.com/jmadan/go-msgstory/user"
	"log"
	"os"
)

type Authenticate struct {
	email           string
	password        string
	user_id         int
	isAuthenticated bool
}

//private function to verify credentials with MySQL
func (a *Authenticate) authorize() {
	// var person User.User
	dburl := os.Getenv("DATABASE_URL")
	// "mysql", "root:password@tcp(localhost:3306)/msgstory"
	db, err := sql.Open("mysql", dburl[8:])
	if err != nil {
		log.Fatal("Phat Gayee : " + err.Error())
	}
	defer db.Close()

	stmtOut, err := db.Prepare("SELECT USER_ID, USEREMAIL FROM USERS WHERE USEREMAIL = ? AND PASSWORD = ?")
	if err != nil {
		log.Fatal("stmtError :" + err.Error())
	}
	defer stmtOut.Close()

	// err = stmtOut.QueryRow(useremail, userpassword).Scan(&authorize.user_id, &authorize.email)
	err = stmtOut.QueryRow(a.email, a.password).Scan(&a.user_id, &a.email)

	if err != nil {
		log.Print("stmtExecution: " + err.Error())
		a.isAuthenticated = false
	} else {
		a.isAuthenticated = true
		log.Println(a.user_id)
	}
}

//Login function to check users credentials
func Login(email, password string) Authenticate {
	a := Authenticate{email, password, 0, false}
	a.authorize()
	return a
}
