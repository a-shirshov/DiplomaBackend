package delivery

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/microservices/auth"
	"Diploma/internal/models"
	"Diploma/utils"
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	gomail "gopkg.in/mail.v2"
	log "Diploma/pkg/logger"
)

const logMessage = "auth:delivery:"

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
	message := logMessage + "SignUp:"
	log.Debug(message + "started")

	userCtxValue, ok := c.Get("user")
	if !ok {
		utils.SendMessage(c, http.StatusUnprocessableEntity, customErrors.ErrWrongJson.Error())
		return
	}
	user := userCtxValue.(models.User)

	imgUrl, err := utils.SaveImageFromRequest(c, "image")
	if err != nil {
		if err == customErrors.ErrWrongExtension {
			utils.SendMessage(c, http.StatusBadRequest, err.Error())
			return
		}

		log.Error(message + err.Error())
		utils.SendMessage(c, http.StatusInternalServerError, customErrors.ErrSmthWentWrong.Error())
		return
	}
	user.ImgUrl = imgUrl
	

	resultUser, tokenDetails, err := uD.authUsecase.CreateUser(&user)
	if err != nil {
		if err == customErrors.ErrUserExists {
			utils.SendMessage(c, http.StatusConflict, err.Error())
		} else {
			log.Error(message + err.Error())
			utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	tokens := &models.Tokens{
		AccessToken:  tokenDetails.AccessToken,
		RefreshToken: tokenDetails.RefreshToken,
	}

	userWithTokens := &models.UserWithTokens{
		User:   *resultUser,
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
	message := logMessage + "SignIn:"
	log.Debug(message + "started")

	loginUserCtxValue, ok := c.Get("login_user")
	if !ok {
		utils.SendMessage(c, http.StatusUnprocessableEntity, customErrors.ErrWrongJson.Error())
		return
	}
	user := loginUserCtxValue.(models.LoginUser)

	resultUser, tokenDetails, err := uD.authUsecase.SignIn(&user)
	if err != nil {
		if err == customErrors.ErrWrongPassword || err == customErrors.ErrWrongEmail {
			utils.SendMessage(c, http.StatusForbidden, err.Error())
		} else {
			log.Error(message + err.Error())
			utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	tokens := &models.Tokens{
		AccessToken:  tokenDetails.AccessToken,
		RefreshToken: tokenDetails.RefreshToken,
	}

	userWithTokens := &models.UserWithTokens{
		User:   *resultUser,
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
	message := logMessage + "Logout:"
	log.Debug(message + "started")

	au, err := utils.GetAUFromContext(c)
	if err != nil {
		utils.SendMessage(c, http.StatusUnauthorized, err.Error())
		return
	}

	err = uD.authUsecase.Logout(au)
	if err != nil {
		utils.SendMessage(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
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
func (aD *AuthDelivery) Refresh(c *gin.Context) {
	message := logMessage + "Refresh:"
	log.Debug(message + "started")

	var inputTokens models.Tokens
	if err := c.ShouldBindJSON(&inputTokens); err != nil {
		utils.SendMessage(c, http.StatusUnprocessableEntity, customErrors.ErrWrongJson.Error())
		return
	}

	tokens, err := aD.authUsecase.Refresh(inputTokens.RefreshToken)
	if err != nil {
		log.Error(message + err.Error())
		utils.SendMessage(c, http.StatusInternalServerError, customErrors.ErrSmthWentWrong.Error())
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (aD *AuthDelivery) SendEmail(c *gin.Context) {
	message := logMessage + "SendEmail:"
	log.Debug(message + "started")

	redeemStructCtxValue, ok := c.Get("redeem_struct")
	if !ok {
		utils.SendMessage(c, http.StatusUnprocessableEntity, customErrors.ErrWrongJson.Error())
		return
	}
	redeemCodeStruct := redeemStructCtxValue.(models.RedeemCodeStruct)

	user, err := aD.authUsecase.FindUserByEmail(redeemCodeStruct.Email)
	if err != nil {
		utils.SendMessage(c, http.StatusNotFound, err.Error())
		return
	}

	redeemCode, err := aD.authUsecase.CreateAndSavePasswordRedeemCode(redeemCodeStruct.Email)
	if err != nil {
		utils.SendMessage(c, http.StatusNotFound, err.Error())
		return
	}

	m := gomail.NewMessage()
	from := viper.GetString("EMAIL_SENDER")
	m.SetHeader("From", from)
	m.SetHeader("To", redeemCodeStruct.Email)
	m.SetHeader("Subject", "PartyPoint. Заявка на смену пароля.")

	resultMessage := fmt.Sprintf("%s%s.\n\n%s%d.\n%s\n\n%s",
		"Здравствуйте, ", user.Name, "Ваш проверочный код для смены пароля:", redeemCode,
		"Если вы не делали заявку на смену пароля, игнорируйте это сообщение.", "C уважением, команда Partypoint.")
	m.SetBody("text/plain", resultMessage)

	password := viper.GetString("EMAIL_PASSWORD")
	smtpHost := viper.GetString("SMTP_HOST")
	smtpPort := viper.GetInt("SMTP_PORT")

	d := gomail.NewDialer(smtpHost, smtpPort, from, password)
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         viper.GetString("DOMAIN_NAME"),
	}

	if err := d.DialAndSend(m); err != nil {
		log.Error(message + err.Error())
		utils.SendMessage(c, http.StatusBadRequest, "Something went wrong")
		return
	}

	utils.SendMessage(c, http.StatusOK, "OK")
}

func (aD *AuthDelivery) CheckRedeemCode(c *gin.Context) {
	message := logMessage + "CheckRedeemCode:"
	log.Debug(message + "started")

	redeemStructCtxValue, ok := c.Get("redeem_struct")
	if !ok {
		utils.SendMessage(c, http.StatusUnprocessableEntity, customErrors.ErrWrongJson.Error())
		return
	}
	redeemCodeStruct := redeemStructCtxValue.(models.RedeemCodeStruct)

	err := aD.authUsecase.CheckRedeemCode(&redeemCodeStruct)
	if err != nil {
		log.Error(message + err.Error())
		utils.SendMessage(c, http.StatusInternalServerError, customErrors.ErrSmthWentWrong.Error())
		return
	}

	utils.SendMessage(c, http.StatusOK, "OK")
}

func (aD *AuthDelivery) UpdatePassword(c *gin.Context) {
	message := logMessage + "UpdatePassword:"
	log.Debug(message + "started")

	redeemStructCtxValue, ok := c.Get("redeem_struct")
	if !ok {
		utils.SendMessage(c, http.StatusUnprocessableEntity, customErrors.ErrWrongJson.Error())
		return
	}
	redeemCodeStruct := redeemStructCtxValue.(models.RedeemCodeStruct)

	err := aD.authUsecase.UpdatePassword(&redeemCodeStruct)
	if err != nil {
		log.Error(message + err.Error())
		utils.SendMessage(c, http.StatusInternalServerError, customErrors.ErrSmthWentWrong.Error())
		return
	}

	utils.SendMessage(c, http.StatusOK, "OK")
}

func (aD *AuthDelivery) DeleteUser (c *gin.Context) {
	message := logMessage + "DeleteUser:"
	log.Debug(message + "started")

	au, err := utils.GetAUFromContext(c)
	if err != nil {
		utils.SendMessage(c, http.StatusUnauthorized, err.Error())
		return
	}

	err = aD.authUsecase.DeleteUser(au.UserId)
	if err != nil {
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendMessage(c, http.StatusOK, "OK")
}