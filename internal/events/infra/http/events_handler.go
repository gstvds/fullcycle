package http

import (
	"encoding/json"
	"net/http"

	"github.com/gstvds/fullcycle/internal/events/usecase"
)

type EventsHandler struct {
	listEventsUseCase *usecase.ListEventsUseCase
	listSpostUseCase  *usecase.ListSpotsUseCase
	getEventUseCase   *usecase.GetEventUseCase
	buyTicketsUseCase *usecase.BuyTicketsUseCase
}

func NewEventsHandler(
	listEventsUseCase *usecase.ListEventsUseCase,
	listSpostUseCase *usecase.ListSpotsUseCase,
	getEventUseCase *usecase.GetEventUseCase,
	buyTicketsUseCase *usecase.BuyTicketsUseCase,
) *EventsHandler {
	return &EventsHandler{
		listEventsUseCase: listEventsUseCase,
		listSpostUseCase:  listSpostUseCase,
		getEventUseCase:   getEventUseCase,
		buyTicketsUseCase: buyTicketsUseCase,
	}
}

func (handler *EventsHandler) ListEvents(writer http.ResponseWriter, request *http.Request) {
	output, err := handler.listEventsUseCase.Execute()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(output)
}

func (handler *EventsHandler) GetEvent(writer http.ResponseWriter, request *http.Request) {
	eventID := request.PathValue("eventID")
	input := usecase.GetEventInputDTO{ID: eventID}

	output, err := handler.getEventUseCase.Execute(input)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(output)
}

func (handler *EventsHandler) ListSpots(writer http.ResponseWriter, request *http.Request) {
	eventID := request.PathValue("eventID")
	input := usecase.ListSpotsInputDTO{EventID: eventID}

	output, err := handler.listSpostUseCase.Execute(input)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(output)
}

func (handler *EventsHandler) BuyTickets(writer http.ResponseWriter, request *http.Request) {
	var input usecase.BuyTicketsInputDTO
	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := handler.buyTicketsUseCase.Execute(input)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(output)
}
