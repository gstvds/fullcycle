package domain

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
