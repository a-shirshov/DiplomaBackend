package user

import (
	"Diploma/internal/models"
)

type Repository interface {
	GetUser(int) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	GetFavouriteKudagoEventsIDs(userID int) ([]int, error)
}
