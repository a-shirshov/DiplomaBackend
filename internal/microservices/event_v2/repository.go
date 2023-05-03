package eventV2

import "Diploma/internal/models"

type Repository interface {
	GetExternalEvents(userID int, page int) (*[]models.MyEvent, error)
	GetTodayEvents(startTime int64, endTime int64, userID int, page int) (*[]models.MyEvent, error)
	GetCloseEvents(lat string, lon string, userID int, page int) (*[]models.MyEvent, error)
	GetExternalEvent(userID int, eventID int) (*models.MyFullEvent, error)
	GetRandomEvents(userID int) (*[]models.MyEvent, error)
	GetVector(eventID int) (*[]float64, error)
	GetVectorTitle(eventID int) (*[]float64, error)
	GetEvent(userID int, eventID int) (*models.MyEvent, error)
	SwitchLikeEvent(userID, eventID int) (error)
	GetFavourites(userID, checkedUserID, page int) (*[]models.MyEvent, error)
	SearchEvents(userID int, searchingEvent string, page int) (*[]models.MyEvent, error)
	GetExternalEventsWithCity(userID int, city string, page int) (*[]models.MyEvent, error)
	GetTodayEventsWithCity(startTime int64, endTime int64, city string, userID int, page int) (*[]models.MyEvent, error)
}
