package usecase

import (
	"Diploma/internal/microservices/event"
	"Diploma/internal/models"
	"database/sql"
)

type EventUsecase struct {
	eventRepo event.Repository
}

func NewEventUsecase(eventR event.Repository) (*EventUsecase) {
	return &EventUsecase{
		eventRepo: eventR,
	}
}

func (eU *EventUsecase) GetEvents(page int) ([]*models.Event, error) {
	return eU.eventRepo.GetEvents(page)
}

func (eU *EventUsecase) GetEvent(eventId int) (*models.Event, error) {
	return eU.eventRepo.GetEvent(eventId)
}

func (eU *EventUsecase) GetPeopleCountAndCheckMeeting(userID int, eventID int) (int, bool, error) {
	var peopleCount int
	isGoing := false
	peopleCount, err := eU.eventRepo.GetPeopleCount(eventID)
	if err != nil {
		if err != sql.ErrNoRows {
			return 0, isGoing, err
		}

		err = eU.eventRepo.CreateKudaGoEvent(eventID)
		if err != nil {
			return 0, isGoing, err
		}

		peopleCount = 0
		return peopleCount, isGoing, nil
	}
	if (userID != 0) {
		isGoing, err = eU.eventRepo.CheckKudaGoMeeting(userID, eventID)
		if err != nil {
			return 0, isGoing, err
		}
	}
	return peopleCount, isGoing, nil
}

func (eU *EventUsecase) SwitchEventMeeting(userID int, eventID int) (error) {
	return eU.eventRepo.SwitchEventMeeting(userID, eventID)
}