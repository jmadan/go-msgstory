package main

import (
	"code.google.com/p/gorest"
	"encoding/json"
	"fmt"
	Authenticate "msgstory/authenticate"
	Register "msgstory/register"
	User "msgstory/user"
	"net/http"
)

type person struct {
	Name  string
	Email string
}

type testMessage struct {
	Msg string
}

type authenticate struct {
	InputEmail    string
	InputPassword string
}

type AppService struct {
	gorest.RestService `root:"/"`
	getApp             gorest.EndPoint `method:"GET" path:"/" output:"string"`
}

type UserService struct {
	gorest.RestService `root:"/user-service/"`
	// getUser            gorest.EndPoint `method:"GET" path:"/getuser/{name:string}" output:"string"`
	getUser gorest.EndPoint `method:"GET" path:"/user" output:"string"`
	getAll  gorest.EndPoint `method:"GET" path:"/all" output:"string"`
}

type MsgService struct {
	gorest.RestService `root:"/getmessage/"`
	getMessage         gorest.EndPoint `method:"GET" path:"/get-message" output:"string"`
	getLogin           gorest.EndPoint `method:"GET" path:"/login" output:"string"`
}

type RegisterService struct {
	gorest.RestService `root:"/"`
	registerUser       gorest.EndPoint `method:"POST" path:"/register/" postdata:"string"`
	createUser         gorest.EndPoint `method:"GET" path:"/new/{uemail:string}/{pass:string}" output:"string"`
}

func (serv RegisterService) RegisterUser(posted string) {
	fmt.Println(posted)
	fmt.Println("I am here to register!!!")
}

func (serv RegisterService) CreateUser(uemail, pass string) string {
	Register.Register(uemail, pass)
	return "Executed!!!"
}

func (serv MsgService) GetMessage() string {
	m := testMessage{"here it is!"}

	b, err := json.Marshal(m)

	if err != nil {
		return err.Error()
	}

	return string(b)
}

func (serv MsgService) GetLogin() string {
	return Authenticate.Login()
}

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

func (serv AppService) GetApp() string {

	m := testMessage{"Welcome to Message Story"}

	b, err := json.Marshal(m)

	if err != nil {
		return err.Error()
	}

	return string(b)
}

func getData(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("inputEmail"))
}

func main() {
	gorest.RegisterService(new(AppService))
	gorest.RegisterService(new(UserService))
	gorest.RegisterService(new(MsgService))
	gorest.RegisterService(new(RegisterService))
	http.Handle("/", gorest.Handle())
	http.HandleFunc("/tempurl", getData)
	http.ListenAndServe(":8080", nil)
	// User.GetAll()
}
