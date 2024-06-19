package domain

import "time"

type Rating string

const (
	RatingFree Rating = "L"
	Rating10   Rating = "L10"
	Rating12   Rating = "L12"
	Rating14   Rating = "L14"
	Rating16   Rating = "L16"
	Rating18   Rating = "L18"
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

// This will be moved in the future
type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}

type TicketType string

const (
	TicketTypeHalf TicketType = "half" // Half-price Ticket
	TicketTypeFull TicketType = "full" // Full-price Ticket
)

// This will be moved in the future
type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}

type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartnerID    string
	Spots        []Spot
	Tickets      []Ticket
}
