package delivery

import (
	"Diploma/internal/errors"
	"Diploma/internal/microservices/user"
	"Diploma/internal/models"
	"Diploma/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserDelivery struct {
	userUsecase user.Usecase
}

func NewUserDelivery(userUsecase user.Usecase) *UserDelivery {
	return &UserDelivery{
		userUsecase: userUsecase,
	}
}

// @Summary GetUser
// @Tags Users
// @Description Find a user by id
// @Accept json
// @Produce json
// @Param id path int true "User id"
// @Success 200 {object} models.RegistrationUserResponse
// @Failure 400 {object} models.ErrorMessageBadRequest
// @Failure 500 {object} models.ErrorMessageInternalServer
// @Router /users/{id} [get]
func (uD *UserDelivery) GetUser(c *gin.Context) {
	userIdString := c.Param("id")

	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resultUser, err := uD.userUsecase.GetUser(userId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK,resultUser)
}

// @Summary Update user
// @Security ApiKeyAuth
// @Tags Users
// @Description Update user profile
// @Accept json
// @Produce json
// @Param id path int true "User id"
// @Param newUserInformation body models.UserProfile true "Updated user information"
// @Success 200 {object} models.UserProfile
// @Failure 401 {object} models.ErrorMessageUnauthorized
// @Failure 422 {object} models.ErrorMessageUnprocessableEntity
// @Failure 500 {object} models.ErrorMessageInternalServer
// @Router /users/{id} [post]
func (uD *UserDelivery) UpdateUser(c *gin.Context) {
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		log.Print("here1")
		c.JSON(http.StatusInternalServerError,  err.Error())
		return
	}

	user := c.MustGet("user").(models.User)

	userIdStr := c.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusNotFound,  err.Error())
		return
	}

	if userId != au.UserId {
		c.JSON(http.StatusForbidden,  err.Error())
		return
	}

	imgUrl, err := utils.SaveImageFromRequest(c,"image")
	if err == errors.ErrWrongExtension {
		c.JSON(http.StatusBadRequest,err.Error())
	}
	if err == nil {
		user.ImgUrl = imgUrl
	}

	newUser, err := uD.userUsecase.UpdateUser(au.UserId, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK,newUser)
}
