package main

import (
	"code.google.com/p/gorest"
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
	"strconv"
	"strings"
)

type FormData struct {
	UserName string
	UserPass string
}

type AppService struct {
	gorest.RestService `root:"/" consumes:"application/json" produces:"application/json"`
	getApp             gorest.EndPoint `method:"GET" path:"/" output:"string"`
}

type UserService struct {
	gorest.RestService `root:"/users/" consumes:"application/json" produces:"application/json"`

	getUser      gorest.EndPoint `method:"GET" path:"/{userid:string}" output:"string"`
	getAll       gorest.EndPoint `method:"GET" path:"/" output:"string"`
	registerUser gorest.EndPoint `method:"POST" path:"/" postdata:"string"`
	// createUser         gorest.EndPoint `method:"GET" path:"/new/{uemail:string}/{pass:string}" output:"string"`
}

type ConversationService struct {
	gorest.RestService `root:"/conversations/" consumes:"application/json" produces:"application/json"`

	createConvo   gorest.EndPoint `method:"POST" path:"/" postdata:"string"`
	getAllConvo   gorest.EndPoint `method:"GET" path:"/{placeId:string}" output:"[]byte"`
	getConvo      gorest.EndPoint `method:"GET" path:"/{convoId:string}" output:"[]byte"`
	putMessage    gorest.EndPoint `method:"POST" path:"/{convoId:string}/messages/" postdata:"string"`
	deleteConvo   gorest.EndPoint `method:"DELETE" path:"/{convoId:string}/"`
	deleteMessage gorest.EndPoint `method:"DELETE" path:"/{convoId:string}/messages/{msgId:string}"`
}

type MsgService struct {
	gorest.RestService `root:"/messages/" consumes:"application/json" produces:"application/json"`
	getMessage         gorest.EndPoint `method:"GET" path:"/" output:"string"`
	postMessage        gorest.EndPoint `method:"POST" path:"/postit/" postdata:"string"`
}

type AuthenticateService struct {
	gorest.RestService `root:"/auth/" consumes:"application/json" produces:"application/json"`
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

func (serv AuthenticateService) LoginUser(posted string) {
	var str []string
	log.Println(posted)
	// var jsonResp []byte
	str = strings.Split(posted, "&")
	useremail := strings.SplitAfter(str[0], "=")
	password := strings.SplitAfter(str[1], "=")
	auth := Authenticate.Login(useremail[1], password[1])
	if auth.IsAuthenticated {
		user := User.User{}
		user.UserId = auth.User_id
		user.Email = auth.Email
		serv.ResponseBuilder().SetResponseCode(200).Write([]byte(user.GetUser()))
		return
	} else {
		res := "{\"error\":\"authentication failed\"}"
		serv.ResponseBuilder().SetResponseCode(404).WriteAndOveride([]byte(res))
		return
	}
}

//*************Message Service Methods ***************
func (serv MsgService) GetMessage() string {
	m := "{\"Message\": \"Welcome to Mesiji\"}"
	return m
}

func (serv MsgService) PostMessage(posted string) {
	Mesiji.Save_Message(posted)
}

//*************User Service Methods ***************
func (serv UserService) RegisterUser(posted string) {
	var formData []string
	var questionmark int
	// var jsonObject string
	if strings.Contains(posted, "?") {
		questionmark = strings.Index(posted, "?")
	}

	if questionmark == 0 {
		formData = strings.Split(posted[1:], "&")
	} else {
		formData = strings.Split(posted, "&")
	}

	// jsonObject = "{"
	// for i := 0; i < len(formData); i++ {
	// 	jsonObject += "\"" + formData[i][:strings.Index(formData[i], "=")] + "\":\"" + formData[i][strings.Index(formData[i], "=")+1:] + "\""
	// 	if i != len(formData)-1 {
	// 		jsonObject += ","
	// 	}
	// }
	// jsonObject += "}"
	// jsonObject = "{\"" + formData[0][:strings.Index(formData[0], "=")] + "\":\"" + formData[0][strings.Index(formData[0], "=")+1:] + "\""
	// jsonObject += "\"" + formData[1][:strings.Index(formData[1], "=")] + "\":\"" + formData[1][strings.Index(formData[1], "=")+1:] + "\"}"
	user_id := User.CreateUserLogin(formData[1][strings.Index(formData[1], "=")+1:], formData[3][strings.Index(formData[3], "=")+1:])
	user := User.User{}
	user.UserId, _ = strconv.Atoi(user_id)
	user.Name = formData[0][strings.Index(formData[0], "=")+1:]
	user.Email = formData[1][strings.Index(formData[1], "=")+1:]
	user.Handle = formData[2][strings.Index(formData[2], "=")+1:]
	user.CreateUser()
	log.Println(user)
}

func (serv UserService) CreateUser(uemail, pass string) string {
	Register.Register(uemail, pass)
	return "Executed!!!"
}

func (serv UserService) GetUser(userid string) string {
	// user := User.User{}
	// per := "{User:[" + User.User.GetUser() + "]}"
	// serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	return "Some User"
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
	// gorest.RegisterService(new(MsgService))
	gorest.RegisterService(new(AuthenticateService))
	gorest.RegisterService(new(CircleService))
	gorest.RegisterService(new(LocationService))
	http.Handle("/", gorest.Handle())
	// http.HandleFunc("/tempurl", getData)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	// User.GetAll()
}

func getResponse() string {
	log.Println("something works")
	return "All is Well"
}
