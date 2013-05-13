package main

import (
	"code.google.com/p/gorest"
	"encoding/json"
	"fmt"
	Authenticate "github.com/jmadan/go-msgstory/authenticate"
	Circle "github.com/jmadan/go-msgstory/circle"
	Glocation "github.com/jmadan/go-msgstory/geolocation"
	Mesiji "github.com/jmadan/go-msgstory/message"
	Register "github.com/jmadan/go-msgstory/register"
	User "github.com/jmadan/go-msgstory/user"
	"log"
	"net/http"
	"os"
	"strings"
)

type AppService struct {
	gorest.RestService `root:"/" consumes:"application/json" produces:"application/json"`
	getApp             gorest.EndPoint `method:"GET" path:"/" output:"string"`
}

type UserService struct {
	gorest.RestService `root:"/user-service/" consumes:"application/json" produces:"application/json"`
	getUser            gorest.EndPoint `method:"GET" path:"/user" output:"string"`
	getAll             gorest.EndPoint `method:"GET" path:"/all" output:"string"`
}

type ConversationService struct {
	gorest.RestService `root:"/convoservice/" consumes:"application/json" produces:"application/json"`

	buildConversation gorest.EndPoint `method:"GET" path:"/buildconversation/{convoid:string}" output:"[]byte"`
}

type MsgService struct {
	gorest.RestService `root:"/message/" consumes:"application/json" produces:"application/json"`
	getMessage         gorest.EndPoint `method:"GET" path:"/getmessage" output:"string"`
	postMessage        gorest.EndPoint `method:"POST" path:"/postit/" postdata:"string"`
}

type AuthenticateService struct {
	gorest.RestService `root:"/auth/" consumes:"application/json" produces:"application/json"`
	registerUser       gorest.EndPoint `method:"POST" path:"/register/" postdata:"string"`
	createUser         gorest.EndPoint `method:"GET" path:"/new/{uemail:string}/{pass:string}" output:"string"`
	loginUser          gorest.EndPoint `method:"POST" path:"/login/" postdata:"string"`
}

type CircleService struct {
	gorest.RestService `root:"/circle/" consumes:"application/json" produces:"application/json"`
	createCircle       gorest.EndPoint `method:"POST" path:"/new/" postdata:"string"`
	getCircles         gorest.EndPoint `method:"GET" path:"/circles/" output:"string"`
}

type LocationService struct {
	gorest.RestService     `root:"/location/" consumes:"application/json" produces:"application/json"`
	getLocations           gorest.EndPoint `method:"GET" path:"/near/{place:string}" output:"string"`
	getLocationsWithLatLng gorest.EndPoint `method:"GET" path:"/coordinates/{lat:string}/{lng:string}" output:"string"`
}

// ************Location Service Methods ***********
func (serv LocationService) GetLocations(place string) string {
	fmt.Println(place)
	resp := Glocation.GetVenues("Chelsea,London")
	serv.ResponseBuilder().SetResponseCode(200)
	return resp
}

func (serv LocationService) GetLocationsWithLatLng(lat, lng string) string {
	str := Glocation.GetVenuesWithLatitudeAndLongitude(lat, lng)
	serv.ResponseBuilder().SetResponseCode(200)
	return str
}

//*************Circle Service Methods ***************
func (serv CircleService) CreateCircle(posted string) {
	var str []string
	str = strings.Split(posted, "=")
	msg, err := Circle.CreateCircle(str[1], "", "", nil)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(msg)
		serv.ResponseBuilder().SetResponseCode(200)
	}
}

func (serv CircleService) GetCircles() string {
	return Circle.GetUserCircles("")[0]
}

//*************Authentication Service Methods ***************
func (serv AuthenticateService) RegisterUser(posted string) {
	//var str []string
	// var jsonResp []byte
  fmt.Println(posted)
  jsObject, err := json.Marshal(posted)
  if err != nil {
    log.Println(err.Error())
  }
  fmt.Println(jsObject)
}

func (serv AuthenticateService) CreateUser(uemail, pass string) string {
	Register.Register(uemail, pass)
	return "Executed!!!"
}

func (serv AuthenticateService) LoginUser(posted string) {
	var str []string
	// var jsonResp []byte
	str = strings.Split(posted, "&")
	useremail := strings.SplitAfter(str[0], "=")
	password := strings.SplitAfter(str[1], "=")
	_, err := Authenticate.Login(useremail[1], password[1])
	if err != nil {
		log.Println(err.Error())
		serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	} else {
		serv.ResponseBuilder().SetResponseCode(200)

	}
	return
}

//*************Message Service Methods ***************
func (serv MsgService) GetMessage() string {
	m := "{\"Message\": \"Welcome to Mesiji\"}"
	return m
}

func (serv MsgService) PostMessage(posted string) {
	Mesiji.PostIt(posted)
}

//*************User Service Methods ***************
func (serv UserService) GetUser() string {
	per := "{User:[" + User.GetUser("asd") + "]}"
	return per
	// serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	// return
}

func (serv UserService) GetAll() string {
	fmt.Print("I am here")
	per := "User:[" + User.GetAll() + "]"
	return per
}

//*************App Service Methods ***************
func (serv AppService) GetApp() string {
	m := "{\"Message\": \"Welcome to Mesiji\"}"
	return m
}

func getData(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("inputEmail"))
}

func main() {
	log.Println(os.Getenv("PORT"))
	gorest.RegisterService(new(AppService))
	gorest.RegisterService(new(UserService))
	gorest.RegisterService(new(MsgService))
	gorest.RegisterService(new(AuthenticateService))
	gorest.RegisterService(new(CircleService))
	gorest.RegisterService(new(LocationService))
	http.Handle("/", gorest.Handle())
	// http.HandleFunc("/tempurl", getData)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	// User.GetAll()
}
