package message

import (
	"encoding/json"
	Connection "github.com/jmadan/go-msgstory/connection"
	"labix.org/v2/mgo"
	"log"
	// "os"
)

type Message struct {
	MsgText string
	OwnerID string
}

func (M *Message) MsgToJSON() string {
	mjson, err := json.Marshal(M)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(mjson)
}

func (M *Message) JsonToMsg(msgtext string) {
	err := json.Unmarshal([]byte(msgtext), &M)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Save_Message(msgtext string) bool {
	var tempMsg Message
	tempMsg.JsonToMsg(msgtext)

	// err := json.Unmarshal([]byte(msgtext), &tempMsg)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// session, err := mgo.Dial(os.Getenv("MONGOHQ_URL"))
	session := Connection.GetDBSession()
	// if err != nil {
	// 	log.Fatal("Phat gayee : " + err.Error())
	// }

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("msgme").C("message")

	err := c.Insert(&tempMsg)
	if err != nil {
		log.Print(err.Error())
		return false
	}
	return true
}
