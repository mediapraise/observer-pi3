package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"observer-go/src/service"
	"observer-go/src/structs/DTO"
)

type Event struct {
	log          *log.Logger
	EventService *service.EventService
}

func (i *Event) CreateEvent(res http.ResponseWriter, req *http.Request) {
	var eventBody DTO.EventDTO
	if err := json.NewDecoder(req.Body).Decode(&eventBody); err != nil {
		i.log.Println("Error decoding request body:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	if err := i.EventService.CreateEvent(eventBody); err != nil {
		i.log.Println("Error creating event:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	i.log.Println("Event created")
	res.WriteHeader(http.StatusCreated)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"message": "Event created"}`))
}

func (i *Event) GetEvents(res http.ResponseWriter, req *http.Request) {
	events, err := i.EventService.GetAllEvents()
	if err != nil {
		i.log.Println("Error getting events:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(events); err != nil {
		i.log.Println("Error encoding response:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (i *Event) GetEventById(res http.ResponseWriter, req *http.Request) {
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	eventDTO, err := i.EventService.GetEventByID(id)
	if err != nil {
		i.log.Println("Error getting event:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(eventDTO); err != nil {
		i.log.Println("Error encoding response:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (i *Event) UpdateEventById(res http.ResponseWriter, req *http.Request) {
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	var eventBody DTO.EventDTO
	if err := json.NewDecoder(req.Body).Decode(&eventBody); err != nil {
		i.log.Println("Error decoding request body:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	eventBody.ID = id
	if err := i.EventService.UpdateEvent(eventBody); err != nil {
		i.log.Println("Error updating event:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	i.log.Println("Event updated")
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(`{"message": "Event updated"}`))
}

func (i *Event) DeleteEventById(res http.ResponseWriter, req *http.Request) {
	id, err := getIdValue(req)
	if err != nil {
		i.log.Println("Error getting ID:", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if err := i.EventService.DeleteEvent(id); err != nil {
		i.log.Println("Error deleting event:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	i.log.Println("Event deleted")
	res.WriteHeader(http.StatusNoContent)
}
