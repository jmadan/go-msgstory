package conversation

import (
	// "fmt"
	Circle "github.com/jmadan/go-msgstory/circle"
	Location "github.com/jmadan/go-msgstory/geolocation"
	// Msg "github.com/jmadan/go-msgstory/message"
	Msg "go-msgstory/message"
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
	Messages:  ,
	Venue:     Loc,
	Group:     []Circle.Circle{Cir},
	ConvOwner: "001",
}

func Test_CreateConversation(t *testing.T) {

	res := Conv.CreateConversation()
	if res == "201" {
		t.Log("Test_CreateConversation PASSED")
	} else {
		t.Fail()
		t.Log("Test_CreateConversation FAILED")
	}
}

func Test_GetConversationForGroup(t *testing.T) {

}
