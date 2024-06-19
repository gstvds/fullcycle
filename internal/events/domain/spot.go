package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidSpotNumber       = errors.New("invalid spot number")
	ErrSpotNotFound            = errors.New("spot not found")
	ErrSpotAlreadyReserved     = errors.New("spot already reserved")
	ErrSpotMustStartWithLetter = errors.New("spot name must start with a letter")
	ErrSpotMustEndWithNumber   = errors.New("spot name must end with a number")
	ErrSpotNameTooShort        = errors.New("spot name is too short")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}

func (spot Spot) Validate() error {
	if spot.Name == "" {
		return ErrInvalidSpotNumber
	}

	if len(spot.Name) < 2 {
		return ErrSpotNameTooShort
	}

	if spot.Name[0] < 'A' || spot.Name[0] > 'Z' {
		return ErrSpotMustStartWithLetter
	}
	if spot.Name[len(spot.Name)-1] < '0' || spot.Name[len(spot.Name)-1] > '9' {
		return ErrSpotMustEndWithNumber
	}

	return nil
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		ID:      uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvailable,
	}

	if err := spot.Validate(); err != nil {
		return nil, err
	}

	return spot, nil
}
