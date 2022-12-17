package usecase

import (
	"Diploma/internal/microservices/auth/repository/mock"
	"Diploma/internal/models"
	"Diploma/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mocklib "github.com/stretchr/testify/mock"
)

type createUserTest struct {
	inputUser *models.User
	outputUser *models.User
	outputErr error
}

type logoutTest struct {
	au *utils.AccessDetails
	outputErr error
}

type signInTest struct {
	inputUser *models.LoginUser
	outputUser *models.User
	outputErr error
}

var createUserTests = []createUserTest{
	{
		&models.User{
			Name: "User_1",
			Surname: "Surname_1",
			Email:  "Email_1",
			Password: "Password_1",
		},
		&models.User{
			Name: "User_1",
			Surname: "Surname_1",
			Email:  "Email_1",
		}, nil,
	},
}

var logoutTests = []logoutTest{
	{
		&utils.AccessDetails{}, nil,
	},
}

var signInTests = []signInTest{
	{
		&models.LoginUser{
			Email:  "Email_1",
			Password: "Password_1",
		},
		&models.User{
			ID: 1,
			Name: "User_1",
			Surname: "Surname_1",
			Email:  "Email_1",
			About: "About_1",
			ImgUrl: "ImgUrl_1",
		},
		nil,
	},
}

// func TestCreateUser(t *testing.T){
// 	authRepositoryMock := new(mock.AuthRepositoryMock)
// 	authSessionRepositoryMock := new(mock.AuthSessionRepositoryMock)
// 	authUsecaseTest := NewAuthUsecase(authRepositoryMock, authSessionRepositoryMock)
// 	for _, test := range createUserTests{
// 		authRepositoryMock.On("CreateUser", test.inputUser).Return(test.outputUser, test.outputErr)
// 		actualUser, actualErr := authUsecaseTest.CreateUser(test.inputUser)
// 		assert.Equal(t, test.outputUser, actualUser)
// 		assert.Nil(t, actualErr)
// 	}
// }

func TestSignIn(t *testing.T){
	authRepositoryMock := new(mock.AuthRepositoryMock)
	authSessionRepositoryMock := new(mock.AuthSessionRepositoryMock)
	authUsecaseTest := NewAuthUsecase(authRepositoryMock, authSessionRepositoryMock)
	for _, test := range signInTests {
		hash, hashErr := utils.GenerateHashFromPassword(test.inputUser.Password)
		require.Nil(t,hashErr)
		test.outputUser.Password = hash

		authRepositoryMock.On("GetUserByEmail", test.inputUser.Email).Return(test.outputUser, nil)

		authSessionRepositoryMock.On("SaveTokens", test.outputUser.ID, mocklib.Anything).Return(test.outputErr)
		actualUser, _, actualErr := authUsecaseTest.SignIn(test.inputUser)
		assert.Equal(t, test.outputUser, actualUser)
		assert.Nil(t, actualErr)
	}
}

func TestLogout(t *testing.T){
	authRepositoryMock := new(mock.AuthRepositoryMock)
	authSessionRepositoryMock := new(mock.AuthSessionRepositoryMock)
	authUsecaseTest := NewAuthUsecase(authRepositoryMock, authSessionRepositoryMock)
	for _, test := range logoutTests{
		authSessionRepositoryMock.On("DeleteAuth", test.au.AccessUuid).Return(test.outputErr)
		actualErr := authUsecaseTest.Logout(test.au)
		assert.Nil(t, actualErr)
	}
}

func TestRefresh(t *testing.T){

}