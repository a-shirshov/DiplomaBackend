package mock

import (
	"Diploma/internal/models"

	"github.com/stretchr/testify/mock"
)

type AuthRepositoryMock struct {
	mock.Mock
}

func(m *AuthRepositoryMock) CreateUser(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func(m *AuthRepositoryMock) GetUserByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}