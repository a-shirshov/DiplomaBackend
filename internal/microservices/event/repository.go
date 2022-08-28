package event

import "Diploma/internal/models"

type Repository interface {
	GetEvents(placeId, page int) (*[]*models.Event, error)
	GetEvent(eventId int) (*models.Event, error)
}