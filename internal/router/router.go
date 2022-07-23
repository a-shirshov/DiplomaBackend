package router

import (
	userD "Diploma/internal/server/delivery"
	eventD "Diploma/internal/event/delivery"
	"Diploma/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserEndpoints(r *gin.RouterGroup, uD *userD.UserDelivery) {
	r.POST("/create", uD.SignUp)
	r.POST("/login", uD.SignIn)
	r.GET("/logout", middleware.TokenAuthMiddleware(), uD.Logout)
	r.POST("/update", middleware.TokenAuthMiddleware(),uD.UpdateUser)
	r.GET("/:id/profile", uD.GetUser)
	r.POST("/refresh", uD.Refresh)
}

func EventEndpoints(r *gin.RouterGroup, eD *eventD.EventDelivery) {
	r.GET("/:page", eD.GetEvents)
}

