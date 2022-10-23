package mock

import (
	"Diploma/utils"

	"github.com/stretchr/testify/mock"
)

type AuthSessionRepositoryMock struct {
	mock.Mock
}

func(m *AuthSessionRepositoryMock) SaveTokens(userId int, td *utils.TokenDetails) error {
	args := m.Called(userId, td)
	return args.Error(0)
}

func(m *AuthSessionRepositoryMock) FetchAuth(accessUuid string) (int, error) {
	args := m.Called(accessUuid)
	return args.Int(0), args.Error(1)
}

func(m *AuthSessionRepositoryMock) DeleteAuth(accessUuid string) error {
	args := m.Called(accessUuid)
	return args.Error(0)
}