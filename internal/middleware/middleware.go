package middleware

import (
	"github.com/gin-gonic/gin"
	"Diploma/utils"
	"net/http"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	   err := utils.TokenValid(c.Request)
	   if err != nil {
		  c.JSON(http.StatusUnauthorized, err.Error())
		  c.Abort()
		  return
	   }
	   c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
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