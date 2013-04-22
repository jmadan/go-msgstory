package user

import (
	"testing"
)

func Test_CreatePerson(t *testing.T) {
	person := User{"jasdeep", 30, "jasdeepm@gmail.com", "JD", nil, "07818912893"}
	res := person.CreateUser()
	if res {
		t.Log("PASSED")
	} else {
		t.Fail()
	}
}

func Test_GetByEmail(t *testing.T) {
	// person := User{"jasdeep", 30, "jasdeepm@gmail.com", "JD", nil, "07818912893"}
	res := GetByEmail("jasdeepm@gmail.com")
	if res.Name == "jasdeep" {
		t.Log("PASSED")
	} else {
		t.Fail()
	}
}

// func Test_GetByHandle(t *testing.T) {
// 	person := User{"jasdeep", 30, "jasdeepm@gmail.com", "JD", nil, "07818912893"}
// res := person.GetHandle()
// if res.Handle == "JD" {
// 	t.Log("PASSED")
// } else {
// 	t.Fail()
// }
// }
