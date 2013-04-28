package conversation

import (
	Msg "go-msgstory/message"
	User "go-msgstory/user"
	"log"
)

type Conversation struct {
	Title        string
	Messages     []*Msg.Message
	Conv_Starter *User
}
