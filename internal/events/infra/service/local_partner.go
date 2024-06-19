package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type LocalPartner struct {
	BaseURL string
}

type LocalPartnerReservationRequest struct {
	Spots      []string `json:"spots"`
	TicketKind string   `json:"ticket_kind"`
	Email      string   `json:"email"`
}

type LocalPartnerReservationResponse struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Spot       string `json:"spot"`
	TicketKind string `json:"ticket_kind"`
	Status     string `json:"status"`
	EventID    string `json:"event_id"`
}

func (partner *LocalPartner) MakeReservation(request *ReservationRequest) ([]ReservationResponse, error) {
	partnerRequest := LocalPartnerReservationRequest{
		Spots:      request.Spots,
		TicketKind: request.TicketType,
		Email:      request.Email,
	}

	body, err := json.Marshal(partnerRequest)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/events/%s/reserve", partner.BaseURL, request.EventID)

	bufferedBody := bytes.NewBuffer(body)
	httpRequest, err := http.NewRequest(http.MethodPost, url, bufferedBody)
	if err != nil {
		return nil, err
	}

	httpRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("reservation failed with status code: %d", httpResponse.StatusCode)
	}

	var partnerResponse []LocalPartnerReservationResponse
	if err := json.NewDecoder(httpResponse.Body).Decode(&partnerResponse); err != nil {
		return nil, err
	}

	response := make([]ReservationResponse, len(partnerResponse))
	for i, reservation := range partnerResponse {
		response[i] = ReservationResponse{
			ID:         reservation.ID,
			Email:      reservation.Email,
			Spot:       reservation.Spot,
			TicketType: reservation.TicketKind,
			Status:     reservation.Status,
			EventID:    reservation.EventID,
		}
	}

	return response, nil
}
