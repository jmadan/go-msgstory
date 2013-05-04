package user

import (
	"encoding/json"
	"fmt"
	Message "github.com/jmadan/go-msgstory/message"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

type name struct {
	Fullname string `json:"fullname" bson:"fullname"`
}

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

func GetUser(name string) string {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err.Error())
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("msgme").C("jove")

	result := User{}
	err = c.Find(bson.M{"email": "test@email.com"}).One(&result)
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

// func CreateUser(user string) {

// 	session, err := mgo.Dial("localhost")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	session.SetMode(mgo.Monotonic, true)
// 	c := session.DB("msgme").C("jove")

// 	newUser := User{}
// 	err = json.Unmarshal([]byte(user), &newUser)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	err = c.Insert(newUser)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Print("hello")
// }

// func (user User) MarshalJSON() ([]byte, error) {
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
	session, err := mgo.Dial("localhost")
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
	session, err := mgo.Dial("localhost")
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
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err.Error())
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
