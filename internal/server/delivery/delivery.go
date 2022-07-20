package delivery

import (
	"Diploma/internal/models"
	"Diploma/internal/server"
	"Diploma/utils"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserDelivery struct {
	userUsecase server.Usecase
}

func NewUserDelivery(userUsecase server.Usecase) *UserDelivery {
	return &UserDelivery{
		userUsecase: userUsecase,
	}
}

func (uD *UserDelivery) SignUp(c *gin.Context) {
	var user models.User 
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resultUser, err := uD.userUsecase.CreateUser(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated,resultUser)
}

func (uD *UserDelivery) GetUser(c *gin.Context) {
	userIdString := c.Param("id")

	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	resultUser, err := uD.userUsecase.GetUser(userId)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK,resultUser)
}

func (uD *UserDelivery) SignIn(c *gin.Context) {
	var user *models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "user json problem")
	}

	_, tokenDetails, err := uD.userUsecase.SignIn(user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	tokens := &models.Tokens{
		AccessToken: tokenDetails.AccessToken,
		RefreshToken: tokenDetails.RefreshToken,
	}

	c.JSON(http.StatusOK, tokens)
}

func (uD *UserDelivery) Logout(c *gin.Context) {
	au, err := utils.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	err = uD.userUsecase.Logout(au)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
}

func (uD *UserDelivery) Refresh(c *gin.Context) {
	var inputTokens models.Tokens
	if err := c.ShouldBindJSON(&inputTokens); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return 
	}

	tokens, err := uD.userUsecase.Refresh(inputTokens.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, tokens)
}

func (uD *UserDelivery) UpdateUser(c *gin.Context) {
	au, err := utils.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	var inputUser *models.User
	err = c.ShouldBindJSON(&inputUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "wrong json")
		return
	}

	newUser, err := uD.userUsecase.UpdateUser(au, inputUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK,newUser)
}