package usecase

import (
	"Diploma/internal/microservices/auth/mock"
	"Diploma/internal/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type createUserTest struct {
	name string
	inputUser *models.User
	beforeTest func(userRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository,
		passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager)
	ExpectedUser *models.User
	ExpectedErr error
}

type logoutTest struct {
	au *models.AccessDetails
	outputErr error
}

type signInTest struct {
	inputUser *models.LoginUser
	outputUser *models.User
	outputErr error
}

var createUserTests = []createUserTest{
	{
		"Successfully create user",
		&models.User{
			Name: "Artyom",
			Surname: "Shirshov",
			Email:  "ash@mail.ru",
			Password: "password",
			ImgUrl: "uuid",
		},
		func(userRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository, 
			passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager) {
			passwordHasher.EXPECT().
				GenerateHashFromPassword("password").
				Return("hashed_password", nil)

			userRepo.EXPECT().
				CreateUser(
					&models.User{
						Name: "Artyom",
						Surname: "Shirshov",
						Email:  "ash@mail.ru",
						Password: "hashed_password",
						ImgUrl: "uuid",
					},
				).
				Return(
					&models.User{
						ID: 1,
						Name: "Artyom",
						Surname: "Shirshov",
						Email:  "ash@mail.ru",
						ImgUrl: "uuid",
					},
					nil,
				)

			tokenManager.EXPECT().
				CreateToken(1).
				Return(&models.TokenDetails{}, nil)

			sessionRepo.EXPECT().
				SaveTokens(1, &models.TokenDetails{}).
				Return(nil)	
		},
		&models.User{
			ID: 1,
			Name: "Artyom",
			Surname: "Shirshov",
			Email:  "ash@mail.ru",
			ImgUrl: "/images/uuid",
		}, 
		nil,
	},
}

var logoutTests = []logoutTest{
	{
		&models.AccessDetails{}, nil,
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

func TestCreateUserUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range createUserTests {
		t.Run(test.name, func(t *testing.T){
			mockAuthRepo := mock.NewMockRepository(ctrl)
			mockAuthSessionRepo := mock.NewMockSessionRepository(ctrl)
			mockPasswordHasher := mock.NewMockPasswordHasher(ctrl)
			mockTokenManager := mock.NewMockTokenManager(ctrl)
			testAuthRepository := NewAuthUsecase(mockAuthRepo, mockAuthSessionRepo, mockPasswordHasher, mockTokenManager)

			if test.beforeTest != nil {
				test.beforeTest(mockAuthRepo, mockAuthSessionRepo,
					mockPasswordHasher, mockTokenManager)
			}

			actualUser, _, actualErr := testAuthRepository.CreateUser(test.inputUser)
			assert.Equal(t, test.ExpectedUser, actualUser)
			assert.Equal(t, test.ExpectedErr, actualErr)
		})

	}
}