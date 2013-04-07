package authenticate

import (
	User "msgstory/user"
	"net/http"
	"code.google.com/p/go-mysql-driver/mysql"
)

type authenticate struct {
	email    string
	password string
	auth bool
}

func (a *authenticate) authorize(useremail, userpassword string) bool {
	user := authenticate{}
	user.email := "jasdeepm@gmail.com"
	user.password := "password"
	return true

	
}