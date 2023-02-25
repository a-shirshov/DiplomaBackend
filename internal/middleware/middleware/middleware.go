package middleware

import (
	"Diploma/internal/microservices/auth"
	"Diploma/internal/models"
	"Diploma/pkg"
	"Diploma/utils"
	"strings"

	log "Diploma/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

var allowedOrigins = []string{"", "http://45.141.102.243:8080", "http://127.0.0.1:8080"}

type Middlewares struct {
	auth  auth.SessionRepository
	token pkg.TokenManager
}

func NewMiddleware(auth auth.SessionRepository, token pkg.TokenManager) *Middlewares {
	return &Middlewares{
		auth:  auth,
		token: token,
	}
}

func (m *Middlewares) TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		au, err := m.token.ExtractTokenMetadata(token)
		if err != nil {
			c.Set("access_details", models.AccessDetails{})
			return
		}

		userId, err := m.auth.FetchAuth(au.AccessUuid)
		if err != nil {
			c.Set("access_details", models.AccessDetails{})
			return
		}

		if userId != au.UserId {
			c.Set("access_details", models.AccessDetails{})
			return
		}

		c.Set("access_details", *au)
		c.Next()
	}
}

func (m *Middlewares) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		var isAllowed bool
		for _, orig := range allowedOrigins {
			if origin == orig {
				isAllowed = true
				return
			}
		}

		if !isAllowed {
			log.Error("CORS not allowed origin = ", origin)
			return
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (m *Middlewares) MiddlewareValidateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var inputUser *models.User
		err := c.ShouldBindJSON(&inputUser)
		if err != nil {
			return
		}

		inputUser.Email = strings.ToLower(inputUser.Email)
		log.Debug("Middleware input user", inputUser)
		err = utils.ValidateAndSanitize(inputUser)
		if err != nil {
			log.Error("Error", err)
			return
		}

		c.Set("user", *inputUser)
		c.Next()
	}
}

func (m *Middlewares) MiddlewareValidateLoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var inputUser *models.LoginUser
		err := c.ShouldBindJSON(&inputUser)
		if err != nil {
			return
		}

		inputUser.Email = strings.ToLower(inputUser.Email)
		err = utils.ValidateAndSanitize(inputUser)
		if err != nil {
			return
		}

		c.Set("login_user", *inputUser)
		c.Next()
	}
}

func (m *Middlewares) MiddlewareValidateUserFormData() gin.HandlerFunc {
	return func(c *gin.Context) {
		inputUser := c.Request.FormValue("json")
		user := new(models.User)
		err := json.Unmarshal([]byte(inputUser), &user)
		if err != nil {
			return
		}

		user.Email = strings.ToLower(user.Email)
		err = utils.ValidateAndSanitize(user)
		if err != nil {
			return
		}

		c.Set("user", *user)
		c.Next()
	}
}

func (m *Middlewares) MiddlewareValidateRedeemCode() gin.HandlerFunc {
	return func(c *gin.Context) {
		var inputRedeemCode *models.RedeemCodeStruct
		err := c.ShouldBindJSON(&inputRedeemCode)
		if err != nil {
			return
		}

		err = utils.ValidateAndSanitize(inputRedeemCode)
		if err != nil {
			return
		}
		c.Set("redeem_struct", *inputRedeemCode)
		c.Next()
	}
}

func (m *Middlewares) MiddlewareValidateUserEvent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var inputUserEvent *models.MyEvent
		err := c.ShouldBindJSON(&inputUserEvent)
		if err != nil {
			return
		}

		err = utils.ValidateAndSanitize(inputUserEvent)
		if err != nil {
			return
		}
		c.Set("user_event", *inputUserEvent)
		c.Next()
	}
}
