package delivery

// import (
// 	"Diploma/internal/microservices/auth/mock"
// 	"Diploma/internal/middleware"
// 	"Diploma/internal/models"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// )

// type signInTest struct {
// 	name string
// 	inputJSON string
// 	beforeTest func(userUsecase *mock.MockUsecase)
// 	expectedStatusCode int
// 	wantBody string
// }

// var signInTests = []signInTest {
// 	{
// 		"Succesfully signIn",
// 		`{
// 			"email": "mail@yaimage.ru",
// 			"password": "12345678"
// 		}`,
// 		func(userUsecase *mock.MockUsecase) {
// 			userUsecase.EXPECT().
// 				SignIn(&models.User{
// 					Email: "ash@mail.ru",
// 					Password: "password",
// 				}).
// 				Return(&models.User{
// 					ID: 1,
// 					Name: "Artyom",
// 					Surname: "Shirshov",
// 					Email: "ash@mail.ru",
// 					DateOfBirth: "2001-08-06",
// 					City: "msk",
// 					About: "about",
// 					ImgUrl: "user_face.png",
// 				}, 
// 					&models.TokenDetails{
// 						AccessToken: "jwt.valid",
// 						RefreshToken: "jwt.valid",
// 					},
// 					nil,
// 				)
// 		},
// 		200,
// 		`{
// 			"user": {
// 				"id": 1,
// 				"name": "Artyom",
// 				"surname": "Shirshov",
// 				"email": "ash@mail.ru",
// 				"dateOfBirth": "2001-08-06",
// 				"city": "msk",
// 				"about": "about",
// 				"imgUrl": "user_face.png"
// 			},
// 			"tokens": {
// 				"access_token": "jwt.valid",
// 				"refresh_token": "jwt.valid"
// 			}
// 		}`,
// 	},
// }

// func TestSignIn(t *testing.T){
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	for _, test := range signInTests{
// 		t.Run(test.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			router := gin.Default()
// 			mockUserUsecase := mock.NewMockUsecase(ctrl)
// 			deliveryTest := NewAuthDelivery(mockUserUsecase)

// 			if test.beforeTest != nil {
// 				test.beforeTest(mockUserUsecase)
// 			}
// 			router.Use()
// 			responseRecorder := httptest.NewRecorder()
// 			router.POST("/login", deliveryTest.SignIn)
			
// 			request, err := http.NewRequest(http.MethodPost, "/login", strings.NewReader(test.inputJSON))
// 			assert.NoError(t, err)
// 			router.ServeHTTP(responseRecorder, request)
// 			assert.Equal(t, http.StatusOK, responseRecorder.Code)
// 			assert.Equal(t, test.wantBody, responseRecorder.Body.String())
// 		})
// 	}
// }