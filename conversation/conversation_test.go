package conversation

import (
	"fmt"
	Circle "github.com/jmadan/go-msgstory/circle"
	Location "github.com/jmadan/go-msgstory/geolocation"
	Msg "github.com/jmadan/go-msgstory/message"
	"testing"
)

var M1 = Msg.Message{
	MsgText: "Hola! how is everyone doing?",
	OwnerID: "SomeOwnerID",
}

var Cir = Circle.Circle{
	Name: "Tapori",
}

var Loc = Location.GeoLocation{
	FourID:   "4bce6383ef109521bd238486",
	Name:     "City of Manchester",
	Contact:  "",
	Address:  "",
	Lat:      42,
	Lng:      -71,
	Distance: 0,
	Postcode: "03104",
	City:     "Manchester",
	State:    "NH",
	Country:  "United States",
}
var Conv = Conversation{
	Title:     "This is a test conversation!",
	Messages:  []Msg.Message{},
	Venue:     Loc,
	Circles:   []Circle.Circle{Cir},
	ConvOwner: "001",
}

func Test_CreateConversation(t *testing.T) {
	res := Conv.CreateConversation()
	if res.Success {
		t.Log("Test_CreateConversation PASSED")
	} else {
		t.Fail()
		t.Log("Test_CreateConversation FAILED")
	}
}

func Test_GetConversationForLocation(t *testing.T) {
	res := GetConversationForLocation("4bce6383ef109521bd238486")
	if res.Success {
		t.Log("Test_GetConversationForLocation PASSED")
	} else {
		t.Fail()
		t.Log("Test_GetConversationForLocation FAILED")
	}
}

func Test_DeleteConversation(t *testing.T) {
	res := DeleteConversation("51bc529e2ffc2c5db5e9b215")
	if res.Success {
		t.Log("Test_DeleteConversation PASSED")
	} else {
		fmt.Println(res.ErrorMsg)
		t.Fail()
		t.Log("Test_DeleteConversation FAILED")
	}
}

func Test_GetConversation(t *testing.T) {
	conId := "51bc529e2ffc2c5db5e9b215"
	res := GetConversation(conId)
	if res.Success {
		t.Log("Test_GetConversation PASSED")
		fmt.Println(string(res.JsonData))
	} else {
		fmt.Println(res.ErrorMsg)
		t.Fail()
		t.Log("Test_GetConversation FAILED")
	}
}
