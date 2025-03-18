package handler

import (
	"log"
	"net/http"
)

type Api struct {
	Log *log.Logger
}

func (i *Api) Login(res http.ResponseWriter, req *http.Request) {
	i.Log.Println("Login")
}

func (i *Api) CreateUser(res http.ResponseWriter, req *http.Request) {
	i.Log.Println("CreateUser")
}

func (i *Api) GetUserById(res http.ResponseWriter, req *http.Request) {
	i.Log.Println("GetUser")
}

func (i *Api) UpdateUserById(res http.ResponseWriter, req *http.Request) {
	i.Log.Println("UpdateUser")
}

func (i *Api) DeleteUserById(res http.ResponseWriter, req *http.Request) {
	i.Log.Println("DeleteUser")
}


func (i *Api) CreateCompany(res http.ResponseWriter, req *http.Request) {
	i.Log.Println("CreateCompany")
}

func (i *Api) GetCompanyById(res http.ResponseWriter, req *http.Request) {
	i.Log.Println("GetCompany")
}

func (i *Api) UpdateCompanyById(res http.ResponseWriter, req *http.Request) {
	i.Log.Println("UpdateCompany")
}

func (i *Api) DeleteCompanyById(res http.ResponseWriter, req *http.Request) {
	i.Log.Println("DeleteCompany")
}

func (i *Api) GetUsersByCompanyId(res http.ResponseWriter, req *http.Request) {
	i.Log.Println("GetUsersByCompany")
}
