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
	Event        *Event
	Registration *Registration
	Finance      *Finance
	Db           *db.Database
}

func (i *Api) Initialize() {
	//Initialize repo
	userRepo := repositories.NewUserRepo(i.Db)
	companyRepo := repositories.NewCompanyRepo(i.Db)
	registrationRepo := repositories.NewRegistrationRepo(i.Db)
	eventRepo := repositories.NewEventRepo(i.Db)
	financeRepo := repositories.NewFinanceRepo(i.Db)
	//Initialize Services
	userServ := service.NewUserService(userRepo)
	companyServ := service.NewCompanyService(companyRepo)
	registrationServ := service.NewRegistrationService(registrationRepo)
	eventServ := service.NewEventService(eventRepo)
	financeServ := service.NewHistoryPaymentService(financeRepo)
	//Initialize Handlers
	i.Company = &Company{log: i.Log, CompanyService: companyServ, UserService: userServ}
	i.User = &User{log: i.Log, UserService: userServ}
	i.Registration = &Registration{log: i.Log, RegistrationService: registrationServ}
	i.Event = &Event{log: i.Log, EventService: eventServ}
	i.Finance = &Finance{log: i.Log, HistoryPaymentService: financeServ}
}
