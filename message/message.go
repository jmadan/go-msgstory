package message

import (
	"encoding/json"
	Connection "github.com/jmadan/go-msgstory/connection"
	RD "github.com/jmadan/go-msgstory/util"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"os"
	"strings"
	"time"
)

type Message struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	MsgText     string        `json:"msg_text" bson:"msg_text"`
	UserId      string        `json:"user_id" bson:"user_id"`
	ParentMsgId string        `json:"parent_msg_id" bson:"parent_msg_id"`
	CreatedOn   time.Time     `json:"created_on" 	bson:"created_on, omitempty"`
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

func (msg *Message) SaveMessage(conversationId string) RD.ReturnData {
	returnData := RD.ReturnData{}
	dbSession := Connection.GetDBSession()
	dbSession.SetMode(mgo.Monotonic, true)
	dataBase := strings.SplitAfter(os.Getenv("MONGOHQ_URL"), "/")
	c := dbSession.DB(dataBase[3]).C("conversation")

	var change = mgo.Change{ReturnNew: true, Update: bson.M{
		"$push": bson.M{"messages": bson.M{
			"msg_text":      msg.MsgText,
			"user_id":       msg.UserId,
			"parent_msg_id": msg.ParentMsgId,
			"created_on":    msg.CreatedOn,
		}}}}
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(conversationId)}, change)
	if err != nil {
		log.Println(err.Error())
		returnData.ErrorMsg = err.Error()
		returnData.Success = false
		returnData.Status = "422"
	} else {
		jsonData := []byte("{}")
		returnData.Success = true
		returnData.JsonData = jsonData
		returnData.Status = "201"
	}
	return returnData
}

func GetMessages(conversationId string) RD.ReturnData {
	returnData := RD.ReturnData{}
	dbSession := Connection.GetDBSession()
	dbSession.SetMode(mgo.Monotonic, true)
	dataBase := strings.SplitAfter(os.Getenv("MONGOHQ_URL"), "/")
	c := dbSession.DB(dataBase[3]).C("conversation")

	Msgs := []Message{}
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(conversationId)}).All(&Msgs)
	if err != nil {
		log.Println(err.Error())
		returnData.ErrorMsg = err.Error()
		returnData.Success = false
		returnData.Status = "422"
	} else {
		jsonData, _ := json.Marshal(&Msgs)
		returnData.Success = true
		returnData.JsonData = jsonData
		returnData.Status = "201"
	}
	return returnData
}
