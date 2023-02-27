package usecase

import (
	"Diploma/internal/microservices/user/mock"
	"Diploma/internal/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type getUserTest struct {
	name 			string
	inputUserID 	int
	beforeTest 		func(userRepo *mock.MockRepository)
	ExpectedUser 	*models.User
	ExpectedErr 	error
}

var getUserTests = []getUserTest{
	{
		"Successfully get user with image",
		1,
		func(userRepo *mock.MockRepository) {
			userRepo.EXPECT().
				GetUser(1).
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
		},
		&models.User{
			ID: 1,
			Name: "Artyom",
			Surname: "Shirshov",
			Email:  "ash@mail.ru",
			ImgUrl: "/images/uuid.webp",
		}, 
		nil,
	},
	{
		"Successfully get user with no image",
		1,
		func(userRepo *mock.MockRepository) {
			userRepo.EXPECT().
				GetUser(1).
				Return(
					&models.User{
						ID: 1,
						Name: "Artyom",
						Surname: "Shirshov",
						Email:  "ash@mail.ru",
					},
					nil,
				)
		},
		&models.User{
			ID: 1,
			Name: "Artyom",
			Surname: "Shirshov",
			Email:  "ash@mail.ru",
			ImgUrl: "",
		}, 
		nil,
	},
}

type updateUserTest struct {
	name 			string
	inputUser		*models.User
	beforeTest 		func(userRepo *mock.MockRepository)
	ExpectedUser 	*models.User
	ExpectedErr 	error
}

var updateUserTests = []updateUserTest{
	{
		"Successfully update user (returned with image)",
		&models.User{
			ID: 1,
			Name: "Artyom",
			Surname: "Shirshov",
			Email:  "ash@mail.ru",
		},
		func(userRepo *mock.MockRepository) {
			userRepo.EXPECT().
				UpdateUser(&models.User{
					ID: 1,
					Name: "Artyom",
					Surname: "Shirshov",
					Email:  "ash@mail.ru",
				}).
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
		},
		&models.User{
			ID: 1,
			Name: "Artyom",
			Surname: "Shirshov",
			Email:  "ash@mail.ru",
			ImgUrl: "/images/uuid.webp",
		}, 
		nil,
	},
	{
		"Successfully update user (returned with no image)",
		&models.User{
			ID: 1,
			Name: "Artyom",
			Surname: "Shirshov",
			Email:  "ash@mail.ru",
		},
		func(userRepo *mock.MockRepository) {
			userRepo.EXPECT().
				UpdateUser(&models.User{
					ID: 1,
					Name: "Artyom",
					Surname: "Shirshov",
					Email:  "ash@mail.ru",
				}).
				Return(
					&models.User{
						ID: 1,
						Name: "Artyom",
						Surname: "Shirshov",
						Email:  "ash@mail.ru",
					},
					nil,
				)
		},
		&models.User{
			ID: 1,
			Name: "Artyom",
			Surname: "Shirshov",
			Email:  "ash@mail.ru",
			ImgUrl: "",
		}, 
		nil,
	},
}

type updateUserImageTest struct {
	name 			string
	inputUserID		int
	inputImgUUID	string
	beforeTest 		func(userRepo *mock.MockRepository)
	ExpectedUser 	*models.User
	ExpectedErr 	error
}

var updateUserImageTests = []updateUserImageTest{
	{
		"Successfully update user (returned with image)",
		1,
		"ImgUUID",
		func(userRepo *mock.MockRepository) {
			userRepo.EXPECT().
				UpdateUserImage(1, "ImgUUID").
				Return(
					&models.User{
						ID: 1,
						Name: "Artyom",
						Surname: "Shirshov",
						Email:  "ash@mail.ru",
						ImgUrl: "ImgUUID",
					},
					nil,
				)
		},
		&models.User{
			ID: 1,
			Name: "Artyom",
			Surname: "Shirshov",
			Email:  "ash@mail.ru",
			ImgUrl: "/images/ImgUUID.webp",
		}, 
		nil,
	},
}

// type GetFavouriteKudagoEventsIDsTest struct {
// 	name 				string
// 	inputUserID			int
// 	beforeTest 			func(userRepo *mock.MockRepository)
// 	ExpectedEventIDs 	[]int
// 	ExpectedErr 		error
// }

// var GetFavouriteKudagoEventsIDsTests = []GetFavouriteKudagoEventsIDsTest{
// 	{
// 		"Successfully return favourite event ids",
// 		1,
// 		func(userRepo *mock.MockRepository) {
// 			userRepo.EXPECT().
// 				GetFavouriteKudagoEventsIDs(1).
// 				Return(
// 					[]int{1, 2, 3},
// 					nil,
// 				)
// 		},
// 		[]int{1, 2, 3},
// 		nil,
// 	},
// }

func TestGetUserUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range getUserTests {
		t.Run(test.name, func(t *testing.T){
			mockUserRepo := mock.NewMockRepository(ctrl)
			testUserRepository := NewUserUsecase(mockUserRepo)

			if test.beforeTest != nil {
				test.beforeTest(mockUserRepo)
			}

			actualUser, actualErr := testUserRepository.GetUser(test.inputUserID)
			assert.Equal(t, test.ExpectedUser, actualUser)
			assert.Equal(t, test.ExpectedErr, actualErr)
		})
	}
}

func TestUpdateUserUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range updateUserTests {
		t.Run(test.name, func(t *testing.T){
			mockUserRepo := mock.NewMockRepository(ctrl)
			testUserRepository := NewUserUsecase(mockUserRepo)

			if test.beforeTest != nil {
				test.beforeTest(mockUserRepo)
			}

			actualUser, actualErr := testUserRepository.UpdateUser(test.inputUser)
			assert.Equal(t, test.ExpectedUser, actualUser)
			assert.Equal(t, test.ExpectedErr, actualErr)
		})
	}
}

func TestUpdateUserImageUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range updateUserImageTests {
		t.Run(test.name, func(t *testing.T){
			mockUserRepo := mock.NewMockRepository(ctrl)
			testUserRepository := NewUserUsecase(mockUserRepo)

			if test.beforeTest != nil {
				test.beforeTest(mockUserRepo)
			}

			actualUser, actualErr := testUserRepository.UpdateUserImage(test.inputUserID, test.inputImgUUID)
			assert.Equal(t, test.ExpectedUser, actualUser)
			assert.Equal(t, test.ExpectedErr, actualErr)
		})
	}
}

// func TestGetFavouriteKudagoEventsIDs(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	for _, test := range GetFavouriteKudagoEventsIDsTests {
// 		t.Run(test.name, func(t *testing.T){
// 			mockUserRepo := mock.NewMockRepository(ctrl)
// 			testUserRepository := NewUserUsecase(mockUserRepo)

// 			if test.beforeTest != nil {
// 				test.beforeTest(mockUserRepo)
// 			}

// 			actualEventIDs, actualErr := testUserRepository.GetFavouriteKudagoEventsIDs(test.inputUserID)
// 			assert.Equal(t, test.ExpectedEventIDs, actualEventIDs)
// 			assert.Equal(t, test.ExpectedErr, actualErr)
// 		})
// 	}
// }
