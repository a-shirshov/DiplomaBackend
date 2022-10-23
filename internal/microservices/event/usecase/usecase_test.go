package usecase

import (
	"Diploma/internal/microservices/event/repository/mock"
	"Diploma/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

type getEventTest struct {
	eventId int
	outputEvent *models.Event
	outputErr error
}

type getEventsTest struct {
	page int
	placeId int
	outputEvents []*models.Event
	outputErr error
}

var getEventTests = []getEventTest{
	{
		1, &models.Event{
			ID: 1,
			Name: "Name_1",
			Description: "Description_1",
			About: "About_1",
			Category: "Category_1",
			Tags: []string{"Tag_1"," Tag_2"},
			SpecialInfo: "SpecialInfo_1",
		}, nil,
	},
}

var getEventsTests = []getEventsTest{
	{
		1, 1, []*models.Event{
			{
				ID: 1,
				Name: "Name_1",
				Description: "Description_1",
				About: "About_1",
				Category: "Category_1",
				Tags: []string{"Tag_1"," Tag_2"},
				SpecialInfo: "SpecialInfo_1",
			},
			{
				ID: 2,
				Name: "Name_2",
				Description: "Description_2",
				About: "About_2",
				Category: "Category_2",
				Tags: []string{"Tag_1"," Tag_2"},
				SpecialInfo: "SpecialInfo_2",
			},
		}, nil,
	},
}

func TestGetEvent(t *testing.T){
	placeRepositoryMock := new(mock.EventRepositoryMock)
	eventUsecaseTest := NewEventUsecase(placeRepositoryMock)
	for _, test := range getEventTests {
		placeRepositoryMock.On("GetEvent",test.eventId).Return(test.outputEvent, test.outputErr)
		actualEvent, actualErr := eventUsecaseTest.GetEvent(test.eventId)
		assert.Equal(t, test.outputEvent, actualEvent)
		assert.Nil(t, actualErr)
	}
}

func TestGetEvents(t *testing.T) {
	placeRepositoryMock := new(mock.EventRepositoryMock)
	eventUsecaseTest := NewEventUsecase(placeRepositoryMock)
	for _, test := range getEventsTests {
		placeRepositoryMock.On("GetEvents", test.page, test.placeId).Return(test.outputEvents, test.outputErr)
		actualEvents, actualErr := eventUsecaseTest.GetEvents(test.page, test.placeId)
		assert.Equal(t, test.outputEvents, actualEvents)
		assert.Nil(t, actualErr)
	}
}