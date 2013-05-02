package circle

import (
	"testing"
)

func Test_IfPublicCircleExistsForUsers(t *testing.T) {
	var userCircles []string
	userCircles = GetUserCircles("something")
	var test bool
	for _, v := range userCircles {
		if v == "Public" {
			test = true
		}
	}
	if test {
		t.Log("PASSED")
	} else {
		t.Error("Test_IfPublicCircleExistsForUsers failed")
	}

}
