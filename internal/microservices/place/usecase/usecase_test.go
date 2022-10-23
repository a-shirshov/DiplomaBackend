package usecase

import (
	"Diploma/internal/microservices/place/repository/mock"
	"Diploma/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

type getPlaceTest struct {
	placeId int
	outputPlace *models.Place
	outputError error
}

var getPlaceTests = []getPlaceTest{
	{
		1, &models.Place{
			ID: 1,
			Name: "Name_1",
			Description: "Description_1",
			About: "About",
			Category: "Category",
			ImgUrl: "ImgUrl_1",
		}, nil,
	},
}

func TestGetPlace(t *testing.T){
	placeRepositoryMock := new(mock.PlaceRepositoryMock)
	placeUsecaseTest := NewPlaceUsecase(placeRepositoryMock)
	for _, test := range getPlaceTests {
		placeRepositoryMock.On("GetPlace", test.placeId).Return(test.outputPlace, nil)
		actualPlace, actualErr := placeUsecaseTest.GetPlace(test.placeId)
		assert.Equal(t,test.outputPlace, actualPlace)
		assert.Nil(t, actualErr)
	}
}