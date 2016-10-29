package usecases

import "github.com/cmouli84/xgameodds/domain"

// EventsInteractor struct
type EventsInteractor struct {
	EventsRepository EventsRepository
}

// EventsRepository interface
type EventsRepository interface {
	GetEventsByDate(eventDate string) []domain.Event
}

// GetEventsByDate function
func (interactor *EventsInteractor) GetEventsByDate(eventDate string) []domain.Event {
	return interactor.EventsRepository.GetEventsByDate(eventDate)
}
