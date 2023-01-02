package user

import (
	"Diploma/internal/models"
)

type Usecase interface {
	GetUser(int) (*models.User, error)
	UpdateUser(*models.User) (*models.User, error)
}
