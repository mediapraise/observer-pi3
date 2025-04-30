package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"observer-go/src/service"
	"observer-go/src/structs/DTO"
)

type Company struct {
	log            *log.Logger
	CompanyService *service.CompanyService
	UserService    *service.UserService
}

func (i *Company) CreateCompany(res http.ResponseWriter, req *http.Request) {
	var companyBody DTO.CompanyDTO
	err := json.NewDecoder(req.Body).Decode(&companyBody)
	if err != nil {
		i.log.Println("Error decoding request body:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	err = i.CompanyService.CreateCompany(companyBody)
	if err != nil {
		i.log.Println("Error creating company:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	i.log.Println("Company created")
	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"message": "Company created"}`))
}

func (i *Company) GetCompanyById(res http.ResponseWriter, req *http.Request) {
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	companyDTO, err := i.CompanyService.GetCompanyByID(id)
	if err != nil {
		i.log.Println("Error getting company:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(companyDTO)
	if err != nil {
		i.log.Println("Error encoding response:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (i *Company) UpdateCompanyById(res http.ResponseWriter, req *http.Request) {
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	var companyBody DTO.CompanyDTO
	err = json.NewDecoder(req.Body).Decode(&companyBody)
	if err != nil {
		i.log.Println("Error decoding request body:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	companyBody.ID = id
	err = i.CompanyService.UpdateCompany(companyBody)
	if err != nil {
		i.log.Println("Error updating company:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	i.log.Println("Company updated")
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"message": "Company updated"}`))
}

func (i *Company) DeleteCompanyById(res http.ResponseWriter, req *http.Request) {
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	err = i.CompanyService.DeleteCompany(id)
	if err != nil {
		i.log.Println("Error deleting company:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	i.log.Println("Company deleted")
	res.WriteHeader(http.StatusNoContent)
}

func (i *Company) GetUsersByCompanyId(res http.ResponseWriter, req *http.Request) {
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	users, err := i.UserService.GetAllUsers()
	if err != nil {
		i.log.Println("Error fetching users:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	var filtered []DTO.UserDTO
	for _, u := range users {
		if u.CompanyID == id {
			filtered = append(filtered, u)
		}
	}
	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(filtered); err != nil {
		i.log.Println("Error encoding response:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}
