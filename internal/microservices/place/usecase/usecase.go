package usecase

import (
	"Diploma/internal/microservices/place"
	"Diploma/internal/models"
)

type PlaceUsecase struct {
	placeRepo place.Repository
}

func NewPlaceUsecase(placeR place.Repository) *PlaceUsecase {
	return &PlaceUsecase{
		placeRepo: placeR,
	}
}

func (pU *PlaceUsecase) GetPlaces(page int) (*[]*models.Place, error) {
	return pU.placeRepo.GetPlaces(page)
}

func (pU *PlaceUsecase) GetPlace(id int) (*models.Place, error) {
	return pU.placeRepo.GetPlace(id)
}