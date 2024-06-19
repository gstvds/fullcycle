package usecase

import "github.com/gstvds/fullcycle/internal/events/domain"

type BuyTicketsInputDTO struct {
	EventID    string   `json:"event_id"`
	Spots      []string `json:"spots"`
	TicketType string   `json:"ticket_type"`
	CardHash   string   `json:"card_hash"`
	Email      string   `json:"email"`
}

type BuyTicketsOutputDTO struct {
	Tickets []TicketDTO `json:"tickets"`
}

type BuyTicketsUseCase struct {
	repository     domain.EventRepository
	partnerFactory service.PartnerFactory
}

func NewBuyTicketsUseCase(repository domain.EventRepository, partnerFactory service.PartnerFactory) *BuyTicketsUseCase {
	return &BuyTicketsUseCase{repository: repository, partnerFactory: partnerFactory}
}

func (buyTicketsUseCase *BuyTicketsUseCase) Execute(input BuyTicketsInputDTO) (*BuyTicketsOutputDTO, error) {
	event, err := buyTicketsUseCase.repository.GetEventByID(input.EventID)

	if err != nil {
		return nil, err
	}

	partnerRequest := &service.ReservationRequest{
		EventID:    input.EventID,
		Spots:      input.Spots,
		TicketType: input.TicketType,
		CardHash:   input.CardHash,
		Email:      input.Email,
	}

	partnerService, err := buyTicketsUseCase.partnerFactory.CreatePartner(event.PartnerID)

	if err != nil {
		return nil, err
	}

	reservationResponse, err := partnerService.MakeReservation(partnerRequest)
	if err != nil {
		return nil, err
	}

	tickets := make([]domain.Ticket, len(reservationResponse))

	for i, reservation := range reservationResponse {
		spot, err := buyTicketsUseCase.repository.GetSpotByName(event.ID, reservation.Spot)
		if err != nil {
			return nil, err
		}

		ticket, err := domain.NewTicket(event, spot, domain.TicketType(reservation.TicketType))
		if err != nil {
			return nil, err
		}

		err = buyTicketsUseCase.repository.CreateTicket(ticket)
		if err != nil {
			return nil, err
		}

		spot.Reserve(ticket.ID)
		err = buyTicketsUseCase.repository.ReserveSpot(spot.ID, ticket.ID)
		if err != nil {
			return nil, err
		}
		tickets[i] = *ticket
	}

	ticketDTOs := make([]TicketDTO, len(tickets))
	for i, ticket := range tickets {
		ticketDTOs[i] = TicketDTO{
			ID:         ticket.ID,
			SpotID:     ticket.Spot.ID,
			TicketType: string(ticket.TicketType),
			Price:      ticket.Price,
		}
	}

	return &BuyTicketsOutputDTO{Tickets: ticketDTOs}, nil
}
