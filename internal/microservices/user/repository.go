package user

import (
	"Diploma/internal/models"
)

type Repository interface {
	GetUser(int) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(userId int, user *models.User) (*models.User, error)
}
