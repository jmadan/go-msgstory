package main

import (
	"code.google.com/p/gorest"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
)

type User struct {
	Id    bson.ObjectId
	Name  string
	Email string
}

type UserService struct {
	gorest.RestService `root:"/user-service/"`

	getUser gorest.EndPoint `method:"GET" path:"/getuser/{name:string}" output:"string"`
}

func (serv UserService) GetUser(name string) string {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err.Error())
	}

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("msgme").C("users")

	result := User{}
	err = c.Find(bson.M{"firstName": "Jasdeep"}).One(&result)
	if err != nil {
		panic(err)
	}

	return result.Email

	// serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	// return
}

func main() {
	gorest.RegisterService(new(UserService))
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":8787", nil)

}
