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
