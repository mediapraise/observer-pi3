package handler

import (
	"log"
	"net/http"
)

type Company struct {
	log *log.Logger
}

func (i *Company) CreateCompany(res http.ResponseWriter, req *http.Request) {
	i.log.Println("CreateCompany")
}

func (i *Company) GetCompanyById(res http.ResponseWriter, req *http.Request) {
	i.log.Println("GetCompany")
}

func (i *Company) UpdateCompanyById(res http.ResponseWriter, req *http.Request) {
	i.log.Println("UpdateCompany")
}

func (i *Company) DeleteCompanyById(res http.ResponseWriter, req *http.Request) {
	i.log.Println("DeleteCompany")
}

func (i *Company) GetUsersByCompanyId(res http.ResponseWriter, req *http.Request) {
	i.log.Println("GetUsersByCompany")
}
