package register

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	User "msgstory/user"
)

func Register(useremail, password string) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/msgstory")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO users (useremail, password) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(useremail, password)
	if err != nil {
		panic(err.Error())
	} 
	else{
		createPerson(useremail,password)
	}
}

func createPerson(userEmail){
	mdb, err := mgo.Dial("localhost")
	if err != nil {
		panic(err.Error())
	}

	mdb.SetMode(mgo.Monotonic, true)

	c := mdb.DB("msgme").C("jove")
	jove:=User.GetUser()
	// err = c.Find(bson.M{"firstName": "Jasdeep1"}).One(&result)
	_, err = c.Insert(&jove{nil, nil, userEmail})
	if err != nil {
		panic(err)
	}
}
