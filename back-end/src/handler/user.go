package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"observer-go/src/service"
	"observer-go/src/structs/DTO"

)

type User struct {
	UserService service.UserService
	log *log.Logger
}

func (i *User) Login(res http.ResponseWriter, req *http.Request) {
	i.log.Println("Login")
}

func (i *User) CreateUser(res http.ResponseWriter, req *http.Request) {
	var userBody DTO.UserDTO
	err := json.NewDecoder(req.Body).Decode(&userBody)
	if err != nil {
		i.log.Println("Error decoding request body:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	err = i.UserService.CreateUser(userBody)
	if err != nil {
		i.log.Println("Error creating user:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	i.log.Println("User created")
	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"message": "User created"}`))
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
