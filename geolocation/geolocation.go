package geolocation

import (
	"io/ioutil"
	"log"
	"net/http"
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

const (
	FSQR_CLIENT_ID     = "052K4OL11WWWBVKTTREAXQVMOEMA0SPERTSKUQLVS1ALQMRP"
	FQSR_CLIENT_SECRET = "KCQIMSET51UBMQH4LQEHPJCCERQKYJB2J3LYJLS0X02PJQBR"
)

func NewLocal() GeoLocation {
	return GeoLocation{}
}

func (gl *GeoLocation) GetVenues(near string) string {
	var FSqrUrl string
	FSqrUrl = "https://api.foursquare.com/v2/venues/search?v=20130417&near=<nearLocation>&client_id=" + FSQR_CLIENT_ID + "&client_secret=" + FQSR_CLIENT_SECRET + "\""
	FSqrUrl = strings.Replace(FSqrUrl, "<nearLocation>", near, -1)

	res, err := http.Get(FSqrUrl)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(res.StatusCode))

	return string(res.StatusCode)
}

func (gl *GeoLocation) GetVenuesWithLatitudeAndLongitude(lt, lg string) (string, string) {
	latandlong := lt + "," + lg
	var contents, FSqrUrl string
	FSqrUrl = "https://api.foursquare.com/v2/venues/search?v=20130417&ll=<latandlong>&client_id=" + FSQR_CLIENT_ID + "&client_secret=" + FQSR_CLIENT_SECRET + "\""
	FSqrUrl = strings.Replace(FSqrUrl, "<latandlong>", latandlong, -1)

	response, err := http.Get(FSqrUrl)
	if err != nil {
		log.Println(err.Error())
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf(" err is %s", err)
		}
		log.Printf("contents are %s\n", string(contents))
	}

	return string(contents), string(response.StatusCode)
}

func GetNearVenues(glocation string) (glocale []string) {
	// venues := []string{"MiddleOfNoWhere", "SomeWhere"}
	gLocal := GeoLocation{}
	venues := []string{gLocal.GetVenues("M46BA")}
	return venues
}

func locations(response []string) []*GeoLocation {
	var venues []*GeoLocation

	return venues
}
