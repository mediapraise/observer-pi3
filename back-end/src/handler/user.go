package handler

import (
	"log"
	"net/http"
)

type User struct {
	log *log.Logger
}

func (i *User) Login(res http.ResponseWriter, req *http.Request) {
	i.log.Println("Login")
}

func (i *User) CreateUser(res http.ResponseWriter, req *http.Request) {
	i.log.Println("CreateUser")
}

func (i *User) GetUserById(res http.ResponseWriter, req *http.Request) {
	i.log.Println("GetUser")
}

func (i *User) UpdateUserById(res http.ResponseWriter, req *http.Request) {
	i.log.Println("UpdateUser")
}

func (i *User) DeleteUserById(res http.ResponseWriter, req *http.Request) {
	i.log.Println("DeleteUser")
}
