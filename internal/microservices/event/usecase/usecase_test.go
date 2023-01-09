package usecase

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/microservices/event/mock"
	"Diploma/internal/models"
	"database/sql"
	// "testing"

	// "github.com/golang/mock/gomock"
	// "github.com/stretchr/testify/assert"
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

type getPeopleCountTest struct {
	name string
	inputPlaceID int
	beforeTest func(eventRepo *mock.MockRepository)
	outputPeopleCount int
	outputErr error
}

var getPeopleCountTests = []getPeopleCountTest {
	{
		"Successfully got count from existing event",
		1,
		func(eventRepo *mock.MockRepository) {
			eventRepo.EXPECT().
				GetPeopleCount(1).
				Return(10, nil)
		},
		10,
		nil,
	},
	{
		"Event is not existing so new event created with count 0",
		1,
		func(eventRepo *mock.MockRepository) {
			eventRepo.EXPECT().
				GetPeopleCount(1).
				Return(0, sql.ErrNoRows)
		
			eventRepo.EXPECT().
				CreateKudaGoEvent(1).
				Return(nil)
		},
		0,
		nil,
	},
	{
		"Something wrong on database People Count",
		1,
		func(eventRepo *mock.MockRepository) {
			eventRepo.EXPECT().
				GetPeopleCount(1).
				Return(0, customErrors.ErrPostgres)
		},
		0,
		customErrors.ErrPostgres,
	},
	{
		"Something wrong on database Create KudaGo Event",
		1,
		func(eventRepo *mock.MockRepository) {
			eventRepo.EXPECT().
				GetPeopleCount(1).
				Return(0, sql.ErrNoRows)
			
			eventRepo.EXPECT().
				CreateKudaGoEvent(1).
				Return(customErrors.ErrPostgres)
		},
		0,
		customErrors.ErrPostgres,
	},
}

// func TestGetEvent(t *testing.T) {
// 	placeRepositoryMock := new(mock.EventRepositoryMock)
// 	eventUsecaseTest := NewEventUsecase(placeRepositoryMock)
// 	for _, test := range getEventTests {
// 		placeRepositoryMock.On("GetEvent",test.eventId).Return(test.outputEvent, test.outputErr)
// 		actualEvent, actualErr := eventUsecaseTest.GetEvent(test.eventId)
// 		assert.Equal(t, test.outputEvent, actualEvent)
// 		assert.Nil(t, actualErr)
// 	}
// }

// func TestGetEvents(t *testing.T) {
// 	placeRepositoryMock := new(mock.EventRepositoryMock)
// 	eventUsecaseTest := NewEventUsecase(placeRepositoryMock)
// 	for _, test := range getEventsTests {
// 		placeRepositoryMock.On("GetEvents", test.page).Return(test.outputEvents, test.outputErr)
// 		actualEvents, actualErr := eventUsecaseTest.GetEvents(test.page)
// 		assert.Equal(t, test.outputEvents, actualEvents)
// 		assert.Nil(t, actualErr)
// 	}
// }

// func TestGetPeopleCount(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
	
// 	for _, test := range getPeopleCountTests {
// 		t.Run(test.name, func(t *testing.T) {
// 			EventRepositoryMock := mock.NewMockRepository(ctrl)
// 			UsecaseTest := NewEventUsecase(EventRepositoryMock)
// 			if test.beforeTest != nil {
// 				test.beforeTest(EventRepositoryMock)
// 			}
// 			actualPeopleCount, actualErr := UsecaseTest.GetPeopleCount(test.inputPlaceID)
// 			assert.Equal(t, test.outputPeopleCount, actualPeopleCount)
// 			assert.Equal(t, test.outputErr, actualErr)
// 		})
// 	}
// }