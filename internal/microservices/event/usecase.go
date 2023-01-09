package event

import (
	"Diploma/internal/models"
)

type Usecase interface {
	GetEvents(page int) ([]*models.Event, error)
	GetEvent(eventId int) (*models.Event, error)
	GetPeopleCountAndCheckMeeting(userID int, eventID int) (int, bool, error)
	SwitchEventMeeting(userID int, eventID int) (error)
}