package message

import (
	"fmt"
	"testing"
)

var Msg = "{\"msg_text\":\"this is a text message\",\"user_id\":\"1234567\"}"

func Test_GetMessages(t *testing.T) {
	res := GetMessages("51efb0be5fb5e52f2860a052")
	if res.Success {
		fmt.Println(string(res.JsonData))
		t.Log("Test_GetConversationsForLocation PASSED")
	} else {
		t.Fail()
		t.Log("Test_GetConversationsForLocation FAILED")
	}
}

func Test_SaveMessage(t *testing.T) {
	var m Message
	m.JsonToMsg(Msg)
	res := m.SaveMessage("51efb0be5fb5e52f2860a052")
	if res.Success {
		fmt.Println(string(res.JsonData))
		t.Log("Test_SaveMessage PASSED")
	} else {
		t.Fail()
		t.Log("Test_SaveMessage FAILED")
	}
}
