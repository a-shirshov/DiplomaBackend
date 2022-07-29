package user

import (
	"Diploma/internal/models"
	"Diploma/utils"
)

type Usecase interface {
	GetUser(int) (*models.User, error)
	UpdateUser(*utils.AccessDetails, *models.User) (*models.User, error)
}
