package conversation

import (
	// "encoding/json"
	// SJ "github.com/bitly/go-simplejson"
	Group "github.com/jmadan/go-msgstory/circle"
	Connection "github.com/jmadan/go-msgstory/connection"
	Location "github.com/jmadan/go-msgstory/geolocation"
	Msg "github.com/jmadan/go-msgstory/message"
	// User "github.com/jmadan/go-msgstory/user"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
	"strings"
)

type Conversation struct {
	Title       string               `json:"title" bson:"title"`
	Messages    []Msg.Message        `json:"messages" bson:"messages"`
	Venue       Location.GeoLocation `json:"venue" bson:"venue"`
	Circles     []Group.Circle       `json:"circles" bson:"circles"`
	ConvOwner   string               `json:"creator" bson:"creator"`
	Is_Approved bool                 `json:"is_approved" bson:"is_approved"`
	Created_On  string               `json:"created_on" bson:"created_on"`
}

func (conv *Conversation) CreateConversation() string {
	dbSession := Connection.GetDBSession()
	dbSession.SetMode(mgo.Monotonic, true)
	dataBase := strings.SplitAfter(os.Getenv("MONGOHQ_URL"), "/")
	c := dbSession.DB(dataBase[3]).C("conversation")

	err := c.Insert(&conv)
	if err != nil {
		log.Print(err.Error())
		return "422"
	}

	return "201"
}

func GetConversationForLocation(convID string) Conversation {
	dbSession := Connection.GetDBSession()
	dbSession.SetMode(mgo.Monotonic, true)
	dataBase := strings.SplitAfter(os.Getenv("MONGOHQ_URL"), "/")
	c := dbSession.DB(dataBase[3]).C("conversation")

	res := Conversation{}
	err := c.Find(bson.M{"_id": convID}).One(&res)
	if err != nil {
		log.Println("Found Nothing. Something went wrong fetching the Conversation document")
	}
	return res
}

func GetConversationForGroup(groupID string) Conversation {
	dbSession := Connection.GetDBSession()
	dbSession.SetMode(mgo.Monotonic, true)
	dataBase := strings.SplitAfter(os.Getenv("MONGOHQ_URL"), "/")
	c := dbSession.DB(dataBase[3]).C("conversation")

	res := Conversation{}
	err := c.Find(bson.M{"Group.Id": groupID}).One(&res)
	if err != nil {
		log.Println("Found Nothing. Something went wrong fetching the Conversation document")
	}
	return res
}
