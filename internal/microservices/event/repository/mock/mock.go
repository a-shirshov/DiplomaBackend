package mock

import (
	"Diploma/internal/models"

	"github.com/stretchr/testify/mock"
)

type EventRepositoryMock struct {
	mock.Mock
}

func(m *EventRepositoryMock) GetEvents(page int) ([]*models.Event, error) {
	args := m.Called(page)
	return args.Get(0).([]*models.Event), args.Error(1)
}

func (m *EventRepositoryMock) GetEvent(eventId int) (*models.Event, error) {
	args := m.Called(eventId)
	return args.Get(0).(*models.Event), args.Error(1)
}