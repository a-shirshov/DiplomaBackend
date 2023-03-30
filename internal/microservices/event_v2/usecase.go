package eventV2

import (
	"Diploma/internal/models"
)

type Usecase interface {
	GetExternalEvents(userID int) (*models.MyEvents, error)
}
