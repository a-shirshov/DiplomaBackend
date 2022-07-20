package event

import "Diploma/internal/models"

type Repository interface {
	GetEvents(page int) ([]*models.Event, error)
}