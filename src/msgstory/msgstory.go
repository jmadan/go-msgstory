package main

import (
	"code.google.com/p/gorest"
	Authenticate "msgstory/authenticate"
	Register "msgstory/register"
	User "msgstory/user"
	"net/http"
)

type person struct {
	Name  string
	Email string
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
	gorest.RestService `root:"/"`
	// createUser         gorest.EndPoint `method:"POST" path:"/new/{uemail:string}/{pass:string}" postdata:"person"`
	createUser gorest.EndPoint `method:"GET" path:"/new/{uemail:string}/{pass:string}" output:"string"`
}

func (serv RegisterService) CreateUser(uemail, pass string) string {
	Register.Register(uemail, pass)
	return "Executed!!!"
}

func (serv MsgService) GetMessage() string {
	return "Welcome to Message Story"
}

func (serv MsgService) GetLogin() string {
	return Authenticate.Login()
}

func (serv UserService) GetUser(name string) string {
	return User.GetUser(name)

	// serv.ResponseBuilder().SetResponseCode(404).Overide(true)
	// return
}

func main() {
	gorest.RegisterService(new(UserService))
	gorest.RegisterService(new(MsgService))
	gorest.RegisterService(new(RegisterService))
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":8080", nil)
}
