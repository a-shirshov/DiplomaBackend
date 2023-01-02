package event

import (
	"Diploma/internal/models"
)

type Usecase interface {
	GetEvents(page int) ([]*models.Event, error)
	GetEvent(eventId int) (*models.Event, error)
	GetEventsByToday(page int) ([]*models.Event, error)
}