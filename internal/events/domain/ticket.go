package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidTicketPrice = errors.New("ticket price must be greater than 0")
	ErrInvalidTicketType  = errors.New("invalid ticket type")
)

type TicketType string

const (
	TicketTypeHalf TicketType = "half" // Half-price Ticket
	TicketTypeFull TicketType = "full" // Full-price Ticket
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}

func IsValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

func (ticket *Ticket) CalculatePrice() {
	if ticket.TicketType == TicketTypeHalf {
		ticket.Price /= 2
	}
}

func (ticket Ticket) Validate() error {
	if ticket.Price <= 0 {
		return ErrInvalidTicketPrice
	}

	return nil
}

func NewTicket(event *Event, spot *Spot, ticketType TicketType) (*Ticket, error) {
	if !IsValidTicketType(ticketType) {
		return nil, ErrInvalidTicketType
	}

	ticket := &Ticket{
		ID:         uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
		TicketType: ticketType,
		Price:      event.Price,
	}
	ticket.CalculatePrice()

	if err := ticket.Validate(); err != nil {
		return nil, err
	}

	return ticket, nil
}
