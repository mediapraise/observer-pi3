package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"observer-go/src/service"
	"observer-go/src/structs/DTO"
)

type Finance struct {
	log                   *log.Logger
	HistoryPaymentService *service.HistoryPaymentService
}

func (i *Finance) CreateHistoryPayment(res http.ResponseWriter, req *http.Request) {
	// Decode request body into DTO
	var historyPaymentBody DTO.HistoryPaymentDTO
	if err := json.NewDecoder(req.Body).Decode(&historyPaymentBody); err != nil {
		i.log.Println("Error decoding request body:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	// Create history Payment via service
	if err := i.HistoryPaymentService.Create(historyPaymentBody); err != nil {
		i.log.Println("Error creating history Payment:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"message": "History Payment created"}`))
}

func (i *Finance) UpdateHistoryPayment(res http.ResponseWriter, req *http.Request) {
	// Parse ID and body
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	var historyPaymentBody DTO.HistoryPaymentDTO
	if err := json.NewDecoder(req.Body).Decode(&historyPaymentBody); err != nil {
		i.log.Println("Error decoding request body:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	historyPaymentBody.ID = id

	// Update via service
	if err := i.HistoryPaymentService.Update(historyPaymentBody); err != nil {
		i.log.Println("Error updating history Payment:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"message": "History Payment updated"}`))
}

func (i *Finance) GetHistoryPaymentById(res http.ResponseWriter, req *http.Request) {
	// Parse ID
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// Get history Payment via service
	historyPaymentDTO, err := i.HistoryPaymentService.GetByID(id)
	if err != nil {
		i.log.Println("Error getting history Payment:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	res.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(historyPaymentDTO)
	if err != nil {
		i.log.Println("Error encoding response:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (i *Finance) GetAllHistoryPayment(res http.ResponseWriter, req *http.Request) {
	// Get all history Payments via service
	historyPayments, err := i.HistoryPaymentService.GetAll()
	if err != nil {
		i.log.Println("Error getting all history Payments:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	res.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(historyPayments)
	if err != nil {
		i.log.Println("Error encoding response:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (i *Finance) GetAllHistoryPaymentByCompanyId(res http.ResponseWriter, req *http.Request) {
	// Parse company ID
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting company ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// Get all history Payments by company ID via service
	historyPayments, err := i.HistoryPaymentService.GetAllByCompanyID(id)
	if err != nil {
		i.log.Println("Error getting all history Payments by company ID:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	res.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(historyPayments)
	if err != nil {
		i.log.Println("Error encoding response:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (i *Finance) DeleteHistoryPaymentById(res http.ResponseWriter, req *http.Request) {
	// Parse ID
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// Delete history Payment via service
	if err := i.HistoryPaymentService.Delete(id); err != nil {
		i.log.Println("Error deleting history Payment:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"message": "History Payment deleted"}`))
}