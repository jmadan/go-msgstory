package message

import (
	"testing"
)

var Msg = "{\"msgtext\":\"this is a text message\",\"ownerid\":\"1234567\"}"

func Test_PostIt(t *testing.T) {
	str := Post_Message(Msg)
	if str {
		t.Log("PASSED")
	}
}
