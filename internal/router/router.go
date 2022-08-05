package router

import (
	eventD "Diploma/internal/microservices/event/delivery"
	userD "Diploma/internal/microservices/user/delivery"
	authD "Diploma/internal/microservices/auth/delivery"
	placeD "Diploma/internal/microservices/place/delivery"

	"Diploma/internal/middleware"
	"github.com/gin-gonic/gin"
)

func AuthEndpoints(r *gin.RouterGroup, mws *middleware.Middlewares, aD *authD.AuthDelivery) {
	r.POST("/signup", mws.MiddlewareValidateUser(), aD.SignUp)
	r.POST("/login", mws.MiddlewareValidateLoginUser(), aD.SignIn)
	r.GET("/logout", mws.TokenAuthMiddleware(), aD.Logout)
	r.POST("/refresh", aD.Refresh)
}

func UserEndpoints(r *gin.RouterGroup, mws *middleware.Middlewares, uD *userD.UserDelivery) {
	r.POST("/:id", mws.TokenAuthMiddleware(), mws.MiddlewareValidateUser(), uD.UpdateUser)
	r.GET("/:id", uD.GetUser)
}

func EventEndpoints(r *gin.RouterGroup, eD *eventD.EventDelivery) {
	r.GET("/", eD.GetEvents)
	r.GET("/:id", eD.GetEvent)
}

func PlaceEndpoints(r *gin.RouterGroup, pD *placeD.PlaceDelivery) {
	r.GET("/", pD.GetPlaces)
	r.GET("/:id", pD.GetPlace)
}