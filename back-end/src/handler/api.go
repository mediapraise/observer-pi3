package handler

import (
	"log"
	"observer-go/src/db"
	"observer-go/src/repositories"
	"observer-go/src/service"
)

type Api struct {
	Log          *log.Logger
	Company      *Company
	User         *User
	Registration *Registration
	Db           *db.Database
}

func (i *Api) Initialize() {
	//Initialize repo
	userRepo := repositories.NewUserRepo(i.Db)
	companyRepo := repositories.NewCompanyRepo(i.Db)
	registrationRepo := repositories.NewRegistrationRepo(i.Db)
	//Initialize Services
	userServ := service.NewUserService(userRepo)
	companyServ := service.NewCompanyService(companyRepo)
	registrationServ := service.NewRegistrationService(registrationRepo)

	i.Company = &Company{log: i.Log, CompanyService: companyServ, UserService: userServ}
	i.User = &User{log: i.Log, UserService: userServ}
	i.Registration = &Registration{log: i.Log, RegistrationService: registrationServ}
}
