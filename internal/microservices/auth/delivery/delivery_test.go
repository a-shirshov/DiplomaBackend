package delivery

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/microservices/auth/mock"
	"Diploma/internal/middleware/middleware"
	"Diploma/internal/models"
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type signInTest struct {
	name string
	inputStructToBeJSON *models.LoginUser
	beforeTest func(authUsecase *mock.MockUsecase)
	expectedStatusCode int
	expectedStructToBeJSON interface{}
}

var signInTests = []signInTest {
	{
		"Succesfully signIn",
		&models.LoginUser{
			Email: "ash@mail.ru",
			Password: "password",
		},
		func(authUsecase *mock.MockUsecase) {
			authUsecase.EXPECT().
				SignIn(&models.LoginUser{
					Email: "ash@mail.ru",
					Password: "password",
				}).
				Return(&models.User{
					ID: 1,
					Name: "Artyom",
					Surname: "Shirshov",
					Email: "ash@mail.ru",
					DateOfBirth: "2001-08-06",
					City: "msk",
					About: "about",
					ImgUrl: "user_face.png",
				}, 
					&models.TokenDetails{
						AccessToken: "jwt.valid",
						RefreshToken: "jwt.valid",
					},
					nil,
				)
		},
		http.StatusOK,
		&models.UserWithTokens{
			User: models.User{
				ID: 1,
				Name: "Artyom",
				Surname: "Shirshov",
				Email: "ash@mail.ru",
				DateOfBirth: "2001-08-06",
				City: "msk",
				About: "about",
				ImgUrl: "user_face.png",
			},
			Tokens: models.Tokens{
				AccessToken: "jwt.valid",
				RefreshToken: "jwt.valid",
			},
		},
	},
	{
		"No user with this email",
		&models.LoginUser{
			Email: "ash@mail.ru",
			Password: "password",
		},
		func(authUsecase *mock.MockUsecase) {
			authUsecase.EXPECT().
				SignIn(&models.LoginUser{
					Email: "ash@mail.ru",
					Password: "password",
				}).
				Return(&models.User{}, 
					&models.TokenDetails{},
					customErrors.ErrWrongEmail,
				)
		},
		http.StatusForbidden,
		&models.Message{
			Message: customErrors.ErrWrongEmail.Error(),
		},
	},
	{
		"Wrong Password",
		&models.LoginUser{
			Email: "ash@mail.ru",
			Password: "password",
		},
		func(authUsecase *mock.MockUsecase) {
			authUsecase.EXPECT().
				SignIn(&models.LoginUser{
					Email: "ash@mail.ru",
					Password: "password",
				}).
				Return(&models.User{}, 
					&models.TokenDetails{},
					customErrors.ErrWrongPassword,
				)
		},
		http.StatusForbidden,
		&models.Message{
			Message: customErrors.ErrWrongPassword.Error(),
		},
	},
	{
		"Internal server error",
		&models.LoginUser{
			Email: "ash@mail.ru",
			Password: "password",
		},
		func(authUsecase *mock.MockUsecase) {
			authUsecase.EXPECT().
				SignIn(&models.LoginUser{
					Email: "ash@mail.ru",
					Password: "password",
				}).
				Return(&models.User{}, 
					&models.TokenDetails{},
					customErrors.ErrPostgres,
				)
		},
		http.StatusInternalServerError,
		&models.Message{
			Message: customErrors.ErrPostgres.Error(),
		},
	},
}

func TestSignIn(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range signInTests{
		t.Run(test.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			router := gin.Default()
			mockauthUsecase := mock.NewMockUsecase(ctrl)
			deliveryTest := NewAuthDelivery(mockauthUsecase)
			mws := middleware.NewMiddleware(mock.NewMockSessionRepository(ctrl), mock.NewMockTokenManager(ctrl))

			if test.beforeTest != nil {
				test.beforeTest(mockauthUsecase)
			}

			router.Use(mws.MiddlewareValidateLoginUser())
			responseRecorder := httptest.NewRecorder()
			router.POST("/login", deliveryTest.SignIn)

			inputJSON, _ := json.Marshal(test.inputStructToBeJSON)
			request, err := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(inputJSON))
			assert.NoError(t, err)
			router.ServeHTTP(responseRecorder, request)

			assert.Equal(t, test.expectedStatusCode, responseRecorder.Code)
			expectedJSON, _ := json.Marshal(test.expectedStructToBeJSON)
			assert.Equal(t, string(expectedJSON), responseRecorder.Body.String())
		})
	}
}

type logoutTest struct {
	name string
	inputAccessDetails *models.AccessDetails
	beforeTest func(authUsecase *mock.MockUsecase)
	expectedStatusCode int
}

var logoutTests = []logoutTest {
	{
		"Successfully logout",
		&models.AccessDetails{
			AccessUuid: "access_uuid",
			UserId: 1,
		},
		func(authUsecase *mock.MockUsecase) {
			authUsecase.EXPECT().
				Logout(&models.AccessDetails{
					AccessUuid: "access_uuid",
					UserId: 1,
				}).
				Return(nil)
		},
		http.StatusOK,
	},
	{
		"No authorization token",
		&models.AccessDetails{},
		nil,
		http.StatusUnauthorized,
	},
	{
		"Error during Logout",
		&models.AccessDetails{
			AccessUuid: "access_uuid",
			UserId: 1,
		},
		func(authUsecase *mock.MockUsecase) {
			authUsecase.EXPECT().
				Logout(&models.AccessDetails{
					AccessUuid: "access_uuid",
					UserId: 1,
				}).
				Return(errors.New("error"))
		},
		http.StatusUnauthorized,
	},
}

func TestLogout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range logoutTests{
		t.Run(test.name, func(t *testing.T) {
			mockauthUsecase := mock.NewMockUsecase(ctrl)
			deliveryTest := NewAuthDelivery(mockauthUsecase)

			if test.beforeTest != nil {
				test.beforeTest(mockauthUsecase)
			}

			responseRecorder := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			ctx, router := gin.CreateTestContext(responseRecorder)
			router.Use(func(ctx *gin.Context){
				ctx.Set("access_details", *test.inputAccessDetails)
			})
			
			router.POST("/logout", deliveryTest.Logout)
			ctx.Request, _ = http.NewRequest(http.MethodPost, "/logout", nil)
			router.ServeHTTP(responseRecorder, ctx.Request)

			assert.Equal(t, test.expectedStatusCode, responseRecorder.Code)
		})
	}
}