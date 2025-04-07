package handler

import (
	"log"
	"net/http"
)

type Registration struct {
	log *log.Logger
}

func (r *Registration) Register(res http.ResponseWriter, req *http.Request) {

}
