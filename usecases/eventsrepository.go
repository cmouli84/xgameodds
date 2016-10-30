package usecases

import "github.com/cmouli84/xgameodds/domain"

// EventsInteractor struct
type EventsInteractor struct {
	EventsRepository EventsRepository
}

// EventsRepository interface
type EventsRepository interface {
	GetNflEventsByDate(eventDate string) []domain.Event
}

// GetNflEventsByDate function
func (interactor *EventsInteractor) GetNflEventsByDate(eventDate string) []domain.Event {
	return interactor.EventsRepository.GetNflEventsByDate(eventDate)
}
