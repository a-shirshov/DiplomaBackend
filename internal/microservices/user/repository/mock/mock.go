package mock

import (
	"Diploma/internal/models"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetUser(userId int) (*models.User, error) {
	args := m.Called(userId)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserRepositoryMock) GetUserByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *UserRepositoryMock) UpdateUser(userId int, user *models.User) (*models.User, error) {
	args := m.Called(userId, user)
	return args.Get(0).(*models.User), args.Error(1)
}