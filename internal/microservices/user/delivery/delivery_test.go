package delivery

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/microservices/user/mock"
	middlewareMock "Diploma/internal/middleware/mock"
	"Diploma/internal/models"
	"Diploma/internal/router"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func prepareTestEnvironment() (*gin.Context, *gin.Engine, *gin.RouterGroup, *httptest.ResponseRecorder) {
	responseRecorder := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, testEngine := gin.CreateTestContext(responseRecorder)
	routerAPI := testEngine.Group("/api")
	userRouter := routerAPI.Group("/users")
	return ctx, testEngine, userRouter, responseRecorder
}

type getUserTest struct {
	name string
	queryParamUserID string
	beforeTest func(userUsecase *mock.MockUsecase)
	expectedUserJSON string
	expectedStatusCode int
}

var getUserTests = []getUserTest {
	{
		"Successfully get user",
		"1",
		func(userUsecase *mock.MockUsecase) {
			userUsecase.EXPECT().
				GetUser(1).
				Return(
					&models.User{
						ID: 1,
						Name: "Artyom",
						Surname: "Shirshov",
						About: "About",
						ImgUrl: "ImgUrl",
					},
					nil)
		},
		`{
			"id":1, 
			"name":"Artyom", 
			"surname":"Shirshov", 
			"about":"About", 
			"imgUrl":"ImgUrl", 
			"dateOfBirth":"", 
			"city":""
		}`,
		http.StatusOK,
	},
	{
		"User was not found",
		"1",
		func(userUsecase *mock.MockUsecase) {
			userUsecase.EXPECT().
				GetUser(1).
				Return(&models.User{}, customErrors.ErrUserNotFound)
		},
		`{
			"message":"user not found"
		}`,
		http.StatusNotFound,
	},
	{
		"Database problems",
		"1",
		func(userUsecase *mock.MockUsecase) {
			userUsecase.EXPECT().
				GetUser(1).
				Return(&models.User{}, customErrors.ErrPostgres)
		},
		`{
			"message":"database problems"
		}`,
		http.StatusInternalServerError,
	},
	{
		"Not integer id in request as param user_id",
		"first",
		nil,
		`{
			"message":"bad request"
		}`,
		http.StatusBadRequest,
	},
}

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range getUserTests{
		t.Run(test.name, func(t *testing.T) {
			mockUserUsecase := mock.NewMockUsecase(ctrl)
			deliveryTest := NewUserDelivery(mockUserUsecase)
			mockMiddlewares := middlewareMock.NewMockMiddleware(ctrl)

			if test.beforeTest != nil {
				test.beforeTest(mockUserUsecase)
			}

			mockMiddlewares.EXPECT().TokenAuthMiddleware().Return(func (ctx *gin.Context){ctx.Next()})
			mockMiddlewares.EXPECT().MiddlewareValidateUser().Return(func (ctx *gin.Context){ctx.Next()})

			ctx, testEngine, userRouter, responseRecorder := prepareTestEnvironment()
			router.UserEndpoints(userRouter, mockMiddlewares, deliveryTest)
			var err error
			ctx.Request, err = http.NewRequest(http.MethodGet, fmt.Sprintf("/api/users/%s", test.queryParamUserID), nil)
			assert.Nil(t, err)
			testEngine.ServeHTTP(responseRecorder, ctx.Request)

			assert.Equal(t, test.expectedStatusCode, responseRecorder.Code)
			assert.JSONEq(t, test.expectedUserJSON, responseRecorder.Body.String())
		})
	}
}

type updateUserTest struct {
	name string
	queryParamUserID string
	inputBodyJSON string
	inputAccessDetails models.AccessDetails
	beforeTest func(userUsecase *mock.MockUsecase)
	expectedUserJSON string
	expectedStatusCode int
}

var updateUserTests = []updateUserTest {
	{
		"Successfully update user",
		"1",
		`{
			"id":1, 
			"name":"Artyom", 
			"surname":"Shirshov", 
			"about":"About", 
			"imgUrl":"ImgUrl", 
			"dateOfBirth":"2001-08-06", 
			"city":"msk"
		}`,
		models.AccessDetails{
			AccessUuid: "uuid",
			UserId: 1,
		},
		func(userUsecase *mock.MockUsecase) {
			userUsecase.EXPECT().
				UpdateUser(
					&models.User{
						ID: 1,
						Name: "Artyom",
						Surname: "Shirshov",
						About: "About",
						ImgUrl: "ImgUrl",
						DateOfBirth: "2001-08-06",
						City: "msk",
					}).
				Return(
					&models.User{
						ID: 1,
						Name: "Artyom",
						Surname: "Shirshov",
						About: "About",
						ImgUrl: "ImgUrl",
						DateOfBirth: "2001-08-06",
						City: "msk",
					},
					nil)
		},
		`{
			"id":1, 
			"name":"Artyom", 
			"surname":"Shirshov", 
			"about":"About", 
			"imgUrl":"ImgUrl", 
			"dateOfBirth":"2001-08-06", 
			"city":"msk"
		}`,
		http.StatusOK,
	},
	{
		"No authorization token",
		"1",
		`{
			"id":1, 
			"name":"Artyom", 
			"surname":"Shirshov", 
			"about":"About", 
			"imgUrl":"ImgUrl", 
			"dateOfBirth":"2001-08-06", 
			"city":"msk"
		}`,
		models.AccessDetails{},
		nil,
		`{
			"message": "no authorization token"
		}`,
		http.StatusUnauthorized,
	},
	{
		"Wrong request json",
		"1",
		`{
			"user":1, 
		}`,
		models.AccessDetails{
			AccessUuid: "uuid",
			UserId: 1,
		},
		nil,
		`{
			"message": "input json is not correct"
		}`,
		http.StatusUnprocessableEntity,
	},
	{
		"Wrong request json",
		"1",
		`{
			"user":1, 
		}`,
		models.AccessDetails{
			AccessUuid: "uuid",
			UserId: 1,
		},
		nil,
		`{
			"message": "input json is not correct"
		}`,
		http.StatusUnprocessableEntity,
	},
	{
		"Diffrent query param and token id",
		"1",
		`{
			"id":1, 
			"name":"Artyom", 
			"surname":"Shirshov", 
			"about":"About", 
			"imgUrl":"ImgUrl", 
			"dateOfBirth":"2001-08-06", 
			"city":"msk"
		}`,
		models.AccessDetails{
			AccessUuid: "uuid",
			UserId: 2,
		},
		nil,
		`{
			"message": "bad credentials"
		}`,
		http.StatusForbidden,
	},
	{
		"Not integer id in request as param user_id",
		"first",
		`{
			"id":1, 
			"name":"Artyom", 
			"surname":"Shirshov", 
			"about":"About", 
			"imgUrl":"ImgUrl", 
			"dateOfBirth":"2001-08-06", 
			"city":"msk"
		}`,
		models.AccessDetails{
			AccessUuid: "uuid",
			UserId: 1,
		},
		nil,
		`{
			"message":"bad request"
		}`,
		http.StatusBadRequest,
	},
	{
		"Database problems",
		"1",
		`{
			"id":1, 
			"name":"Artyom", 
			"surname":"Shirshov", 
			"about":"About", 
			"imgUrl":"ImgUrl", 
			"dateOfBirth":"2001-08-06", 
			"city":"msk"
		}`,
		models.AccessDetails{
			AccessUuid: "uuid",
			UserId: 1,
		},
		func(userUsecase *mock.MockUsecase) {
			userUsecase.EXPECT().
				UpdateUser(
					&models.User{
						ID: 1,
						Name: "Artyom",
						Surname: "Shirshov",
						About: "About",
						ImgUrl: "ImgUrl",
						DateOfBirth: "2001-08-06",
						City: "msk",
					}).
				Return(
					&models.User{},
					customErrors.ErrPostgres)
		},
		`{
			"message": "database problems"
		}`,
		http.StatusInternalServerError,
	},
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range updateUserTests{
		t.Run(test.name, func(t *testing.T) {
			mockUserUsecase := mock.NewMockUsecase(ctrl)
			deliveryTest := NewUserDelivery(mockUserUsecase)
			mockMiddlewares := middlewareMock.NewMockMiddleware(ctrl)

			if test.beforeTest != nil {
				test.beforeTest(mockUserUsecase)
			}

			mockMiddlewares.EXPECT().TokenAuthMiddleware().Return(func (ctx *gin.Context){ctx.Next()})
			mockMiddlewares.EXPECT().MiddlewareValidateUser().Return(func (ctx *gin.Context){ctx.Next()})

			ctx, testEngine, userRouter, responseRecorder := prepareTestEnvironment()
			userRouter.Use(func(ctx *gin.Context){
				ctx.Set("access_details", test.inputAccessDetails)
			})
			router.UserEndpoints(userRouter, mockMiddlewares, deliveryTest)

			var err error
			ctx.Request, err = http.NewRequest(http.MethodPost, fmt.Sprintf("/api/users/%s", test.queryParamUserID), strings.NewReader(test.inputBodyJSON))
			assert.Nil(t, err)
			testEngine.ServeHTTP(responseRecorder, ctx.Request)

			assert.Equal(t, test.expectedStatusCode, responseRecorder.Code)
			assert.JSONEq(t, test.expectedUserJSON, responseRecorder.Body.String())
		})
	}
}