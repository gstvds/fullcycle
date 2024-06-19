package domain

type EventRepository interface {
	ListEvents() ([]Event, error)
	GetEventByID(eventID string) (*Event, error)
	GetSpotsByEventID(eventID string) ([]*Spot, error)
	GetSpotByName(eventID, spotName string) (*Spot, error)
	// Creation will be added in the future
	// CreateEvent(event *Event) error
	// CreateSpot(spot *Spot) error
	CreateTicket(ticket *Ticket) error
	ReserveSpot(spotID, ticketID string) error
}
