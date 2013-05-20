package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	Message "github.com/jmadan/go-msgstory/message"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
)

type User struct {
	UserId int    `json:"userid" bson:"userid"`
	Name   string `json:"name" bson:"name"`
	//	Age         int       `json:"age" bson:"age"`
	Email       string `json:"email" bson:"email"`
	Handle      string `json:"handle" bson:"handle"`
	PhoneNumber string `json:"phone" bson:"phone"`
	relations   rels   `json:"relations" bson:"relations"`
}

type rels struct {
	Messages []Message.Message `json:"messages" bson:"messages"`
}

type JSONUser struct {
	UserId int    `json:"userid" bson:"userid"`
	Name   string `json:"name" bson:"name"`
	// Age         int    `json:"age" bson:"age"`
	Email       string `json:"email" bson:"email"`
	Handle      string `json:"handle" bson:"handle"`
	PhoneNumber string `json:"phone" bson:"phone"`
	relations   rels   `json:"relations" bson:"relations"`
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetHandle() string {
	return u.Handle
}

func (u *User) GetMessages() string {
	str, err := json.Marshal(u.relations.Messages)
	if err != nil {
		fmt.Println("what the fuck!")
	}
	return string(str)
}

func (u *User) GetUser() string {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err.Error())
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("msgme").C("jove")

	result := User{}
	err = c.Find(bson.M{"email": u.GetEmail()}).One(&result)
	if err != nil {
		panic(err)
	}
	js, _ := json.Marshal(result)
	// fmt.Printf("%s", js)
	return string(js)
	// serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	// return
}

func GetAll() string {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err.Error())
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("msgme").C("jove")

	result := []User{}
	err = c.Find(nil).Limit(10).All(&result)
	if err != nil {
		panic(err.Error())
	}

	return "hello"
}

func (u *User) MarshalJSON() ([]byte, error) {
	jsonUser := JSONUser{
		u.UserId,
		u.Name,
		u.Email,
		u.Handle,
		u.PhoneNumber,
		u.relations,
	}
	return json.Marshal(jsonUser)
}

func GetByEmailAndUserId(email string, user_id int) (User, error) {
	session, err := mgo.Dial(os.Getenv("MONGOHQ_URL"))
	if err != nil {
		log.Fatal(err.Error())
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("msgme").C("jove")

	result := User{}
	err = c.Find(bson.M{"email": email, "userid": user_id}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	// serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	return result, err
}

func (u *User) GetByHandle() User {
	session, err := mgo.Dial(os.Getenv("MONGOHQ_URL"))
	if err != nil {
		log.Fatal(err.Error())
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("msgme").C("jove")

	result := User{}
	err = c.Find(bson.M{"handle": u.Handle}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
	// serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	// return
}

func (u *User) CreateUser() bool {

	session, err := mgo.Dial(os.Getenv("MONGOHQ_URL"))
	if err != nil {
		log.Fatal("Phat gayee : " + err.Error())
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("msgme").C("jove")

	err = c.Insert(u)
	if err != nil {
		log.Print(err.Error())
		return false
	}
	return true
}

func CreateUserLogin(useremail, password string) string {
	dburl := os.Getenv("DATABASE_URL")
	// "mysql", "root:password@tcp(localhost:3306)/msgstory"
	db, err := sql.Open("mysql", dburl[8:])
	if err != nil {
		log.Fatal("Phat Gayee : " + err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO USERS (USEREMAIL,PASSWORD) VALUES (?,?)")
	if err != nil {
		log.Fatal("stmtError :" + err.Error())
	}
	defer stmtIns.Close()

	// err = stmtOut.QueryRow(useremail, userpassword).Scan(&authorize.user_id, &authorize.email)
	_, err = stmtIns.Exec(useremail, password)
	if err != nil {
		log.Print("stmtExecution: " + err.Error())
	}

	var userid string
	stmtOut, err := db.Prepare("SELECT USER_ID FROM USERS WHERE USEREMAIL=?")
	if err != nil {
		log.Println("stmtError: " + err.Error())
	}

	err = stmtOut.QueryRow(useremail).Scan(&userid)
	if err != nil {
		log.Println(err.Error())
	}

	return userid
}

func getUserByEmail(user_email string) string {
	dburl := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", dburl[8:])
	if err != nil {
		log.Fatal("Phat Gayee : " + err.Error())
	}
	defer db.Close()

	stmtOut, err := db.Prepare("SELECT USER_ID FROM USERS WHERE USEREMAIL = ?")
	if err != nil {
		log.Fatal("stmtError :" + err.Error())
	}
	defer stmtOut.Close()

	var uid string

	// err = stmtOut.QueryRow(useremail, userpassword).Scan(&authorize.user_id, &authorize.email)
	err = stmtOut.QueryRow(user_email).Scan(&uid)

	if err != nil {
		log.Print("stmtExecution: " + err.Error())
		return err.Error()
	} else {
		return uid
	}

	return uid
}
