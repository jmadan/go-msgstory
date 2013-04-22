package authenticate

import (
	"testing"
)

func Test_Authorize_Incorrect_Credentials(t *testing.T) {

	authenticate := Authenticate{"a@a.com", "something", "", false}
	user, loggedIn := authenticate.Authorize()
	if !loggedIn.isAuthenticated {
		t.Log("PASSED")
		t.Log(user.Email)
	} else {
		t.Log("Test_Authorize_Incorrect_Credentials Failed")
		t.Fail()
	}
}

func Test_Authorize_Correct_Credentials(t *testing.T) {
	authenticate := Authenticate{"jasdeepm@gmail.com", "98036054", "", false}
	user, loggedIn := authenticate.Authorize()
	if loggedIn.isAuthenticated {
		t.Log("PASSED")
		t.Log(user.Email)
	} else {
		t.Log("Test_Authorize_Correct_Credentials Failed")
		t.Fail()
	}
}
