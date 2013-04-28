package conversation

import (
	"log"
	Msg "msgstory/message"
	User "msgstory/user"
)

type Conversation struct {
	Title        string
	Messages     []*Msg.Message
	Conv_Starter *User
}
