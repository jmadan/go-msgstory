package geolocation

import (
	"testing"
)

func Test_GetVenuesWithLatitudeAndLongitude(t *testing.T) {
	location := GeoLocation{}
	resbody, status_code := location.GetVenuesWithLatitudeAndLongitude("53.48", "2.24")
	if status_code == "200" {
		t.Log("Passed")
	}
	t.Log(resbody)
}

func Test_GetVenues(t *testing.T) {
	location := GeoLocation{}
	res := location.GetVenues("Manchester")
	if res == "200" {
		t.Log("Passed")
	}

}
