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
	r.POST("/signup", mws.MiddlewareValidateUserFormData(), aD.SignUp)
	r.POST("/login", mws.MiddlewareValidateLoginUser(), aD.SignIn)
	r.GET("/logout", mws.TokenAuthMiddleware(), aD.Logout)
	r.POST("/refresh", aD.Refresh)
}

func UserEndpoints(r *gin.RouterGroup, mws *middleware.Middlewares, uD *userD.UserDelivery) {
	r.POST("/:id", mws.TokenAuthMiddleware(), mws.MiddlewareValidateUserFormData(), uD.UpdateUser)
	r.GET("/:id", uD.GetUser)
}

func EventEndpoints(r *gin.RouterGroup, eD *eventD.EventDelivery) {
	r.GET("/:event_id", eD.GetEvent)
	r.GET("/", eD.GetEvents)
	r.GET("/external", eD.GetExternalEvents)
	r.GET("/external/close", eD.GetCloseExternalEvents)
	r.GET("/external/today", eD.GetTodayEvents)
	r.GET("/external/:place_id/:event_id", eD.GetExternalEvent)
}

func PlaceEndpoints(r *gin.RouterGroup, pD *placeD.PlaceDelivery, eD *eventD.EventDelivery) {
	r.GET("/", pD.GetPlaces)
	r.GET("/:place_id", pD.GetPlace)
}