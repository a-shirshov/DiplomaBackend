package delivery

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/microservices/auth"
	"Diploma/internal/models"
	"Diploma/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthDelivery struct {
	authUsecase auth.Usecase
}

func NewAuthDelivery(authUsecase auth.Usecase) *AuthDelivery {
	return &AuthDelivery{
		authUsecase: authUsecase,
	}
}

// @Summary Registration
// @Tags Auth
// @Description Create a new user
// @Accept json
// @Produce json
// @Param inputUser body models.RegistrationUserRequest true "User data"
// @Success 200 {object} models.RegistrationUserResponse
// @Failure 400 {object} models.ErrorMessageBadRequest
// @Failure 500 {object} models.ErrorMessageInternalServer
// @Router /auth/signup [post]
func (uD *AuthDelivery) SignUp(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	imgUrl, err := utils.SaveImageFromRequest(c,"image")
	if err == customErrors.ErrWrongExtension {
		c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Message: customErrors.ErrWrongExtension.Error(),
		})
		return
	}
	log.Println("Image Err = ", err.Error())
	if err == nil {
		user.ImgUrl = imgUrl
	}

	resultUser, tokenDetails, err := uD.authUsecase.CreateUser(&user)
	if err != nil {
		if err == customErrors.ErrUserExists {
			c.JSON(http.StatusConflict, models.ErrorMessage{
				Message: customErrors.ErrUserExists.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, models.ErrorMessage{
				Message: customErrors.ErrWrongJson.Error(),
			})
		}
		return
	}

	tokens := &models.Tokens{
		AccessToken: tokenDetails.AccessToken,
		RefreshToken: tokenDetails.RefreshToken,
	}

	userWithTokens := &models.UserWithTokens{
		User: *resultUser,
		Tokens: *tokens,
	}

	c.JSON(http.StatusOK, userWithTokens)
}

// @Summary Login
// @Tags Auth
// @Description Login a user
// @Accept json
// @Produce json
// @Param inputCredentials body models.LoginUserRequest true "User credentials"
// @Success 200 {object} models.UserWithTokensResponse
// @Failure 400 {object} models.ErrorMessageBadRequest
// @Failure 500 {object} models.ErrorMessageInternalServer
// @Router /auth/login [post]
func (uD *AuthDelivery) SignIn(c *gin.Context) {
	user := c.MustGet("login_user").(models.LoginUser)

	resultUser, tokenDetails, err := uD.authUsecase.SignIn(&user)
	if err != nil {
		if err == customErrors.ErrWrongPassword || err == customErrors.ErrWrongEmail {
			c.JSON(http.StatusForbidden, models.ErrorMessage{
				Message: err.Error(),
			})
			
		} else {
			c.JSON(http.StatusInternalServerError, models.ErrorMessage{
				Message: err.Error(),
			})
		}
		return
	}

	tokens := &models.Tokens{
		AccessToken: tokenDetails.AccessToken,
		RefreshToken: tokenDetails.RefreshToken,
	}

	userWithTokens := &models.UserWithTokens{
		User: *resultUser,
		Tokens: *tokens,
	}

	c.JSON(http.StatusOK, userWithTokens)
}

// @Summary Logout
// @Security ApiKeyAuth
// @Tags Auth
// @Description Logout
// @Accept json
// @Produce json
// @Success 200 
// @Failure 401 {object} models.ErrorMessageUnauthorized
// @Failure 500 {object} models.ErrorMessageInternalServer
// @Router /auth/logout [get]
func (uD *AuthDelivery) Logout(c *gin.Context) {
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusUnauthorized, err.Error())
		return
	}
	
	err = uD.authUsecase.Logout(au)
	if err != nil {
		utils.SendErrorMessage(c, http.StatusUnauthorized, err.Error())
		return
	}
}

// @Summary Refresh
// @Tags Auth
// @Description Recieve new tokens
// @Accept json
// @Produce json
// @Param RefreshToken body models.RefreshTokenRequest true "RefreshToken"
// @Success 200 {object} models.Tokens
// @Failure 422 {object} models.ErrorMessageBadRequest
// @Failure 500 {object} models.ErrorMessageInternalServer
// @Router /auth/refresh [post]
func (uD *AuthDelivery) Refresh(c *gin.Context) {
	var inputTokens models.Tokens
	if err := c.ShouldBindJSON(&inputTokens); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.ErrorMessage{
			Message: customErrors.ErrWrongJson.Error(),
		})
		return 
	}
	
	log.Print("Refresh token:", inputTokens.RefreshToken)
	tokens, err := uD.authUsecase.Refresh(inputTokens.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, tokens)
}

