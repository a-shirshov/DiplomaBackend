package usecase

import (
	"Diploma/internal/microservices/event"
	"Diploma/internal/models"
	"database/sql"
	"sync"
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

func (eU *EventUsecase) GetPeopleCountAndCheckMeeting(userID int, eventID int) (int, bool, bool, error) {
	var peopleCount int
	isGoing := false
	isFavourite := false
	peopleCount, err := eU.eventRepo.GetPeopleCount(eventID)
	if err != nil {
		if err != sql.ErrNoRows {
			return 0, isGoing, isFavourite, err
		}

		err = eU.eventRepo.CreateKudaGoEvent(eventID)
		if err != nil {
			return 0, isGoing, isFavourite, err
		}

		peopleCount = 0
		return peopleCount, isGoing, isFavourite, nil
	}
	if (userID != 0) {
		var wg sync.WaitGroup
		wg.Add(2)
		go func(){
			defer wg.Done()
			isGoing, err = eU.eventRepo.CheckKudaGoMeeting(userID, eventID)
			if err != nil {
				return 
			}
		}()
		go func(){
			defer wg.Done()
			isFavourite, err = eU.eventRepo.CheckKudaGoFavourite(userID, eventID)
			if err != nil {
				return 
			}
		}()
		wg.Wait()
	}
	return peopleCount, isGoing, isFavourite, nil
}

func (eU *EventUsecase) SwitchEventMeeting(userID int, eventID int) (error) {
	return eU.eventRepo.SwitchEventMeeting(userID, eventID)
}

func (eU *EventUsecase) SwitchEventFavourite(userID int, eventID int) (error) {
	return eU.eventRepo.SwitchEventFavourite(userID, eventID)
}