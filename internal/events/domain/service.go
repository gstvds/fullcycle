package domain

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidQuantity = errors.New("quantity must be greater than 0")
)

type spotService struct{}

func NewSpotService() *spotService {
	return &spotService{}
}

func (spot *spotService) GenerateSpots(event *Event, quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	for i := 0; i < quantity; i++ {
		spotName := fmt.Sprintf("%c%d", 'A'+i/10, i%10+1)
		spot, err := NewSpot(event, spotName)
		if err != nil {
			return err
		}

		event.Spots = append(event.Spots, *spot)
	}

	return nil
}
