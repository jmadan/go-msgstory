package circle

import (
	"encoding/json"
	"fmt"
	Connection "github.com/jmadan/go-msgstory/connection"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"time"
)

type Circle struct {
	Name       string   `json:"name" bson:"name"`
	Descrption string   `json:"description" bson:"description"`
	CreatorID  string   `json:"creator" bson:"creator"`
	CreatedOn  string   `json:"createdon" bson:"createdon"`
	Members    []string `json:"members" bson:"members"`
}

type JsonCircle struct {
	Name string `json:"name" bson:"name"`
}

func (c *Circle) GetName() string {
	return c.Name
}

func GetUserCircles(userID string) []string {
	var userCircles []string
	searchResults := []Circle{}
	query := func(c *mgo.Collection) error {
		fn := c.Find(bson.M{"members": userID}).All(&searchResults)
		return fn
	}
	search := func() error {
		return Connection.WithCollection("circle", query)
	}
	err := search()
	if err != nil {
		searchErr := "Database Error"
		log.Println(searchErr)
	}

	for i, v := range searchResults {
		userCircles[i] = v.Name
	}
	return userCircles
}

func GetCircleMembers(circleName string) []string {
	var circleMembers []string
	searchResults := []Circle{}
	query := func(c *mgo.Collection) error {
		fn := c.Find(bson.M{"name": circleName}).All(&searchResults)
		return fn
	}
	search := func() error {
		return Connection.WithCollection("circle", query)
	}
	err := search()
	if err != nil {
		searchErr := "Database Error"
		log.Println(searchErr)
	}

	for i, v := range searchResults {
		circleMembers[i] = v.Name
	}
	return circleMembers
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
func CreateCircle(name, desc, creatorID string, members []string) string {

	msgCircle := Circle{name, desc, creatorID, time.Now().String(), groupMembers}

	if CheckIfCircleExists(&msgCircle) {
		return "Duplicate! Circle already exists"
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
	return "Circle created"
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
