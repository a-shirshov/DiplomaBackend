package user

import (
	"Diploma/internal/models"
)

type Usecase interface {
	GetUser(int) (*models.User, error)
	UpdateUser(*models.User) (*models.User, error)
	UpdateUserImage(userID int, imgUUID string) (*models.User, error)
	ChangePassword(userID int, password string) (error)
}
