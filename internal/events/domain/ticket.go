package domain

import "errors"

var (
	ErrInvalidTicketPrice = errors.New("ticket price must be greater than 0")
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
