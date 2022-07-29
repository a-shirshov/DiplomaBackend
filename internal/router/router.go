package router

import (
	eventD "Diploma/internal/microservices/event/delivery"
	"Diploma/internal/middleware"
	userD "Diploma/internal/microservices/user/delivery"
	authD "Diploma/internal/microservices/auth/delivery"

	"github.com/gin-gonic/gin"
)

func AuthEndpoints(r *gin.RouterGroup, aD *authD.AuthDelivery) {
	r.POST("/signup", aD.SignUp)
	r.POST("/login", aD.SignIn)
	r.GET("/logout", middleware.TokenAuthMiddleware(), aD.Logout)
	r.POST("/refresh", aD.Refresh)
}

func UserEndpoints(r *gin.RouterGroup, uD *userD.UserDelivery) {
	r.POST("/:id", middleware.TokenAuthMiddleware(), uD.UpdateUser)
	r.GET("/:id", uD.GetUser)
}

func EventEndpoints(r *gin.RouterGroup, eD *eventD.EventDelivery) {
	r.GET("/", eD.GetEvents)
	r.GET("/:id", eD.GetEvent)
}
