package main

import (
	"code.google.com/p/gorest"
	"encoding/json"
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

type AppService struct {
	gorest.RestService `root:"/"`
	getApp             gorest.EndPoint `method:"GET" path:"/" output:"string"`
}

type UserService struct {
	gorest.RestService `root:"/user-service/"`

	getUser gorest.EndPoint `method:"GET" path:"/getuser/{name:string}" output:"string"`
}

type MsgService struct {
	gorest.RestService `root:"/getmessage/"`

	getMessage gorest.EndPoint `method:"GET" path:"/get-message" output:"string"`
	getLogin   gorest.EndPoint `method:"GET" path:"/login" output:"string"`
}

type RegisterService struct {
	gorest.RestService `root:"/create/"`
	// createUser         gorest.EndPoint `method:"POST" path:"/new/{uemail:string}/{pass:string}" postdata:"person"`
	createUser gorest.EndPoint `method:"GET" path:"/new/{uemail:string}/{pass:string}" output:"string"`
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

func (serv UserService) GetUser(name string) string {
	return User.GetUser(name)

	// serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	// return
}

func (serv AppService) GetApp() string {
	m := testMessage{"Welcome to Message Story"}

	b, err := json.Marshal(m)

	if err != nil {
		return err.Error()
	}

	return string(b)
}

func main() {
	gorest.RegisterService(new(AppService))
	gorest.RegisterService(new(UserService))
	gorest.RegisterService(new(MsgService))
	gorest.RegisterService(new(RegisterService))
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":8080", nil)
}
