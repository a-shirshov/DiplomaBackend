package usecase

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/microservices/auth/mock"
	"Diploma/internal/models"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type createUserTest struct {
	name string
	inputUser *models.User
	beforeTest func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository,
		passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager)
	ExpectedUser *models.User
	ExpectedTokenDetails *models.TokenDetails
	ExpectedErr error
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
		func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository, 
			passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager) {
			passwordHasher.EXPECT().
				GenerateHashFromPassword("password").
				Return("hashed_password", nil)

			authRepo.EXPECT().
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
			ImgUrl: "/images/uuid.webp",
		}, 
		&models.TokenDetails{},
		nil,
	},
	{
		"Error during hashing password",
		&models.User{
			Name: "Artyom",
			Surname: "Shirshov",
			Email:  "ash@mail.ru",
			Password: "password",
			ImgUrl: "uuid",
		},
		func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository, 
			passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager) {
			passwordHasher.EXPECT().
				GenerateHashFromPassword("password").
				Return("", customErrors.ErrHashingProblems)

		},
		&models.User{},
		&models.TokenDetails{},
		customErrors.ErrHashingProblems,
	},
	{
		"Error during hashing password",
		&models.User{
			Name: "Artyom",
			Surname: "Shirshov",
			Email:  "ash@mail.ru",
			Password: "password",
			ImgUrl: "uuid",
		},
		func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository, 
			passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager) {
			passwordHasher.EXPECT().
				GenerateHashFromPassword("password").
				Return("", customErrors.ErrHashingProblems)

		},
		&models.User{},
		&models.TokenDetails{},
		customErrors.ErrHashingProblems,
	},
	{
		"Error during saving user to db",
		&models.User{
			Name: "Artyom",
			Surname: "Shirshov",
			Email:  "ash@mail.ru",
			Password: "password",
			ImgUrl: "uuid",
		},
		func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository, 
			passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager) {
			passwordHasher.EXPECT().
				GenerateHashFromPassword("password").
				Return("hashed_password", nil)

				authRepo.EXPECT().
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
					&models.User{},
					customErrors.ErrPostgres,
				)

		},
		&models.User{},
		&models.TokenDetails{},
		customErrors.ErrPostgres,
	},
	{
		"User already exists",
		&models.User{
			Name: "Artyom",
			Surname: "Shirshov",
			Email:  "ash@mail.ru",
			Password: "password",
			ImgUrl: "uuid",
		},
		func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository, 
			passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager) {
			passwordHasher.EXPECT().
				GenerateHashFromPassword("password").
				Return("hashed_password", nil)

				authRepo.EXPECT().
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
					&models.User{},
					customErrors.ErrUserExists,
				)

		},
		&models.User{},
		&models.TokenDetails{},
		customErrors.ErrUserExists,
	},
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range createUserTests {
		t.Run(test.name, func(t *testing.T){
			mockAuthRepo := mock.NewMockRepository(ctrl)
			mockAuthSessionRepo := mock.NewMockSessionRepository(ctrl)
			mockPasswordHasher := mock.NewMockPasswordHasher(ctrl)
			mockTokenManager := mock.NewMockTokenManager(ctrl)
			testAuthUsecase := NewAuthUsecase(mockAuthRepo, mockAuthSessionRepo, mockPasswordHasher, mockTokenManager)

			if test.beforeTest != nil {
				test.beforeTest(mockAuthRepo, mockAuthSessionRepo,
					mockPasswordHasher, mockTokenManager)
			}

			actualUser, actualTokenDetails, actualErr := testAuthUsecase.CreateUser(test.inputUser)
			assert.Equal(t, test.ExpectedUser, actualUser)
			assert.Equal(t, test.ExpectedTokenDetails, actualTokenDetails)
			assert.Equal(t, test.ExpectedErr, actualErr)
		})

	}
}

type signInTest struct {
	name string
	inputUser *models.LoginUser
	beforeTest func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository,
		passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager)
	ExpectedUser *models.User
	ExpectedErr error
}

var signInTests = []signInTest{
	{
		"Successfully signIn",
		&models.LoginUser{
			Email:  "ash@mail.ru",
			Password: "requestPassword",
		},
		func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository, 
			passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager) {
				authRepo.EXPECT().
					GetUserByEmail("ash@mail.ru").
					Return(&models.User{
						ID: 1,
						Name: "Artyom",
						Surname: "Shirshov",
						Password: "dbPasswordHashed",
						ImgUrl: "uuid",
					},
					nil,
				)

				passwordHasher.EXPECT().
					VerifyPassword("requestPassword", "dbPasswordHashed").
					Return(nil)

				tokenManager.EXPECT().
					CreateToken(1).
					Return(&models.TokenDetails{}, nil)

				sessionRepo.EXPECT().
					SaveTokens(1,&models.TokenDetails{})
				
			},
		&models.User{
			ID: 1,
			Name: "Artyom",
			Surname: "Shirshov",
			ImgUrl: "/images/uuid.webp",
		},
		nil,
	},
}

func TestSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range signInTests {
		t.Run(test.name, func(t *testing.T){
			mockAuthRepo := mock.NewMockRepository(ctrl)
			mockAuthSessionRepo := mock.NewMockSessionRepository(ctrl)
			mockPasswordHasher := mock.NewMockPasswordHasher(ctrl)
			mockTokenManager := mock.NewMockTokenManager(ctrl)
			testAuthUsecase := NewAuthUsecase(mockAuthRepo, mockAuthSessionRepo, mockPasswordHasher, mockTokenManager)

			if test.beforeTest != nil {
				test.beforeTest(mockAuthRepo, mockAuthSessionRepo,
					mockPasswordHasher, mockTokenManager)
			}

			actualUser, _, actualErr := testAuthUsecase.SignIn(test.inputUser)
			assert.Equal(t, test.ExpectedUser, actualUser)
			assert.Equal(t, test.ExpectedErr, actualErr)
		})

	}
}

type logoutTest struct {
	name string
	inputAccessDetails *models.AccessDetails
	inputRefresh string
	beforeTest func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository,
		passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager)
	ExpectedErr error
}

var logoutTests = []logoutTest{
	{
		"Successfully logout",
		&models.AccessDetails{
			AccessUuid: "uuid",
		},
		"refresh_uuid",
		func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository, 
			passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager) {
				sessionRepo.EXPECT().
					DeleteAuth("uuid").
					Return(nil)
				
				sessionRepo.EXPECT().
					DeleteAuth("refresh_uuid").
					Return(nil)
			},
		nil,
	},
}

func TestLogout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range logoutTests {
		t.Run(test.name, func(t *testing.T){
			mockAuthRepo := mock.NewMockRepository(ctrl)
			mockAuthSessionRepo := mock.NewMockSessionRepository(ctrl)
			mockPasswordHasher := mock.NewMockPasswordHasher(ctrl)
			mockTokenManager := mock.NewMockTokenManager(ctrl)
			testAuthUsecase := NewAuthUsecase(mockAuthRepo, mockAuthSessionRepo, mockPasswordHasher, mockTokenManager)

			if test.beforeTest != nil {
				test.beforeTest(mockAuthRepo, mockAuthSessionRepo,
					mockPasswordHasher, mockTokenManager)
			}

			actualErr := testAuthUsecase.Logout(test.inputAccessDetails, test.inputRefresh)
			assert.Equal(t, test.ExpectedErr, actualErr)
		})

	}
}

type refreshTest struct {
	name string
	inputRefreshToken string
	beforeTest func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository,
		passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager)
	ExpectedTokens *models.Tokens
	ExpectedErr error
}

var refreshTests = []refreshTest{
	{
		"Successfully refresh",
		"refresh_token",
		func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository, 
			passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager) {
				claims := jwt.MapClaims{}
				claims["refresh_uuid"] = "refresh_uuid"
				claims["user_id"] = 1.0
				tokenManager.EXPECT().CheckTokenAndGetClaims("refresh_token").Return(claims, nil)
				sessionRepo.EXPECT().DeleteAuth("refresh_uuid").Return(nil)

				ts := &models.TokenDetails{
					AccessToken: "access_token",
					RefreshToken: "refresh_token",
				}
				tokenManager.EXPECT().CreateToken(1).Return(ts, nil)
				sessionRepo.EXPECT().SaveTokens(1, ts).Return(nil)
			},
			&models.Tokens{
				AccessToken: "access_token",
				RefreshToken: "refresh_token",
			},
		nil,
	},
}

func TestRefresh(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range refreshTests {
		t.Run(test.name, func(t *testing.T){
			mockAuthRepo := mock.NewMockRepository(ctrl)
			mockAuthSessionRepo := mock.NewMockSessionRepository(ctrl)
			mockPasswordHasher := mock.NewMockPasswordHasher(ctrl)
			mockTokenManager := mock.NewMockTokenManager(ctrl)
			testAuthUsecase := NewAuthUsecase(mockAuthRepo, mockAuthSessionRepo, mockPasswordHasher, mockTokenManager)

			if test.beforeTest != nil {
				test.beforeTest(mockAuthRepo, mockAuthSessionRepo,
					mockPasswordHasher, mockTokenManager)
			}

			actualTokens, actualErr := testAuthUsecase.Refresh(test.inputRefreshToken)
			assert.Equal(t, test.ExpectedTokens, actualTokens)
			assert.Equal(t, test.ExpectedErr, actualErr)
		})

	}
}

type findUserByEmailTest struct {
	name string
	inputEmail string
	beforeTest func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository,
		passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager)
	ExpectedUser *models.User
	ExpectedErr error
}

var findUserByEmailTests = []findUserByEmailTest{
	{
		"Successfully found user by email",
		"ash@mail.ru",
		func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository, 
			passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager) {
				authRepo.EXPECT().GetUserByEmail("ash@mail.ru").
					Return(
						&models.User{
						ID: 1,
						Name: "Artyom",
						Surname: "Shirshov",
						ImgUrl: "uuid",
					},
					nil)
			},
		&models.User{
			ID: 1,
			Name: "Artyom",
			Surname: "Shirshov",
			ImgUrl: "/images/uuid.webp",
		},
		nil,
	},
}

func TestFindUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range findUserByEmailTests {
		t.Run(test.name, func(t *testing.T){
			mockAuthRepo := mock.NewMockRepository(ctrl)
			mockAuthSessionRepo := mock.NewMockSessionRepository(ctrl)
			mockPasswordHasher := mock.NewMockPasswordHasher(ctrl)
			mockTokenManager := mock.NewMockTokenManager(ctrl)
			testAuthUsecase := NewAuthUsecase(mockAuthRepo, mockAuthSessionRepo, mockPasswordHasher, mockTokenManager)

			if test.beforeTest != nil {
				test.beforeTest(mockAuthRepo, mockAuthSessionRepo,
					mockPasswordHasher, mockTokenManager)
			}

			actualUser, actualErr := testAuthUsecase.FindUserByEmail(test.inputEmail)
			assert.Equal(t, test.ExpectedUser, actualUser)
			assert.Equal(t, test.ExpectedErr, actualErr)
		})

	}
}

type CreateAndSavePasswordRedeemCodeTest struct {
	name string
	inputEmail string
	beforeTest func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository,
		passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager)
	ExpectedErr error
}

var CreateAndSavePasswordRedeemCodeTests = []CreateAndSavePasswordRedeemCodeTest{
	{
		"Successfully created and saved redeem code",
		"ash@mail.ru",
		func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository, 
			passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager) {
				sessionRepo.EXPECT().SavePasswordRedeemCode("ash@mail.ru", gomock.Any()).Return(nil)
			},
		nil,
	},
}

func TestCreateAndSavePasswordRedeemCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range CreateAndSavePasswordRedeemCodeTests {
		t.Run(test.name, func(t *testing.T){
			mockAuthRepo := mock.NewMockRepository(ctrl)
			mockAuthSessionRepo := mock.NewMockSessionRepository(ctrl)
			mockPasswordHasher := mock.NewMockPasswordHasher(ctrl)
			mockTokenManager := mock.NewMockTokenManager(ctrl)
			testAuthUsecase := NewAuthUsecase(mockAuthRepo, mockAuthSessionRepo, mockPasswordHasher, mockTokenManager)

			if test.beforeTest != nil {
				test.beforeTest(mockAuthRepo, mockAuthSessionRepo,
					mockPasswordHasher, mockTokenManager)
			}

			actualRedeemCode, actualErr := testAuthUsecase.CreateAndSavePasswordRedeemCode(test.inputEmail)
			assert.True(t, (actualRedeemCode >= 100000) && (actualRedeemCode < 1000000), "actualRedeemCode is", actualRedeemCode)
			assert.Equal(t, test.ExpectedErr, actualErr)
		})

	}
}

type CheckRedeemCodeTest struct {
	name string
	inputRedeemCodeStruct *models.RedeemCodeStruct
	beforeTest func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository,
		passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager)
	ExpectedErr error
}

var CheckRedeemCodeTests = []CheckRedeemCodeTest{
	{
		"Successfully checked code",
		&models.RedeemCodeStruct{
			Email: "ash@mail.ru",
			RedeemCode: 123456,
		},
		func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository, 
			passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager) {
				sessionRepo.EXPECT().CheckRedeemCode("ash@mail.ru", 123456).Return(nil)
			},
		nil,
	},
}

func TestCheckRedeemCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range CheckRedeemCodeTests {
		t.Run(test.name, func(t *testing.T){
			mockAuthRepo := mock.NewMockRepository(ctrl)
			mockAuthSessionRepo := mock.NewMockSessionRepository(ctrl)
			mockPasswordHasher := mock.NewMockPasswordHasher(ctrl)
			mockTokenManager := mock.NewMockTokenManager(ctrl)
			testAuthUsecase := NewAuthUsecase(mockAuthRepo, mockAuthSessionRepo, mockPasswordHasher, mockTokenManager)

			if test.beforeTest != nil {
				test.beforeTest(mockAuthRepo, mockAuthSessionRepo,
					mockPasswordHasher, mockTokenManager)
			}

			actualErr := testAuthUsecase.CheckRedeemCode(test.inputRedeemCodeStruct)
			assert.Equal(t, test.ExpectedErr, actualErr)
		})

	}
}

type UpdatePasswordTest struct {
	name string
	inputRedeemCodeStruct *models.RedeemCodeStruct
	beforeTest func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository,
		passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager)
	ExpectedErr error
}

var UpdatePasswordTests = []UpdatePasswordTest{
	{
		"Successfully checked code",
		&models.RedeemCodeStruct{
			Email: "ash@mail.ru",
			Password: "password",
		},
		func(authRepo *mock.MockRepository, sessionRepo *mock.MockSessionRepository, 
			passwordHasher *mock.MockPasswordHasher, tokenManager *mock.MockTokenManager) {
				sessionRepo.EXPECT().CheckAccessToNewPassword("ash@mail.ru").Return(true)
				passwordHasher.EXPECT().GenerateHashFromPassword("password").Return("hashed_password", nil)
				authRepo.EXPECT().UpdatePassword("hashed_password", "ash@mail.ru").Return(nil)
			},
		nil,
	},
}

func TestUpdatePasswordCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range UpdatePasswordTests {
		t.Run(test.name, func(t *testing.T){
			mockAuthRepo := mock.NewMockRepository(ctrl)
			mockAuthSessionRepo := mock.NewMockSessionRepository(ctrl)
			mockPasswordHasher := mock.NewMockPasswordHasher(ctrl)
			mockTokenManager := mock.NewMockTokenManager(ctrl)
			testAuthUsecase := NewAuthUsecase(mockAuthRepo, mockAuthSessionRepo, mockPasswordHasher, mockTokenManager)

			if test.beforeTest != nil {
				test.beforeTest(mockAuthRepo, mockAuthSessionRepo,
					mockPasswordHasher, mockTokenManager)
			}

			actualErr := testAuthUsecase.UpdatePassword(test.inputRedeemCodeStruct)
			assert.Equal(t, test.ExpectedErr, actualErr)
		})

	}
}