package user

import (
	"encoding/json"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	// GeoLocation "msgstory/geolocation"
)

type name struct {
	Fullname string `json:"fullname" bson:"fullname"`
}

type message struct {
	Msgtext  string `json:"messageText" bson:"messageText"`
	Circle   string `json:"circle" bson:"circle"`
	Location string `json:"location" bson:"location"`
}

type User struct {
	Name        string    `json:"name" bson:"name"`
	Age         int       `json:"age" bson:"age"`
	Email       string    `json:"email" bson:"email"`
	Handle      string    `json:"handle" bson:"handle"`
	Messages    []message `json:"messages" bson:"messages"`
	PhoneNumber string    `json:"phone" bson:"phone"`
}

type JSONUser struct {
	Name        string    `json:"name" bson:"name"`
	Age         int       `json:"age" bson:"age"`
	Email       string    `json:"email" bson:"email"`
	Handle      string    `json:"handle" bson:"handle"`
	Messages    []message `json:"messages" bson:"messages"`
	PhoneNumber string    `json:"phone" bson:"phone"`
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
	str, err := json.Marshal(u.Messages)
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

// func GetUserByEmail(email string) string {
// 	session, err := mgo.Dial("localhost")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	session.SetMode(mgo.Monotonic, true)
// 	c := session.DB("msgme").C("jove")

// 	result := User{}
// 	err = c.Find(bson.M{"email": email}).One(&result)
// 	if err != nil {
// 		panic(err)
// 	}
// 	js, _ := json.Marshal(result)
// 	// fmt.Printf("%s", js)
// 	return string(js)
// 	// serv.ResponseBuilder().SetResponseCode(404).Overide(true)
// 	// return
// }

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
		u.Name,
		u.Age,
		u.Email,
		u.Handle,
		u.Messages,
		u.PhoneNumber,
	}
	return json.Marshal(jsonUser)
}

func (u *User) GetByEmail() User {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err.Error())
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("msgme").C("jove")

	result := User{}
	err = c.Find(bson.M{"email": u.Email}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	// serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	return result
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
