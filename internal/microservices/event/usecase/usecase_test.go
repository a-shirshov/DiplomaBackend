package usecase

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/microservices/event/mock"
	"database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type getPeopleCountWithEventCreatedIfNecessaryTest struct {
	name string
	inputEventID int
	beforeTest func(eventRepo *mock.MockRepository)
	outputPeopleCount int
	outputErr error
}

var getPeopleCountWithEventCreatedIfNecessaryTests = []getPeopleCountWithEventCreatedIfNecessaryTest {
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

type checkKudaGoMeetingTest struct {
	name string
	inputUserID int
	inputEventID int
	beforeTest func(eventRepo *mock.MockRepository)
	expectedIsGoing bool
	outputErr error
}

var checkKudaGoMeetingTests = []checkKudaGoMeetingTest {
	{
		"Found meeting",
		1,
		1,
		func(eventRepo *mock.MockRepository) {
			eventRepo.EXPECT().
				CheckKudaGoMeeting(1, 1).
				Return(true, nil)
		},
		true,
		nil,
	},
}

type checkKudaGoFavouriteTest struct {
	name string
	inputUserID int
	inputEventID int
	beforeTest func(eventRepo *mock.MockRepository)
	expectedIsFavourite bool
	outputErr error
}

var checkKudaGoFavouriteTests = []checkKudaGoFavouriteTest {
	{
		"Found favourite",
		1,
		1,
		func(eventRepo *mock.MockRepository) {
			eventRepo.EXPECT().
				CheckKudaGoFavourite(1, 1).
				Return(true, nil)
		},
		true,
		nil,
	},
}

type switchKudaGoMeetingTest struct {
	name string
	inputUserID int
	inputEventID int
	beforeTest func(eventRepo *mock.MockRepository)
	outputErr error
}

var switchKudaGoMeetingTests = []switchKudaGoMeetingTest {
	{
		"Switched favourite",
		1,
		1,
		func(eventRepo *mock.MockRepository) {
			eventRepo.EXPECT().
				SwitchEventMeeting(1, 1).
				Return(nil)
		},
		nil,
	},
}

type switchKudaGoFavouriteTest struct {
	name string
	inputUserID int
	inputEventID int
	beforeTest func(eventRepo *mock.MockRepository)
	outputErr error
}

var switchKudaGoFavouriteTests = []switchKudaGoFavouriteTest {
	{
		"Switched favourite",
		1,
		1,
		func(eventRepo *mock.MockRepository) {
			eventRepo.EXPECT().
				SwitchEventFavourite(1, 1).
				Return(nil)
		},
		nil,
	},
}

func TestGetPeopleCountWithEventCreatedIfNecessary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	for _, test := range getPeopleCountWithEventCreatedIfNecessaryTests {
		t.Run(test.name, func(t *testing.T) {
			EventRepositoryMock := mock.NewMockRepository(ctrl)
			UsecaseTest := NewEventUsecase(EventRepositoryMock)
			if test.beforeTest != nil {
				test.beforeTest(EventRepositoryMock)
			}
			actualPeopleCount, actualErr := UsecaseTest.GetPeopleCountWithEventCreatedIfNecessary(test.inputEventID)
			assert.Equal(t, test.outputPeopleCount, actualPeopleCount)
			assert.Equal(t, test.outputErr, actualErr)
		})
	}
}

func TestCheckKudaGoMeeting(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	for _, test := range checkKudaGoMeetingTests {
		t.Run(test.name, func(t *testing.T) {
			EventRepositoryMock := mock.NewMockRepository(ctrl)
			UsecaseTest := NewEventUsecase(EventRepositoryMock)
			if test.beforeTest != nil {
				test.beforeTest(EventRepositoryMock)
			}
			actualIsGoing, actualErr := UsecaseTest.CheckKudaGoMeeting(test.inputUserID, test.inputEventID)
			assert.Equal(t, test.expectedIsGoing, actualIsGoing)
			assert.Equal(t, test.outputErr, actualErr)
		})
	}
}

func TestCheckKudaGoFavourite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	for _, test := range checkKudaGoFavouriteTests {
		t.Run(test.name, func(t *testing.T) {
			EventRepositoryMock := mock.NewMockRepository(ctrl)
			UsecaseTest := NewEventUsecase(EventRepositoryMock)
			if test.beforeTest != nil {
				test.beforeTest(EventRepositoryMock)
			}
			actualIsGoing, actualErr := UsecaseTest.CheckKudaGoFavourite(test.inputUserID, test.inputEventID)
			assert.Equal(t, test.expectedIsFavourite, actualIsGoing)
			assert.Equal(t, test.outputErr, actualErr)
		})
	}
}

func TestSwitchKudaGoFavourite(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	for _, test := range switchKudaGoFavouriteTests {
		t.Run(test.name, func(t *testing.T) {
			EventRepositoryMock := mock.NewMockRepository(ctrl)
			UsecaseTest := NewEventUsecase(EventRepositoryMock)
			if test.beforeTest != nil {
				test.beforeTest(EventRepositoryMock)
			}
			actualErr := UsecaseTest.SwitchEventFavourite(test.inputUserID, test.inputEventID)
			assert.Equal(t, test.outputErr, actualErr)
		})
	}
}

func TestSwitchKudaGoMeeting(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	for _, test := range switchKudaGoMeetingTests {
		t.Run(test.name, func(t *testing.T) {
			EventRepositoryMock := mock.NewMockRepository(ctrl)
			UsecaseTest := NewEventUsecase(EventRepositoryMock)
			if test.beforeTest != nil {
				test.beforeTest(EventRepositoryMock)
			}
			actualErr := UsecaseTest.SwitchEventMeeting(test.inputUserID, test.inputEventID)
			assert.Equal(t, test.outputErr, actualErr)
		})
	}
}
