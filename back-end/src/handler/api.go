package handler

import (
	"log"
)

type Api struct {
	Log     *log.Logger
	Company *Company
	User    *User
	Registration *Registration
}

func (i *Api) Initialize() {
	// Initialize the logger for the API handler
	i.Company = &Company{log: i.Log}
	i.User = &User{log: i.Log}
	i.Registration = &Registration{log: i.Log}
}