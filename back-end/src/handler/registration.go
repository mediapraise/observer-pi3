package handler

import (
	"log"
	"net/http"
	"observer-go/src/service"
)

type Registration struct {
	log                 *log.Logger
	RegistrationService *service.RegistrationService
}

func (r *Registration) Register(res http.ResponseWriter, req *http.Request) {

}
