package user

import (
	"encoding/json"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type name struct {
	Firstname string `json:"firstname" bson:"firstname"`
	Lastname  string `json:"lastname" bson:"lastname"`
}

type message struct {
	Msgtext  string `json:"messageText" bson:"messageText"`
	Circle   string `json:"circle" bson:"circle"`
	Location string `json:"location" bson:"location"`
}

type User struct {
	Name     name      `json:"name" bson:"name"`
	Age      string    `json:"age" bson:"age"`
	Email    string    `json:"email" bson:"email"`
	Handle   string    `json:"handle" bson:"handle"`
	Messages []message `json:"messages" bson:"messages"`
}

type JSONUser struct {
	Name     name      `json:"name" bson:"name"`
	Age      string    `json:"age" bson:"age"`
	Email    string    `json:"email" bson:"email"`
	Handle   string    `json:"handle" bson:"handle"`
	Messages []message `json:"messages" bson:"messages"`
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

func GetUserByEmail(email string) string {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err.Error())
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("msgme").C("jove")

	result := User{}
	err = c.Find(bson.M{"email": email}).One(&result)
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

func CreateUser(user string) {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err.Error())
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("msgme").C("jove")

	newUser := User{}
	err = json.Unmarshal([]byte(user), &newUser)
	if err != nil {
		panic(err.Error())
	}

	err = c.Insert(newUser)
	if err != nil {
		panic(err.Error())
	}
	fmt.Print("hello")
}

// func (user User) MarshalJSON() ([]byte, error) {
func (user User) MarshalJSON() {
	jsonUser := JSONUser{
		user.Name,
		user.Age,
		user.Email,
		user.Handle,
		user.Messages,
	}
	json.Marshal(jsonUser)
	// fmt.Println(len(st))
}
