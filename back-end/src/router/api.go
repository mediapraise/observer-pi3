package router

import (
	"log"
	"net/http"
	"observer-go/src/handler"
	"observer-go/src/middleware"

	"github.com/gorilla/mux"
)

type Api struct {
	Log     *log.Logger
	Router  *mux.Router
	Handler *handler.Api
}

func (i *Api) Build() {
	i.Handler.Initialize()
	skipPaths := []string{"/api/user/register", "api/user/login", "/api/company"}
	i.Router.Use(func(next http.Handler) http.Handler {
		return middleware.SkipPathsMiddleware(skipPaths, next)
	})
	i.user()
	i.company()
	i.events()
}

func (i *Api) registration() {
	i.Router.HandleFunc("/registration", i.Handler.Registration.CreateRegister).Methods("POST")
	i.Router.HandleFunc("/registration/{id}", i.Handler.Registration.UpdateRegister).Methods("PUT")
	i.Router.HandleFunc("/registration/{id}", i.Handler.Registration.DeleteRegisterById).Methods("DELETE")
	i.Router.HandleFunc("/registration/{id}", i.Handler.Registration.GetRegisterById).Methods("GET")
}

func (i *Api) user() {
	i.Router.HandleFunc("/user/register", i.Handler.User.CreateUser).Methods("POST")
	i.Router.HandleFunc("/user/login", i.Handler.User.Login).Methods("POST")
	i.Router.HandleFunc("/user/{id}", i.Handler.User.GetUserById).Methods("GET")
	i.Router.HandleFunc("/user/{id}", i.Handler.User.UpdateUserById).Methods("PUT")
	i.Router.HandleFunc("/user/{id}", i.Handler.User.DeleteUserById).Methods("DELETE")
}

func (i *Api) company() {
	i.Router.HandleFunc("/company", i.Handler.Company.CreateCompany).Methods("POST")
	i.Router.HandleFunc("/company/{id}", i.Handler.Company.GetCompanyById).Methods("GET")
	i.Router.HandleFunc("/company/{id}", i.Handler.Company.UpdateCompanyById).Methods("PUT")
	i.Router.HandleFunc("/company/{id}", i.Handler.Company.DeleteCompanyById).Methods("DELETE")
	i.Router.HandleFunc("/company/{id}/user", i.Handler.Company.GetUsersByCompanyId).Methods("GET")
}

func (i *Api) events() {
	i.Router.HandleFunc("/events", i.Handler.Event.GetEvents).Methods("GET")
	i.Router.HandleFunc("/events/{id}", i.Handler.Event.GetEventById).Methods("GET")
	i.Router.HandleFunc("/events", i.Handler.Event.CreateEvent).Methods("POST")
	i.Router.HandleFunc("/events/{id}", i.Handler.Event.UpdateEventById).Methods("PUT")
	i.Router.HandleFunc("/events/{id}", i.Handler.Event.DeleteEventById).Methods("DELETE")
}
