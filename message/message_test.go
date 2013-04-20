package message

import (
	"testing"
)

func Test_PostIt(t *testing.T) {
	str := "{\"msg\":\"this is a text message\", \"circles\":[{\"name\":\"public\"}], \"location\":{\"name\":\"somewhere\"}, \"owner\":{\"phonenumber\":\"1234567\"}}"
	if PostIt(str) == "this is a test message" {
		t.Log("PASSED")
	}
}
