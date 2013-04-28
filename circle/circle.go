package circle

import (
	"encoding/json"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Circle struct {
	Name       string `json:"name" bson:"name"`
	Descrption string `json:"description" bson:"description"`
	// Owner      *User
	// Members    []*User
}

type JsonCircle struct {
	Name string `json:"name" bson:"name"`
}

func (c *Circle) GetName() string {
	return c.Name
}

// func (cir *circle) GetMembers(name string) (circle, exists bool) {
// 	dbSession, err := mgo.Dial("localhost")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	dbSession.SetMode(mgo.Monotonic, true)
// 	c := dbSession.DB("msgme").C("circle")

// 	result := circle{}
// 	err = c.Find(bson.M{"name": name}).One(&result)
// 	if err.Error() == "not found" {
// 		exists = false
// 	} else {
// 		exists = true
// 	}

// 	return result, exists
// }

func (c *Circle) GetJson() string {
	jCircle := JsonCircle{
		c.Name,
		// c.Descrption,
		// c.Owner,
		// c.Members,
	}

	str, err := json.Marshal(jCircle)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(str)
}

// func CreateCircle(name string, desc string, owner User) string {
func CreateCircle(name, desc string) string {
	msgCircle := Circle{name, desc}

	if CheckIfCircleExists(&msgCircle) {
		return "Duplicate"
	}
	dbSession, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println(err.Error())
	}
	dbSession.SetMode(mgo.Monotonic, true)
	c := dbSession.DB("msgme").C("circle")
	// c := Connection.GetDataBase("circle")

	fmt.Println(msgCircle.GetJson())

	err = c.Insert(msgCircle)
	if err != nil {
		fmt.Println(err.Error())
	}
	return "Done!"
}

// func CheckIfCircleExists(name string, owner User) (exists bool, msg string) {
func CheckIfCircleExists(mCircle *Circle) (exists bool) {
	dbSession, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println(err.Error())
	}

	dbSession.SetMode(mgo.Monotonic, true)
	c := dbSession.DB("msgme").C("circle")

	result := Circle{}
	err = c.Find(bson.M{"name": mCircle.Name}).One(&result)
	if err.Error() == "not found" {
		exists = false
	} else {
		exists = true
	}

	return exists
}
