package usecase

import (
	"Diploma/internal/microservices/event_v2"
	"Diploma/internal/models"
)

type eventUsecaseV2 struct {
	eventRepositoryV2 eventV2.Repository
}

func NewEventUsecaseV2 (eventV2R eventV2.Repository) (*eventUsecaseV2){
	return &eventUsecaseV2{
		eventRepositoryV2: eventV2R,
	}
}

func(eU *eventUsecaseV2) GetExternalEvents(userID int) (*models.MyEvents, error) {
	return eU.eventRepositoryV2.GetExternalEvents(userID)
}