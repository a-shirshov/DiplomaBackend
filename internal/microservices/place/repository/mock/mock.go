package mock

import (
	"Diploma/internal/models"

	"github.com/stretchr/testify/mock"
)

type PlaceRepositoryMock struct {
	mock.Mock
}

func(m *PlaceRepositoryMock) GetPlaces(page int) ([]*models.Place, error) {
	args := m.Called(page)
	return args.Get(0).([]*models.Place), args.Error(1)
}

func (m *PlaceRepositoryMock) GetPlace(id int) (*models.Place, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Place), args.Error(1)
}