package place

import (
	"Diploma/internal/models"
)

type Usecase interface {
	GetPlaces(page int) (*[]*models.Place, error)
	GetPlace(id int) (*models.Place, error)
}