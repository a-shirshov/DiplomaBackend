package event

import (
	"Diploma/internal/models"
)

type Usecase interface {
	GetEvents(placeId, page int) ([]*models.Event, error)
	GetEvent(eventId int) (*models.Event, error)
}