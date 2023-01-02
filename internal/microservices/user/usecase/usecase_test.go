package usecase

import (
	"Diploma/internal/microservices/user/repository/mock"
	"Diploma/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

type getUserTest struct {
	userId int
	outputUser *models.User
	outputErr error
}

type updateUserTest struct {
	userId int
	inputUser *models.User
	outputUser *models.User
	outputErr error
}

var getUserTests = []getUserTest{
	{
		1, &models.User{}, nil,
	},
}

var updateUserTests = []updateUserTest{
	{
		1, &models.User{
			Name: "Name_1",
			Surname: "Surname_1",
			About: "About_1",
		}, &models.User{
			ID: 1,
			Name: "Name_1",
			Surname: "Surname_1",
			Email: "Email_1",
			About: "About_1",
			ImgUrl: "ImgUrl_1",
		}, nil,
	},
}

func TestGetUserByID(t *testing.T) {
	userRepositoryMock := new(mock.UserRepositoryMock)
	userUsecaseTest := NewUserUsecase(userRepositoryMock)
	for _, test := range getUserTests {
		userRepositoryMock.On("GetUser", test.userId).Return(test.outputUser, nil)
		actualUser, actualErr := userUsecaseTest.GetUser(test.userId)
		assert.Equal(t, test.outputUser, actualUser)
		assert.Nil(t, actualErr)
	}
}

func TestUpdateUser(t *testing.T) {
	userRepositoryMock := new(mock.UserRepositoryMock)
	userUsecaseTest := NewUserUsecase(userRepositoryMock)
	for _, test := range updateUserTests{
		userRepositoryMock.On("UpdateUser", test.inputUser).Return(test.outputUser, nil)
		actualUser, actualErr := userUsecaseTest.UpdateUser(test.inputUser)
		assert.Equal(t, test.outputUser, actualUser)
		assert.Nil(t, actualErr)
	}	
}