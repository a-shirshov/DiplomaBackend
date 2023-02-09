package router

import (
	eventD "Diploma/internal/microservices/event/delivery"
	userD "Diploma/internal/microservices/user/delivery"
	authD "Diploma/internal/microservices/auth/delivery"

	"Diploma/internal/middleware"
	"github.com/gin-gonic/gin"
)

func AuthEndpoints(r *gin.RouterGroup, mws *middleware.Middlewares, aD *authD.AuthDelivery) {
	r.POST("/signup", mws.MiddlewareValidateUserFormData(), aD.SignUp)
	r.POST("/login", mws.MiddlewareValidateLoginUser(), aD.SignIn)
	r.GET("/logout", mws.TokenAuthMiddleware(), aD.Logout)
	r.POST("/refresh", aD.Refresh)
	r.POST("/redeem", mws.MiddlewareValidateRedeemCode(), aD.SendEmail)
	r.POST("/codecheck", mws.MiddlewareValidateRedeemCode(), aD.CheckRedeemCode)
	r.POST("/credentials", mws.MiddlewareValidateRedeemCode(), aD.UpdatePassword)
}

func UserEndpoints(r *gin.RouterGroup, mws *middleware.Middlewares, uD *userD.UserDelivery) {
	r.POST("/:id", mws.TokenAuthMiddleware(), mws.MiddlewareValidateUserFormData(), uD.UpdateUser)
	r.GET("/:id", uD.GetUser)
	r.GET("/:id/favourites", uD.GetFavourites)
	r.POST("/:id/image", uD.UpdateUserImage)
}

func EventEndpoints(r *gin.RouterGroup, mws *middleware.Middlewares, eD *eventD.EventDelivery) {
	r.GET("/external", eD.GetExternalEvents)
	r.GET("/external/close", eD.GetCloseExternalEvents)
	r.GET("/external/today", eD.GetTodayEvents)
	r.GET("/external/:place_id/:event_id", mws.TokenAuthMiddleware(), eD.GetExternalEvent)
	r.POST("external/:event_id/go", mws.TokenAuthMiddleware(), eD.SwitchEventMeeting)
	r.POST("external/:event_id/favourite", mws.TokenAuthMiddleware(), eD.SwitchEventFavourite)
}