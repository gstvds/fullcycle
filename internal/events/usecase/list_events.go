package usecase

import "github.com/gstvds/fullcycle/internal/events/domain"

type ListEventsOutputDTO struct {
	Events []EventDTO `json:"events"`
}

type ListEventsUseCase struct {
	repository domain.EventRepository
}

func NewListEventsUseCase(repository domain.EventRepository) *ListEventsUseCase {
	return &ListEventsUseCase{repository: repository}
}

func (listEventsUseCase *ListEventsUseCase) Execute() (*ListEventsOutputDTO, error) {
	events, err := listEventsUseCase.repository.ListEvents()

	if err != nil {
		return nil, err
	}

	eventDTO := make([]EventDTO, len(events))

	for i, event := range events {
		eventDTO[i] = EventDTO{
			ID:           event.ID,
			Name:         event.Name,
			Location:     event.Location,
			Organization: event.Organization,
			Rating:       string(event.Rating),
			Date:         event.Date.Format("2006-01-02 15:04:05"),
			ImageURL:     event.ImageURL,
			Capacity:     event.Capacity,
			Price:        event.Price,
			PartnerID:    event.PartnerID,
		}
	}

	return &ListEventsOutputDTO{Events: eventDTO}, nil
}
