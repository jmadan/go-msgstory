package message

import (
	"encoding/json"
	"fmt"
	Circle "github.com/jmadan/go-msgstory/circle"
	Venue "github.com/jmadan/go-msgstory/geolocation"
	User "github.com/jmadan/go-msgstory/user"
	"log"
)

type Message struct {
	Msg      string
	Circles  []*Circle.Circle
	Location *Venue.GeoLocation
	Owner    *User.User
}

func (M *Message) MsgToJSON() string {
	mjson, err := json.Marshal(M)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(mjson)
}

func PostIt(msgtext string) string {
	var tempMsg Message
	err := json.Unmarshal([]byte(msgtext), &tempMsg)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(tempMsg)

	return tempMsg.Msg
}
