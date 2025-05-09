package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"observer-go/src/service"
	"observer-go/src/structs/DTO"
)

type Registration struct {
	log                 *log.Logger
	RegistrationService *service.RegistrationService
}

func (i *Registration) CreateRegister(res http.ResponseWriter, req *http.Request) {
	// Decode request body into DTO
	var regBody DTO.RegistrationDTO
	if err := json.NewDecoder(req.Body).Decode(&regBody); err != nil {
		i.log.Println("Error decoding request body:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	// Create registration via service
	if err := i.RegistrationService.CreateRegistration(regBody); err != nil {
		i.log.Println("Error creating registration:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"message": "Registration created"}`))
}

func (i *Registration) UpdateRegister(res http.ResponseWriter, req *http.Request) {
	// Parse ID and body
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	var regBody DTO.RegistrationDTO
	if err := json.NewDecoder(req.Body).Decode(&regBody); err != nil {
		i.log.Println("Error decoding request body:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	regBody.ID = id

	// Update via service
	if err := i.RegistrationService.UpdateRegistration(regBody); err != nil {
		i.log.Println("Error updating registration:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"message": "Registration updated"}`))
}

func (i *Registration) GetRegisterById(res http.ResponseWriter, req *http.Request) {
	// Parse ID from request
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve registration via service
	regDTO, err := i.RegistrationService.GetRegistrationByID(id)
	if err != nil {
		i.log.Println("Error fetching registration:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return registration data
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(regDTO)
	res.WriteHeader(http.StatusOK)
}

func (i *Registration) GetRegisterByUserId(res http.ResponseWriter, req *http.Request) {
	// Parse user ID
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting user ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// Fetch all and filter by Owner (user)
	all, err := i.RegistrationService.GetAllRegistrations()
	if err != nil {
		i.log.Println("Error fetching registrations:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	var filtered []DTO.RegistrationDTO
	for _, r := range all {
		if r.Owner == fmt.Sprint(id) {
			filtered = append(filtered, r)
		}
	}

	// Return filtered registrations
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(filtered)
	res.WriteHeader(http.StatusOK)
}

func (i *Registration) GetRegistersByCompanyId(res http.ResponseWriter, req *http.Request) {
	// Parse company ID
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting company ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// Fetch all and filter by CompanyID
	all, err := i.RegistrationService.GetAllRegistrations()
	if err != nil {
		i.log.Println("Error fetching registrations:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	var filtered []DTO.RegistrationDTO
	for _, r := range all {
		if r.CompanyID == fmt.Sprint(id) {
			filtered = append(filtered, r)
		}
	}

	// Return filtered list
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(filtered)
	res.WriteHeader(http.StatusOK)
}

func (i *Registration) DeleteRegisterById(res http.ResponseWriter, req *http.Request) {
	// Parse ID from request
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// Delete registration via service
	if err := i.RegistrationService.DeleteRegistration(id); err != nil {
		i.log.Println("Error deleting registration:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"message": "Registration deleted"}`))
}