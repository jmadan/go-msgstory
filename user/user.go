package user

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// var instantiated *user = nil

type user struct {
	Name  string
	Age   int
	Email string
}

func MyUser() user {
	return user{}
}

func GetUser(name string) string {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err.Error())
	}

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("msgme").C("users")

	result := user{}
	err = c.Find(bson.M{"firstName": "Jasdeep1"}).One(&result)
	if err != nil {
		panic(err)
	}

	return result.Email

	// serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	// return
}
