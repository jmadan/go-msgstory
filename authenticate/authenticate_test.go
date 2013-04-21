package authenticate

import (
	"testing"
)

func Test_Authorize_Incorrect_Credentials(t *testing.T) {
	auth := Authenticate{"a@a.com", "something", "", false}
	str := auth.Authorize()
	if str == "Not found" {
		t.Log("PASSED")
	} else {
		t.Fail()
	}
}

func Test_Authorize_Correct_Credentials(t *testing.T) {
	auth := Authenticate{"jasdeepm@gmail.com", "98036054", "", false}
	str := auth.Authorize()
	if str == "Logged In" {
		t.Log("PASSED")
	} else {
		t.Fail()
	}
}
