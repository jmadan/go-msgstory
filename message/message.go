package message

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Msg        string
	CircleID   []string
	LocationID string
	OwnerID    string
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
