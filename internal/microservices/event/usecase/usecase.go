package usecase

import (
	"Diploma/internal/microservices/event"
	"Diploma/internal/models"
)

type EventUsecase struct {
	eventRepo event.Repository
}

func NewEventUsecase(eventR event.Repository) (*EventUsecase) {
	return &EventUsecase{
		eventRepo: eventR,
	}
}

func (eU *EventUsecase) GetEvents(page int) ([]*models.Event, error) {
	return eU.eventRepo.GetEvents(page)
}

func (eU *EventUsecase) GetEvent(eventId int) (*models.Event, error) {
	return eU.eventRepo.GetEvent(eventId)
}

func (eU *EventUsecase) GetEventsByToday(page int) ([]*models.Event, error) {
	return eU.eventRepo.GetEvents(page)
}