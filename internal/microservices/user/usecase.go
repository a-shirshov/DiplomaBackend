package user

import (
	"Diploma/internal/models"
)

type Usecase interface {
	GetUser(int) (*models.User, error)
	UpdateUser(int, *models.User) (*models.User, error)
}
