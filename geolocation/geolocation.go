package geolocation

import (
	// simplejson "github.com/bitly/go-simplejson"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type GeoLocation struct {
	FourID   string
	Name     string
	Contact  string
	Address  string
	Lat      int
	Lng      int
	Distance int
	Postcode string
	City     string
	State    string
	Country  string
}

type Feed struct {
	Gmeta     Metaf    `json:"meta"`
	Gresponse Response `json:"response"`
}

type Response struct {
	Rvenue []Venue `json:"venues"`
}

type Venue struct {
	Id        string   `xml:"id"`
	Name      string   `xml:"name"`
	Gcontact  Contact  `json:"contact"`
	Glocation Location `json:"location"`
	// canonicalUrl string
	// Categories   []Category
	// Verified     bool
	// Restricted   bool
	// Stats        Stat
	// Url          string
	// ReferralId   string
}

type Contact struct {
	Phone          float64 `xml:"phone"`
	FormattedPhone string  `xml:"formatedPhone"`
}

type Location struct {
	Address    string  `xml:"address"`
	Lat        float64 `xml:"lat"`
	Lng        float64 `xml:"lng"`
	PostalCode string  `xml:"postalCode"`
	City       string  `xml:"city"`
	State      string  `xml:"state"`
	Country    string  `xml:"country"`
	CC         string  `xml:"cc"`
}

type Stat struct {
	CheckinsCount float64
	UserCount     float64
	TipCount      float64
}

type Special struct {
	Count float64
	Items []string
}

type Metaf struct {
	Code float64 `xml:"code"`
}
type Category struct {
	Id         string
	Name       string
	PluralName string
	ShortName  string
	Logo       Icon
	Primary    bool
}

type Icon struct {
	Prefix string
	Suffix string
}

const (
	FSQR_CLIENT_ID     = "052K4OL11WWWBVKTTREAXQVMOEMA0SPERTSKUQLVS1ALQMRP"
	FQSR_CLIENT_SECRET = "KCQIMSET51UBMQH4LQEHPJCCERQKYJB2J3LYJLS0X02PJQBR"
)

func GetVenues(near string) []Venue {
	FSqrUrl := "https://api.foursquare.com/v2/venues/search?v=20130417&near=<nearLocation>&client_id=" + FSQR_CLIENT_ID + "&client_secret=" + FQSR_CLIENT_SECRET
	FSqrUrl = strings.Replace(FSqrUrl, "<nearLocation>", near, -1)
	log.Println(FSqrUrl)

	return getLocations(FSqrUrl)
}

func GetVenuesWithLatitudeAndLongitude(lt, lg string) []Venue {
	latandlong := lt + "," + lg
	var FSqrUrl string
	FSqrUrl = "https://api.foursquare.com/v2/venues/search?v=20130417&ll=<latandlong>&client_id=" + FSQR_CLIENT_ID + "&client_secret=" + FQSR_CLIENT_SECRET + "\""
	FSqrUrl = strings.Replace(FSqrUrl, "<latandlong>", latandlong, -1)

	return getLocations(FSqrUrl)
}

func getLocations(FSqrUrl string) []Venue {
	res, err := http.Get(FSqrUrl)
	if err != nil {
		log.Println("getVenues error: " + err.Error())
	}
	defer res.Body.Close()

	var jsonFeed *Feed
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&jsonFeed); err != nil {
		log.Println(err.Error())
	}

	locations := make([]Venue, len(jsonFeed.Gresponse.Rvenue))
	for i, tvenue := range jsonFeed.Gresponse.Rvenue {
		locations[i].Id = tvenue.Id
		locations[i].Name = tvenue.Name
		locations[i].Gcontact.Phone = tvenue.Gcontact.Phone
		locations[i].Gcontact.FormattedPhone = tvenue.Gcontact.FormattedPhone
		// locations[i].Glocation = tvenue.Glocation
	}
	return locations
}

func locations(response []byte) {
	var object *Response
	// log.Println(string(response))
	// u := &FsqrLocation{}
	err := json.Unmarshal(response, &object)
	// err := simplejson.Unmarshal(response, &u)
	if err != nil {
		panic(err)
	}
	log.Println(object.Rvenue)
	// jsonObject := object.(map[string]interface{})

	// Print out mother and father
	// log.Println(Util.JsonObjectAsString(jsonObject))
}
