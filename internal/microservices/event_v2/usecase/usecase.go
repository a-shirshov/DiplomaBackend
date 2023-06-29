package usecase

import (
	"Diploma/internal/microservices/event_v2"
	"Diploma/internal/models"
	"context"
	"log"
	"time"

	pb "Diploma/protos"

	"google.golang.org/grpc"
)

type eventUsecaseV2 struct {
	eventRepositoryV2 eventV2.Repository
	grpcConn grpc.ClientConnInterface
}

func NewEventUsecaseV2 (eventV2R eventV2.Repository, grpcConn grpc.ClientConnInterface) (*eventUsecaseV2){
	return &eventUsecaseV2{
		eventRepositoryV2: eventV2R,
		grpcConn: grpcConn,
	}
}

func(eU *eventUsecaseV2) GetExternalEvents(userID int, city string, page int) (*[]models.MyEvent, error) {
	if (city == "") {
		return eU.eventRepositoryV2.GetExternalEvents(userID, page)
	}
	return eU.eventRepositoryV2.GetExternalEventsWithCity(userID, city, page)
}

func(eU *eventUsecaseV2) GetTodayEvents(userID int, city string, page int) (*[]models.MyEvent, error) {
	startTime := time.Now().Unix()
	endTime := time.Now().Add(24 * time.Hour).Unix()
	if (city == "") {
		return eU.eventRepositoryV2.GetTodayEvents(startTime, endTime, userID, page)
	}
	return eU.eventRepositoryV2.GetTodayEventsWithCity(startTime, endTime, city, userID, page)
}

func(eU *eventUsecaseV2) GetCloseEvents(lat string, lon string, userID int, page int) (*[]models.MyEvent, error) {
	return eU.eventRepositoryV2.GetCloseEvents(lat, lon, userID, page) 
}

func(eU *eventUsecaseV2) GetExternalEvent(userID int, eventID int) (*models.MyFullEvent, error) {
	return eU.eventRepositoryV2.GetExternalEvent(userID, eventID)
}

func(eU *eventUsecaseV2) GetNLPVector(description string) ([]float64, error) {
	c := pb.NewMyProtoClient(eU.grpcConn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()
	r, err := c.ReturnVector(ctx, &pb.StringRequest{Description: description})
	if err != nil {
		log.Println("Error grpc:", err)
		return nil, err
	}
	return r.GetVector(), nil
}

func (eU *eventUsecaseV2) GetRandomEvents(userID int) (*[]models.MyEvent, error) {
	return eU.eventRepositoryV2.GetRandomEvents(userID)
}

func (eU *eventUsecaseV2) GetVector(eventID int) (*[]float64, error) {
	return eU.eventRepositoryV2.GetVector(eventID)
}

func (eU *eventUsecaseV2) GetVectorTitle(eventID int) (*[]float64, error) {
	return eU.eventRepositoryV2.GetVectorTitle(eventID)
}

func (eU *eventUsecaseV2) SwitchLikeEvent(userID int, eventID int) (*models.MyEvent, error) {
	err := eU.eventRepositoryV2.SwitchLikeEvent(userID, eventID)
	if err != nil {
		return nil, err
	}

	return eU.eventRepositoryV2.GetEvent(userID, eventID)
}

func (eU *eventUsecaseV2) GetFavourites(userID int, checkedUserID, page int) (*[]models.MyEvent, error) {
	return eU.eventRepositoryV2.GetFavourites(userID, checkedUserID, page)
}

func (eU *eventUsecaseV2) SearchEvents(userID int, searchingEvent string, page int) (*[]models.MyEvent, error) {
	return eU.eventRepositoryV2.SearchEvents(userID, searchingEvent, page)
}