package event

import (
	"Diploma/internal/models"
)

type Usecase interface {
	GetEvents(page int) ([]*models.Event, error)
	GetEvent(id int) (*models.Event, error)
}