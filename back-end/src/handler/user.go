package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"observer-go/src/service"
	"observer-go/src/structs/DTO"
	"strconv"
)

type User struct {
	UserService *service.UserService
	log         *log.Logger
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
	idInt, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := i.UserService.GetUserByID(uint(idInt))
	if err != nil {
		i.log.Println("Error getting user:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(user)
	if err != nil {
		i.log.Println("Error encoding response:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (i *User) UpdateUserById(res http.ResponseWriter, req *http.Request) {
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	i.log.Println(id)
}

func (i *User) DeleteUserById(res http.ResponseWriter, req *http.Request) {
	i.log.Println("DeleteUser")
}

func getIdValue(req *http.Request) (uint, error) {
	id := req.URL.Query().Get("id")
	if id == "" {
		return 0, fmt.Errorf("missing id")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return uint(idInt), nil
}
