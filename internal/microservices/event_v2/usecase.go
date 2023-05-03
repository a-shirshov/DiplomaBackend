package eventV2

import (
	"Diploma/internal/models"
)

type Usecase interface {
	GetExternalEvents(userID int, city string, page int) (*[]models.MyEvent, error)
	GetTodayEvents(userID int, city string, page int) (*[]models.MyEvent, error)
	GetCloseEvents(lat string, lon string, userID int, page int) (*[]models.MyEvent, error)
	GetExternalEvent(userID int, eventID int) (*models.MyFullEvent, error)
	GetNLPVector(description string) ([]float64, error)
	GetRandomEvents(userID int) (*[]models.MyEvent, error)
	GetVector(eventID int) (*[]float64, error)
	GetVectorTitle(eventID int) (*[]float64, error)
	SwitchLikeEvent(userID, eventID int) (*models.MyEvent, error)
	GetFavourites(userID, checkedUserID, page int) (*[]models.MyEvent, error)
	SearchEvents(userID int, searchingEvent string, page int) (*[]models.MyEvent, error)
}
