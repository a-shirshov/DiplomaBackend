package user

import (
	"Diploma/internal/models"
)

type Repository interface {
	GetUser(int) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	UpdateUserImage(userID int, imgUUID string) (*models.User, error)
	ChangePassword(userID int, password string) (error)
}
