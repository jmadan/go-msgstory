package util

import (
	"fmt"
	"testing"
)

var msg = []byte("{\"name\":\"something\"}")

// var msg = []byte("eyJuYW1lIjoic29tZXRoaW5nIn0=")
var some_data = ReturnData{true, "No Error", msg, "200"}

func Test_ToString(t *testing.T) {
	testMsg := some_data.ToString()
	fmt.Println(testMsg)
}
