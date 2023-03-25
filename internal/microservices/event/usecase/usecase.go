package usecase

import (
	"Diploma/internal/microservices/event"
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

func (eU *EventUsecase) GetPeopleCountWithEventCreatedIfNecessary(eventID int) (int, error) {
	var peopleCount int
	peopleCount, err := eU.eventRepo.GetPeopleCount(eventID)
	if err != nil {
		if err != sql.ErrNoRows {
			return 0, err
		}

		err = eU.eventRepo.CreateKudaGoEvent(eventID)
		if err != nil {
			return 0, err
		}

		peopleCount = 0
		return peopleCount, nil
	}

	return peopleCount, nil
}

func (eU *EventUsecase) CheckKudaGoMeeting(userID, eventID int) (bool, error) {
	isGoing, err := eU.eventRepo.CheckKudaGoMeeting(userID, eventID)
	if err != nil {
		return false, err
	}
	
	return isGoing, nil
}

func (eU *EventUsecase) CheckKudaGoFavourite(userID, eventID int) (bool, error) {
	isFavourite, err := eU.eventRepo.CheckKudaGoFavourite(userID, eventID)
	if err != nil {
		return false, err
	}
	
	return isFavourite, nil
}

func (eU *EventUsecase) SwitchEventMeeting(userID int, eventID int) (error) {
	return eU.eventRepo.SwitchEventMeeting(userID, eventID)
}

func (eU *EventUsecase) SwitchEventFavourite(userID int, eventID int) (error) {
	return eU.eventRepo.SwitchEventFavourite(userID, eventID)
}

func (eU *EventUsecase) GetFavouriteKudagoEventsIDs(userID int) ([]int, error) {
	return eU.eventRepo.GetFavouriteKudagoEventsIDs(userID)
}