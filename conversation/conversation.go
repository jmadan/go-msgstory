package conversation

import (
	"log"
	Msg "message"
	User "user"
)

type Conversation struct {
	Title        string
	Messages     []*Msg.Message
	Conv_Starter *User
}
