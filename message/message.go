package message

import (
	"encoding/json"
	"labix.org/v2/mgo/bson"
	"log"
)

type Message struct {
	Id 			bson.ObjectId `json:"id" bson:"id"`
	MsgText     string `json:"msg_text" bson:"msg_text"`
	UserId      string `json:"user_id" bson:"user_id"`
	ParentMsgId string `json:"parent_msg_id" bson:"parent_msg_id"`
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