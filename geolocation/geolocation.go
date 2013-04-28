package geolocation

import (
	"strings"
)

type GeoLocation struct {
	FourID   string
	Name     string
	Lat      int
	Lng      int
	Distance int
	Postcode string
	City     string
	State    string
	Country  string
}

// foursquare api call 
// https://api.foursquare.com/v2/venues/search?intent=match%20&ll=40.768853,-73.967792%20&query=Creed%20&client_id=052K4OL11WWWBVKTTREAXQVMOEMA0SPERTSKUQLVS1ALQMRP&client_secret=KCQIMSET51UBMQH4LQEHPJCCERQKYJB2J3LYJLS0X02PJQBR
//https://api.foursquare.com/v2/venues/search?near=ewloe&client_id=052K4OL11WWWBVKTTREAXQVMOEMA0SPERTSKUQLVS1ALQMRP&client_secret=KCQIMSET51UBMQH4LQEHPJCCERQKYJB2J3LYJLS0X02PJQBR

func NewLocal() GeoLocation {
	return GeoLocation{}
}

func (gl *GeoLocation) GetVenues(near string) string {
	var FSqrUrl string
	FSqrUrl = "https://api.foursquare.com/v2/venues/search?v=20130417&near=<nearLocation>&client_id=052K4OL11WWWBVKTTREAXQVMOEMA0SPERTSKUQLVS1ALQMRP&client_secret=KCQIMSET51UBMQH4LQEHPJCCERQKYJB2J3LYJLS0X02PJQBR"
	FSqrUrl = strings.Replace(FSqrUrl, "<nearLocation>", near, -1)

	return "this is one location"
}

func GetNearVenues(glocation string) (glocale []string) {
	// venues := []string{"MiddleOfNoWhere", "SomeWhere"}
	gLocal := GeoLocation{}
	venues := []string{gLocal.GetVenues("M46BA")}
	return venues
}
