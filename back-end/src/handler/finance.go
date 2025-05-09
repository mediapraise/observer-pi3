package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"observer-go/src/structs/DTO"

)

type Finance struct {
	log                   *log.Logger
	HistoryPaymantService *service.HistoryPaymantService
}

func (i *Finance) CreateHistoryPaymant(res http.ResponseWriter, req *http.Request) {
	// Decode request body into DTO
	var historyPaymantBody DTO.HistoryPaymentDTO
	if err := json.NewDecoder(req.Body).Decode(&historyPaymantBody); err != nil {
		i.log.Println("Error decoding request body:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	// Create history paymant via service
	if err := i.HistoryPaymantService.CreateHistoryPaymant(historyPaymantBody); err != nil {
		i.log.Println("Error creating history paymant:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"message": "History paymant created"}`))
}

func (i *Finance) UpdateHistoryPaymant(res http.ResponseWriter, req *http.Request) {
	// Parse ID and body
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	var historyPaymantBody DTO.HistoryPaymentDTO
	if err := json.NewDecoder(req.Body).Decode(&historyPaymantBody); err != nil {
		i.log.Println("Error decoding request body:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	historyPaymantBody.ID = id

	// Update via service
	if err := i.HistoryPaymantService.UpdateHistoryPaymant(historyPaymantBody); err != nil {
		i.log.Println("Error updating history paymant:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"message": "History paymant updated"}`))
}
