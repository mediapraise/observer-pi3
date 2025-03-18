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
	skipPaths := []string{"/user/register", "/user/login"}
	i.Router.Use(func(next http.Handler) http.Handler {
		return middleware.SkipPathsMiddleware(skipPaths, next)
	})

	i.Router.HandleFunc("/user/register", i.Handler.CreateUser).Methods("POST")
	i.Router.HandleFunc("/user/login", i.Handler.Login).Methods("POST")
	i.Router.HandleFunc("/user/{id}", i.Handler.GetUserById).Methods("GET")
	i.Router.HandleFunc("/user/{id}", i.Handler.UpdateUserById).Methods("PUT")
	i.Router.HandleFunc("/user/{id}", i.Handler.DeleteUserById).Methods("DELETE")
	i.Router.HandleFunc("/company", i.Handler.CreateCompany).Methods("POST")
	i.Router.HandleFunc("/company/{id}", i.Handler.GetCompanyById).Methods("GET")
	i.Router.HandleFunc("/company/{id}", i.Handler.UpdateCompanyById).Methods("PUT")
	i.Router.HandleFunc("/company/{id}", i.Handler.DeleteCompanyById).Methods("DELETE")
	i.Router.HandleFunc("/company/{id}/user", i.Handler.GetUsersByCompanyId).Methods("GET")
}
