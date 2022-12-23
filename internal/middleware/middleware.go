package middleware

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/microservices/auth"
	"Diploma/internal/models"
	"Diploma/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

var allowedOrigins = []string{"", "http://45.141.102.243:8080", "http://127.0.0.1:8080"}

type Middlewares struct {
	auth auth.SessionRepository
}

func NewMiddleware(auth auth.SessionRepository) *Middlewares {
	return &Middlewares{
		auth: auth,
	}
}

func (m *Middlewares) TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	   	au, err := utils.ExtractTokenMetadata(c.Request)
	   	if err != nil {
		  c.JSON(http.StatusUnauthorized, err.Error())
		  c.Abort()
		  return
	   	}

		userId, err := m.auth.FetchAuth(au.AccessUuid)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		if userId != au.UserId {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Set("access_details", *au)
		c.Next()
	}
}

func (m *Middlewares) CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		isAllowed := false
		for _, orig := range allowedOrigins {
			if origin == orig {
				isAllowed = true
				return
			}
		}

		if !isAllowed {
			log.Print("CORS not allowed origin = ", origin)
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
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, models.ErrorMessage{
				Message: customErrors.ErrWrongJson.Error(),
			})
			return
		}

		err = utils.ValidateAndSanitize(inputUser)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, models.ErrorMessage{
				Message: err.Error(),
			})
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
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, models.ErrorMessage{
				Message: customErrors.ErrWrongJson.Error(),
			})
			return
		}

		err = utils.ValidateAndSanitize(inputUser)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, models.ErrorMessage{
				Message: err.Error(),
			})
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
		json.Unmarshal([]byte(inputUser), &user)

		err := utils.ValidateAndSanitize(user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, models.ErrorMessage{
				Message: err.Error(),
			})
			return
		}

		c.Set("user", *user)
		c.Next()
	}
}