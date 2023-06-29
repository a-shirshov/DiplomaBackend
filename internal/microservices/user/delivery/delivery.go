package delivery

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/microservices/user"
	"Diploma/internal/models"
	"Diploma/utils"
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
	userIdString := c.Param("user_id")

	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		utils.SendMessage(c, http.StatusBadRequest, customErrors.ErrBadRequest.Error())
		return
	}

	resultUser, err := uD.userUsecase.GetUser(userId)
	if err != nil {
		if err == customErrors.ErrUserNotFound{
			utils.SendMessage(c, http.StatusNotFound, err.Error())
			return
		}
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resultUser)
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
		utils.SendMessage(c, http.StatusUnauthorized, err.Error())
		return
	}

	var inputUser models.User
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		utils.SendMessage(c, http.StatusUnprocessableEntity, customErrors.ErrWrongJson.Error())
		return
	}

	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		utils.SendMessage(c, http.StatusBadRequest, "bad request")
		return
	}

	if userID != au.UserId {
		utils.SendMessage(c, http.StatusForbidden, "bad credentials")
		return
	}
	inputUser.ID = userID

	user, err := uD.userUsecase.UpdateUser(&inputUser)
	if err != nil {
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uD *UserDelivery) UpdateUserImage(c *gin.Context) {
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		utils.SendMessage(c, http.StatusUnauthorized, err.Error())
		return
	}

	imgUUID, err := utils.SaveImageFromRequest(c, "image")
	if err != nil {
		if err == customErrors.ErrWrongExtension {
			utils.SendMessage(c, http.StatusUnprocessableEntity, err.Error())
		} else {
			utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	user, err := uD.userUsecase.UpdateUserImage(au.UserId, imgUUID)
	if err != nil {
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uD *UserDelivery) ChangePassword(c *gin.Context) {
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		utils.SendMessage(c, http.StatusUnauthorized, err.Error())
		return
	}

	redeemStructCtxValue, ok := c.Get("redeem_struct")
	if !ok {
		utils.SendMessage(c, http.StatusUnprocessableEntity, customErrors.ErrWrongJson.Error())
		return
	}
	redeemCodeStruct := redeemStructCtxValue.(models.RedeemCodeStruct)

	err = uD.userUsecase.ChangePassword(au.UserId, redeemCodeStruct.Password)
	if err != nil {
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendMessage(c, http.StatusOK, "OK")
}