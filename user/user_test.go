package user

import (
	"encoding/json"
	"testing"
)

func Test_CreatePerson(t *testing.T) {
	person := User{0, "jasdeep", "jasdeepm@gmail.com", "JD", "07818912893", rels{}}
	res := person.CreateUser()
	if res {
		t.Log("PASSED")
	} else {
		t.Fail()
	}
}

// func Test_GetByHandle(t *testing.T) {
// 	res := GetByHandle("jasdeepm@gmail.com")
// 	if res.Name == "JD" {
// 		t.Log("PASSED")
// 	} else {
// 		t.Fail()
// 	}
// }

func Test_GetUser(t *testing.T) {
	person := User{0, "jasdeep", "jasdeepm@gmail.com", "JD", "07818912893", rels{}}
	personjs, err := json.Marshal(person)
	if err != nil {
		t.Log("error converting to json object")
	}
	if person.GetUser() == string(personjs) {
		t.Log("Test_GetUser PASSED")
	} else {
		t.Fail()
	}
}

func Test_CreateUserLogin(t *testing.T) {
	useremail := "test@test.com"
	password := "password"
	user_id := CreateUserLogin(useremail, password)
	uid := getUserByEmail(useremail)

	if uid == user_id {
		t.Log("Test_CreateUserLogn PASSED")
	} else {
		t.Fail()
		t.Log("Test_CreateUserLogn FAILED")
	}
}
