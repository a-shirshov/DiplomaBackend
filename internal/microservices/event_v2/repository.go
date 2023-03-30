package eventV2

import "Diploma/internal/models"

type Repository interface {
	GetExternalEvents(userID int) (*models.MyEvents, error)
}
