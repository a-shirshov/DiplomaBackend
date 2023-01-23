package event

import "Diploma/internal/models"

type Repository interface {
	GetEvents(page int) ([]*models.Event, error)
	GetEvent(eventId int) (*models.Event, error)
	GetPeopleCount(placeID int) (int, error)
	CreateKudaGoEvent(placeID int) (error)
	SwitchEventMeeting(userID int, eventID int) (error)
	CheckKudaGoMeeting(userID, eventID int) (bool, error)
	SwitchEventFavourite(userID int, eventID int) error
	CheckKudaGoFavourite(userID int, eventID int) (bool, error)
}